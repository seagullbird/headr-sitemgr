package service

//go:generate mockgen -destination=./mock/mock_service.go -package=mock github.com/seagullbird/headr-repoctl/service Service

import (
	"context"
	"errors"
	"github.com/go-kit/kit/log"
	"github.com/seagullbird/headr-common/mq"
	"github.com/seagullbird/headr-common/mq/dispatch"
	"github.com/seagullbird/headr-repoctl/config"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// Service describes a service that deals with local files in the persistent volume (repoctl).
type Service interface {
	NewSite(ctx context.Context, siteID uint, theme string) error
	DeleteSite(ctx context.Context, siteID uint) error
	WritePost(ctx context.Context, siteID uint, filename, content string) error
	RemovePost(ctx context.Context, siteID uint, filename string) error
	ReadPost(ctx context.Context, siteID uint, filename string) (content string, err error)
	WriteConfig(ctx context.Context, siteID uint, config string) error
	ReadConfig(ctx context.Context, siteID uint) (string, error)
	UpdateAbout(ctx context.Context, siteID uint, content string) error
	ReadAbout(ctx context.Context, siteID uint) (string, error)
}

// New returns a basic Service with all of the expected middlewares wired in.
func New(dispatcher dispatch.Dispatcher, logger log.Logger) Service {
	var svc Service
	{
		svc = NewBasicService(dispatcher)
		svc = LoggingMiddleware(logger)(svc)
	}
	return svc
}

type basicService struct {
	dispatcher dispatch.Dispatcher
}

// NewBasicService returns a naïve, stateless implementation of Service.
func NewBasicService(dispatcher dispatch.Dispatcher) Service {
	return basicService{
		dispatcher: dispatcher,
	}
}

// ErrPathNotExist indicates a PathNotExist error
var ErrPathNotExist = errors.New("path does not exist")

// ErrUnexpected indicates an unexpected error
var ErrUnexpected = errors.New("unexpected error")

// ErrInvalidSiteID indicates an invalid SiteID
// Typically a SiteID <= 0
var ErrInvalidSiteID = errors.New("invalid siteID")

func (s basicService) NewSite(ctx context.Context, siteID uint, theme string) error {
	if siteID <= 0 {
		return ErrInvalidSiteID
	}

	evt := mq.SiteUpdatedEvent{
		SiteID:     siteID,
		Theme:      theme,
		ReceivedOn: time.Now().Unix(),
	}
	return s.dispatcher.DispatchMessage("new_site", evt)
}

func (s basicService) DeleteSite(ctx context.Context, siteID uint) error {
	if siteID <= 0 {
		return ErrInvalidSiteID
	}

	sitepath := SitePath(siteID)
	if _, err := os.Stat(sitepath); err != nil {
		if os.IsNotExist(err) {
			return ErrPathNotExist
		}
		return ErrUnexpected
	}
	return os.RemoveAll(sitepath)
}

func (s basicService) WritePost(ctx context.Context, siteID uint, filename, content string) error {
	if siteID <= 0 {
		return ErrInvalidSiteID
	}

	postsPath := PostsPath(siteID)
	if _, err := os.Stat(postsPath); err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(postsPath, 0644)
		} else {
			return err
		}
	}
	postPath := PostPath(siteID, filename)
	if err := ioutil.WriteFile(postPath, []byte(content), 0644); err != nil {
		return err
	}
	// Generate site
	evt := mq.SiteUpdatedEvent{
		SiteID:     siteID,
		ReceivedOn: time.Now().Unix(),
	}
	return s.dispatcher.DispatchMessage("re_generate", evt)
}

func (s basicService) RemovePost(ctx context.Context, siteID uint, filename string) error {
	if siteID <= 0 {
		return ErrInvalidSiteID
	}

	postPath := PostPath(siteID, filename)
	if _, err := os.Stat(postPath); err != nil {
		if os.IsNotExist(err) {
			return ErrPathNotExist
		}
		return ErrUnexpected
	}
	if err := os.Remove(postPath); err != nil {
		return err
	}
	// Generate site
	evt := mq.SiteUpdatedEvent{
		SiteID:     siteID,
		ReceivedOn: time.Now().Unix(),
	}
	return s.dispatcher.DispatchMessage("re_generate", evt)
}

func (s basicService) ReadPost(ctx context.Context, siteID uint, filename string) (content string, err error) {
	if siteID <= 0 {
		return "", ErrInvalidSiteID
	}

	postPath := PostPath(siteID, filename)
	if _, err := os.Stat(postPath); err != nil {
		if os.IsNotExist(err) {
			return "", ErrPathNotExist
		}
		return "", ErrUnexpected
	}
	contentRaw, err := ioutil.ReadFile(postPath)
	if err != nil {
		return "", ErrUnexpected
	}
	return string(contentRaw), nil
}

func (s basicService) WriteConfig(ctx context.Context, siteID uint, config string) error {
	if siteID <= 0 {
		return ErrInvalidSiteID
	}

	sitePath := SitePath(siteID)
	configFilePath := filepath.Join(sitePath, "source", "config.json")
	err := ioutil.WriteFile(configFilePath, []byte(config), 0644)
	if err != nil {
		return err
	}
	// Generate site
	evt := mq.SiteUpdatedEvent{
		SiteID:     siteID,
		ReceivedOn: time.Now().Unix(),
	}
	return s.dispatcher.DispatchMessage("re_generate", evt)
}

func (s basicService) ReadConfig(ctx context.Context, siteID uint) (string, error) {
	if siteID <= 0 {
		return "", ErrInvalidSiteID
	}

	sitePath := SitePath(siteID)
	configFilePath := filepath.Join(sitePath, "source", "config.json")
	configRaw, err := ioutil.ReadFile(configFilePath)
	return string(configRaw), err
}

func (s basicService) UpdateAbout(ctx context.Context, siteID uint, content string) error {
	if siteID <= 0 {
		return ErrInvalidSiteID
	}

	sitePath := SitePath(siteID)
	aboutDir := filepath.Join(sitePath, "source", "content", "about")
	if _, err := os.Stat(aboutDir); err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(aboutDir, 0644)
		} else {
			return err
		}
	}
	aboutPath := filepath.Join(aboutDir, "_index.md")
	if err := ioutil.WriteFile(aboutPath, []byte(content), 0644); err != nil {
		return err
	}
	// Generate site
	evt := mq.SiteUpdatedEvent{
		SiteID:     siteID,
		ReceivedOn: time.Now().Unix(),
	}
	return s.dispatcher.DispatchMessage("re_generate", evt)
}

func (s basicService) ReadAbout(ctx context.Context, siteID uint) (string, error) {
	if siteID <= 0 {
		return "", ErrInvalidSiteID
	}

	sitePath := SitePath(siteID)
	aboutPath := filepath.Join(sitePath, "source", "content", "about", "_index.md")
	if _, err := os.Stat(aboutPath); err != nil {
		if os.IsNotExist(err) {
			return "", ErrPathNotExist
		}
		return "", ErrUnexpected
	}
	contentRaw, err := ioutil.ReadFile(aboutPath)
	if err != nil {
		return "", ErrUnexpected
	}
	return string(contentRaw), nil
}

// SitePath is the root directory of a site. Typically has a public as well as a source sub-directory.
func SitePath(siteID uint) string {
	return filepath.Join(config.SITESDIR, strconv.Itoa(int(siteID)))
}

// PostsPath is the root directory for all posts.
func PostsPath(siteID uint) string {
	return filepath.Join(config.SITESDIR, strconv.Itoa(int(siteID)), "source", "content", "posts")
}

// PostPath is the path of a particular post file.
func PostPath(siteID uint, filename string) string {
	return filepath.Join(config.SITESDIR, strconv.Itoa(int(siteID)), "source", "content", "posts", filename)
}
