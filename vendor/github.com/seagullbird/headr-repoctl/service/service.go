package service

import (
	"context"
	"github.com/seagullbird/headr-repoctl/mq_helper"
	"time"
	"github.com/go-kit/kit/log"
	"os/exec"
	"path/filepath"
	"github.com/seagullbird/headr-repoctl/config"
	"os"
)

type Service interface {
	NewSite(ctx context.Context, email, sitename string) error
	DeleteSite(ctx context.Context, email, sitename string) error
}

func New(dispatcher mq_helper.Dispatcher, logger log.Logger) Service {
	var svc Service
	{
		svc = NewBasicService(dispatcher)
		svc = LoggingMiddleware(logger)(svc)
	}
	return svc
}

type basicService struct {
	dispatcher mq_helper.Dispatcher
}

func NewBasicService(dispatcher mq_helper.Dispatcher) basicService {
	return basicService{
		dispatcher: dispatcher,
	}
}

func (s basicService) NewSite(ctx context.Context, email, sitename string) error {
	evt := mq_helper.NewSiteEvent{
		email,
		sitename,
		time.Now().Unix(),
	}
	return s.dispatcher.DispatchMessage(evt)
}

func (s basicService) DeleteSite(ctx context.Context, email, sitename string) error {
	sitepath := filepath.Join(config.SITESDIR, email, sitename)
	if _, err := os.Stat(sitepath); err != nil {
		if os.IsNotExist(err) {
			return MakeErrPathNotExist(sitepath)
		}
		return MakeErrUnexpected(err)
	}
	cmd := exec.Command("rm", "-rf", sitepath)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
