package service_test

import (
	"bytes"
	"testing"
)

// MiddlewareTest describes a ServiceTest middleware.
type MiddlewareTest func(ServiceTest) ServiceTest

// LoggingMiddlewareTest takes a buffer as a dependency
// and returns  a ServiceTest middleware.
func LoggingMiddlewareTest(buffer *bytes.Buffer) MiddlewareTest {
	return func(nextTest ServiceTest) ServiceTest {
		return loggingMiddlewareTest{
			buffer: buffer,
			next:   nextTest,
		}
	}
}

type loggingMiddlewareTest struct {
	buffer *bytes.Buffer
	next   ServiceTest
}

func (mwt loggingMiddlewareTest) TestNewSite(t *testing.T) {
	mwt.next.TestNewSite(t)
	want := "method=NewSite siteID=0 err=\"invalid siteID\"\n" +
		"method=NewSite siteID=1 err=null\n"
	get := mwt.buffer.String()
	mwt.buffer.Reset()
	if want != get {
		t.Fatal("NewSite log mismatches\n", "want:\n", want, "get\n", get)
	}
}

func (mwt loggingMiddlewareTest) TestDeleteSite(t *testing.T) {
	mwt.next.TestDeleteSite(t)
	want := "method=DeleteSite siteID=0 err=\"invalid siteID\"\n" +
		"method=DeleteSite siteID=2 err=\"path does not exist\"\n" +
		"method=DeleteSite siteID=1 err=null\n"
	get := mwt.buffer.String()
	mwt.buffer.Reset()
	if want != get {
		t.Fatal("DeleteSite log mismatches\n", "want:\n", want, "get\n", get)
	}
}

func (mwt loggingMiddlewareTest) TestWritePost(t *testing.T) {
	mwt.next.TestWritePost(t)
	want := "method=WritePost siteID=0 filename= err=\"invalid siteID\"\n" +
		"method=WritePost siteID=1 filename=test-post.md err=null\n"
	get := mwt.buffer.String()
	mwt.buffer.Reset()
	if want != get {
		t.Fatal("WritePost log mismatches\n", "want:\n", want, "get\n", get)
	}
}

func (mwt loggingMiddlewareTest) TestRemovePost(t *testing.T) {
	mwt.next.TestRemovePost(t)
	want := "method=RemovePost siteID=0 filename= err=\"invalid siteID\"\n" +
		"method=RemovePost siteID=2 filename=test-post.md err=\"path does not exist\"\n" +
		"method=RemovePost siteID=1 filename=test-post.md err=null\n"
	get := mwt.buffer.String()
	mwt.buffer.Reset()
	if want != get {
		t.Fatal("RemovePost log mismatches\n", "want:\n", want, "get\n", get)
	}
}

func (mwt loggingMiddlewareTest) TestReadPost(t *testing.T) {
	mwt.next.TestReadPost(t)
	want := "method=ReadPost siteID=0 filename= err=\"invalid siteID\"\n" +
		"method=ReadPost siteID=2 filename=test-post.md err=\"path does not exist\"\n" +
		"method=ReadPost siteID=1 filename=test-post.md err=null\n"
	get := mwt.buffer.String()
	mwt.buffer.Reset()
	if want != get {
		t.Fatal("ReadPost log mismatches\n", "want:\n", want, "get\n", get)
	}
}

func (mwt loggingMiddlewareTest) TestWriteConfig(t *testing.T) {
	mwt.next.TestWriteConfig(t)
	want := "method=WriteConfig siteID=0 err=\"invalid siteID\"\n" +
		"method=WriteConfig siteID=1 err=null\n"
	get := mwt.buffer.String()
	mwt.buffer.Reset()
	if want != get {
		t.Fatal("WriteConfig log mismatches\n", "want:\n", want, "get\n", get)
	}
}

func (mwt loggingMiddlewareTest) TestReadConfig(t *testing.T) {
	mwt.next.TestReadConfig(t)
	want := "method=ReadConfig siteID=0 err=\"invalid siteID\"\n" +
		"method=ReadConfig siteID=1 err=null\n"
	get := mwt.buffer.String()
	mwt.buffer.Reset()
	if want != get {
		t.Fatal("ReadConfig log mismatches\n", "want:\n", want, "get\n", get)
	}
}

func (mwt loggingMiddlewareTest) TestUpdateAbout(t *testing.T) {
	mwt.next.TestUpdateAbout(t)
	want := "method=UpdateAbout siteID=0 err=\"invalid siteID\"\n" +
		"method=UpdateAbout siteID=1 err=null\n"
	get := mwt.buffer.String()
	mwt.buffer.Reset()
	if want != get {
		t.Fatal("UpdateAbout log mismatches\n", "want:\n", want, "get\n", get)
	}
}

func (mwt loggingMiddlewareTest) TestReadAbout(t *testing.T) {
	mwt.next.TestReadAbout(t)
	want := "method=ReadAbout siteID=0 err=\"invalid siteID\"\n" +
		"method=ReadAbout siteID=1 err=null\n"
	get := mwt.buffer.String()
	mwt.buffer.Reset()
	if want != get {
		t.Fatal("ReadAbout log mismatches\n", "want:\n", want, "get\n", get)
	}
}
