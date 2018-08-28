package main

import (
	//"log"
	"testing"
)

const BAD_HOST_NAME = "abfbfbfbfbfbfbfbfbfbfb"
const BAD_HTTP_HOST_NAME = "http://" + BAD_HOST_NAME

const GOOD_HOST_URL = "http://google.com/"

const GOOD_HOST_BAD_FILE = "http://google.com/mmmmmmmmm"

func TestFailsBadHostNameHttp(t *testing.T) {
	u, err := newUrl(BAD_HTTP_HOST_NAME)
	if err != nil {
		t.Error(err)
	}
	_, err = u.GetRemoteSize()
	if err == nil {
		t.Error(err)
	}
}

func TestGoodHost(t *testing.T) {
	u, err := newUrl(GOOD_HOST_URL)
	if err != nil {
		t.Error(err)
	}
	_, err = u.GetRemoteSize()
	if err != nil {
		t.Fatal(err)
	}

	err = u.SampleTime()
	if err != nil {
		t.Error(err)
	}
	_, err = u.Get()
	if err != nil {
		t.Error(err)
	}
}
func TestGoodHostBadFileName(t *testing.T) {
	u, err := newUrl(GOOD_HOST_BAD_FILE)
	if err != nil {
		t.Error(err)
	}
	_, err = u.GetRemoteSize()
	if err == nil {
		t.Fatal(err)
	}

	err = u.SampleTime()
	if err == nil {
		t.Error(err)
	}
	_, err = u.Get()
	if err == nil {
		t.Error(err)
	}
}
