package service_test

import (
	"bytes"
	"context"
	"github.com/go-kit/kit/log"
	"github.com/golang/mock/gomock"
	mqdispatchmock "github.com/seagullbird/headr-common/mq/dispatch/mock"
	"github.com/seagullbird/headr-repoctl/config"
	"github.com/seagullbird/headr-repoctl/service"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestService(t *testing.T) {
	mockctrl := gomock.NewController(t)
	defer mockctrl.Finish()
	mockDispatcher := mqdispatchmock.NewMockDispatcher(mockctrl)
	mockDispatcher.EXPECT().DispatchMessage(gomock.Any(), gomock.Any()).Return(nil).Times(5)

	var buf bytes.Buffer
	logger := log.NewLogfmtLogger(&buf)
	svc := service.New(mockDispatcher, logger)
	svctest := New(svc, &buf)

	RunTests(t, svctest)
}

// ServiceTest describes a test suite against service.Service
type ServiceTest interface {
	TestNewSite(t *testing.T)
	TestDeleteSite(t *testing.T)
	TestWritePost(t *testing.T)
	TestRemovePost(t *testing.T)
	TestReadPost(t *testing.T)
	TestWriteConfig(t *testing.T)
	TestReadConfig(t *testing.T)
	TestUpdateAbout(t *testing.T)
	TestReadAbout(t *testing.T)
}

// New wires up all ServiceTest middlewares and returns a ServiceTest instance.
func New(svc service.Service, buf *bytes.Buffer) ServiceTest {
	var svctest ServiceTest
	{
		svctest = NewBasicServiceTest(svc)
		svctest = LoggingMiddlewareTest(buf)(svctest)
	}
	return svctest
}

type basicServiceTest struct {
	svc service.Service
}

// NewBasicServiceTest returns a a naïve, stateless implementation of ServiceTest.
func NewBasicServiceTest(svc service.Service) ServiceTest {
	return basicServiceTest{
		svc: svc,
	}
}

// RunTests is the entry of running a ServiceTest.
func RunTests(t *testing.T, svctest ServiceTest) {
	t.Run("NewSite", func(t *testing.T) { clearEnvWrapper(t, svctest.TestNewSite) })
	t.Run("DeleteSite", func(t *testing.T) { clearEnvWrapper(t, svctest.TestDeleteSite) })
	t.Run("WritePost", func(t *testing.T) { clearEnvWrapper(t, svctest.TestWritePost) })
	t.Run("RemovePost", func(t *testing.T) { clearEnvWrapper(t, svctest.TestRemovePost) })
	t.Run("ReadPost", func(t *testing.T) { clearEnvWrapper(t, svctest.TestReadPost) })
	t.Run("WriteConfig", func(t *testing.T) { clearEnvWrapper(t, svctest.TestWriteConfig) })
	t.Run("ReadConfig", func(t *testing.T) { clearEnvWrapper(t, svctest.TestReadConfig) })
	t.Run("UpdateAbout", func(t *testing.T) { clearEnvWrapper(t, svctest.TestUpdateAbout) })
	t.Run("ReadAbout", func(t *testing.T) { clearEnvWrapper(t, svctest.TestReadAbout) })
}

func clearEnvWrapper(t *testing.T, tester func(t *testing.T)) {
	if err := os.RemoveAll(config.SITESDIR); !(err == nil || os.IsNotExist(err)) {
		t.Fatalf("Removing SITESDIR failed: %v", err)
	}

	if err := os.MkdirAll(config.SITESDIR, 0644); err != nil {
		t.Fatalf("Creating SITESDIR failed: %v", err)
	}

	tester(t)
}

func (s basicServiceTest) TestNewSite(t *testing.T) {
	tests := []struct {
		name     string
		siteID   uint
		theme    string
		expected error
	}{
		{"Invalid SiteID 0", 0, "theme", service.ErrInvalidSiteID},
		{"Normal Functioning", 1, "theme", nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := s.svc.NewSite(context.Background(), tt.siteID, tt.theme)
			if output != tt.expected {
				t.Fatalf("siteID=%d, theme=%s, output=%s, expected=%s", tt.siteID, tt.theme, output, tt.expected)
			}
		})
	}
}

func (s basicServiceTest) TestDeleteSite(t *testing.T) {
	tests := []struct {
		name     string
		input    uint
		expected error
	}{
		{"Invalid SiteID 0", 0, service.ErrInvalidSiteID},
		{"Path Not Exists", 2, service.ErrPathNotExist},
		{"Normal Functioning", 1, nil},
	}
	// Create sitepath for site 1
	sitepath := service.SitePath(1)
	if err := os.MkdirAll(sitepath, 0644); err != nil {
		t.Fatalf("Creating sitepath %s failed", sitepath)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := s.svc.DeleteSite(context.Background(), tt.input)
			if output != tt.expected {
				t.Fatalf("siteID=%d, output=%s, expected=%s", tt.input, output, tt.expected)
			}
			if output == nil {
				// Make sure sitepath is truly deleted
				if _, err := os.Stat(sitepath); !os.IsNotExist(err) {
					t.Fatal("Site delete failed, sitepath still exists or something went wrong.")
				}
			}
		})
	}
}

func (s basicServiceTest) TestWritePost(t *testing.T) {
	tests := []struct {
		name     string
		siteID   uint
		filename string
		content  string
		expected error
	}{
		{"Invalid SiteID 0", 0, "", "", service.ErrInvalidSiteID},
		{"Normal Functioning", 1, "test-post.md", "This is a test file", nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := s.svc.WritePost(context.Background(), tt.siteID, tt.filename, tt.content)
			if output != tt.expected {
				t.Fatalf("siteID=%d, filename=%s, config=%s", tt.siteID, tt.filename, tt.content)
			}
			if output == nil {
				// make sure the file is there and its content
				postpath := service.PostPath(1, "test-post.md")
				if _, err := os.Stat(postpath); os.IsNotExist(err) {
					t.Fatalf("write post failed, post path does not exist: %v", err)
				}
				raw, err := ioutil.ReadFile(postpath)
				if err != nil {
					t.Fatalf("Cannot read post: %v", err)
				}
				if string(raw) != "This is a test file" {
					t.Fatalf("write post failed, wrong post content")
				}
			}
		})
	}
}

func (s basicServiceTest) TestRemovePost(t *testing.T) {
	tests := []struct {
		name     string
		siteID   uint
		filename string
		expected error
	}{
		{"Invalid SiteID 0", 0, "", service.ErrInvalidSiteID},
		{"Path Not Exists", 2, "test-post.md", service.ErrPathNotExist},
		{"Normal Functioning", 1, "test-post.md", nil},
	}
	postspath := service.PostsPath(1)
	if err := os.MkdirAll(postspath, 0644); err != nil {
		t.Fatalf("Creating post file directory failed: %v", err)
	}
	postpath := service.PostPath(1, "test-post.md")
	if _, err := os.Create(postpath); err != nil {
		t.Fatalf("Creating post file failed: %v", err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := s.svc.RemovePost(context.Background(), tt.siteID, tt.filename)
			if output != tt.expected {
				t.Fatalf("siteID=%d, filename=%s, config=%s", tt.siteID, tt.filename)
			}
			// make sure its removed
			if output == nil {
				if _, err := os.Stat(postpath); !os.IsNotExist(err) {
					t.Fatal("Post remove failed, sitepath still exists or something went wrong.")
				}
			}
		})
	}
}

func (s basicServiceTest) TestReadPost(t *testing.T) {
	tests := []struct {
		name            string
		siteID          uint
		filename        string
		expectedContent string
		expectedError   error
	}{
		{"Invalid SiteID 0", 0, "", "", service.ErrInvalidSiteID},
		{"Path Not Exists", 2, "test-post.md", "", service.ErrPathNotExist},
		{"Normal Functioning", 1, "test-post.md", "This is a test file", nil},
	}
	postspath := service.PostsPath(1)
	if err := os.MkdirAll(postspath, 0644); err != nil {
		t.Fatalf("Creating post file directory failed: %v", err)
	}
	postpath := service.PostPath(1, "test-post.md")
	if err := ioutil.WriteFile(postpath, []byte("This is a test file"), 0644); err != nil {
		t.Fatalf("Failed to write file content: %v", err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			outputContent, outputError := s.svc.ReadPost(context.Background(), tt.siteID, tt.filename)
			if outputError != tt.expectedError {
				t.Fatalf("siteID=%d, filename=%s, output_content=%s, expected_content=%s, expected_error=%v", tt.siteID, tt.filename, outputContent, tt.expectedContent, tt.expectedError)
			}
			if outputError == nil && outputContent != tt.expectedContent {
				t.Fatalf("siteID=%d, filename=%s, output_content=%s, expected_content=%s, expected_error=%v", tt.siteID, tt.filename, outputContent, tt.expectedContent, tt.expectedError)
			}
		})
	}
}

func (s basicServiceTest) TestWriteConfig(t *testing.T) {
	tests := []struct {
		name     string
		siteID   uint
		config   string
		expected error
	}{
		{"Invalid SiteID 0", 0, "", service.ErrInvalidSiteID},
		{"Normal Functioning", 1, "This is a test config file", nil},
	}

	siteSourcePath := filepath.Join(service.SitePath(1), "source")
	if err := os.MkdirAll(siteSourcePath, 0644); err != nil {
		t.Fatalf("Creating config file directory failed: %v", err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := s.svc.WriteConfig(context.Background(), tt.siteID, tt.config)
			if output != tt.expected {
				t.Fatalf("siteID=%d, config=%s, output=%v, expected=%v", tt.siteID, tt.config, output, tt.expected)
			}
			if output == nil {
				// make sure the file is there and its config
				configPath := filepath.Join(service.SitePath(tt.siteID), "source", "config.json")
				if _, err := os.Stat(configPath); os.IsNotExist(err) {
					t.Fatalf("Config failed, path does not exist: %v", err)
				}
				raw, err := ioutil.ReadFile(configPath)
				if err != nil {
					t.Fatalf("Cannot read config: %v", err)
				}
				if string(raw) != "This is a test config file" {
					t.Fatalf("write post failed, wrong post config: %s", string(raw))
				}
			}
		})
	}
}

func (s basicServiceTest) TestReadConfig(t *testing.T) {
	tests := []struct {
		name            string
		siteID          uint
		expectedContent string
		expectedError   error
	}{
		{"Invalid SiteID 0", 0, "", service.ErrInvalidSiteID},
		{"Normal Functioning", 1, "This is a test file", nil},
	}
	siteSourcePath := filepath.Join(service.SitePath(1), "source")
	if err := os.MkdirAll(siteSourcePath, 0644); err != nil {
		t.Fatalf("Creating config file directory failed: %v", err)
	}
	configPath := filepath.Join(siteSourcePath, "config.json")
	if err := ioutil.WriteFile(configPath, []byte("This is a test file"), 0644); err != nil {
		t.Fatalf("Failed to write file config: %v", err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			outputContent, outputError := s.svc.ReadConfig(context.Background(), tt.siteID)
			if outputError != tt.expectedError {
				t.Fatalf("siteID=%d, output_content=%s, expected_content=%s, expected_error=%v", tt.siteID, outputContent, tt.expectedContent, tt.expectedError)
			}
			if outputError == nil && outputContent != tt.expectedContent {
				t.Fatalf("siteID=%d, output_content=%s, expected_content=%s, expected_error=%v", tt.siteID, outputContent, tt.expectedContent, tt.expectedError)
			}
		})
	}
}

func (s basicServiceTest) TestUpdateAbout(t *testing.T) {
	tests := []struct {
		name     string
		siteID   uint
		content  string
		expected error
	}{
		{"Invalid SiteID 0", 0, "", service.ErrInvalidSiteID},
		{"Normal Functioning", 1, "This is an about file", nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := s.svc.UpdateAbout(context.Background(), tt.siteID, tt.content)
			if output != tt.expected {
				t.Fatalf("siteID=%d, filename=%s, config=%s", tt.siteID, tt.content)
			}
			if output == nil {
				// make sure the file is there and its content
				sitePath := service.SitePath(tt.siteID)
				aboutPath := filepath.Join(sitePath, "source", "content", "about", "_index.md")
				if _, err := os.Stat(aboutPath); os.IsNotExist(err) {
					t.Fatalf("write post failed, post path does not exist: %v", err)
				}
				raw, err := ioutil.ReadFile(aboutPath)
				if err != nil {
					t.Fatalf("Cannot read about: %v", err)
				}
				if string(raw) != "This is an about file" {
					t.Fatalf("update about failed, wrong about content")
				}
			}
		})
	}
}

func (s basicServiceTest) TestReadAbout(t *testing.T) {
	tests := []struct {
		name            string
		siteID          uint
		expectedContent string
		expectedError   error
	}{
		{"Invalid SiteID 0", 0, "", service.ErrInvalidSiteID},
		{"Normal Functioning", 1, "This is an about file", nil},
	}
	sitePath := service.SitePath(1)
	aboutDir := filepath.Join(sitePath, "source", "content", "about")
	if err := os.MkdirAll(aboutDir, 0644); err != nil {
		t.Fatalf("Creating about file directory failed: %v", err)
	}
	aboutPath := filepath.Join(aboutDir, "_index.md")
	if err := ioutil.WriteFile(aboutPath, []byte("This is an about file"), 0644); err != nil {
		t.Fatalf("Failed to write about content: %v", err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			outputContent, outputError := s.svc.ReadAbout(context.Background(), tt.siteID)
			if outputError != tt.expectedError {
				t.Fatalf("siteID=%d, output_content=%s, expected_content=%s, expected_error=%v", tt.siteID, outputContent, tt.expectedContent, tt.expectedError)
			}
			if outputError == nil && outputContent != tt.expectedContent {
				t.Fatalf("siteID=%d, output_content=%s, expected_content=%s, expected_error=%v", tt.siteID, outputContent, tt.expectedContent, tt.expectedError)
			}
		})
	}
}
