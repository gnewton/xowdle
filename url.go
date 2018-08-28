package main

import (
	"errors"
	"io"
	"log"
	"strings"
	"time"
)

type Url interface {
	Init()
	Url() string
	SetUrl(string)
	GetRemoteSize() (int64, error)
	GetSampleTime() time.Duration
	GetResponseCode() int
	SampleTime() error
	Get() (io.ReadCloser, error)
}

type UrlBase struct {
	url                   string
	remoteSize            int64
	remoteSizeFromConnect bool
	sampleTime            time.Duration
	statusCode            int
}

func (u *UrlBase) GetResponseCode() int {
	return u.statusCode
}

func (u *UrlBase) Url() string {
	return u.url
}

func (u *UrlBase) Init() {

}

func (u *UrlBase) SetUrl(url string) {
	u.url = url
}

func (u *UrlBase) GetSampleTime() time.Duration {
	return u.sampleTime
}

func newUrl(url string) (Url, error) {
	var newUrl Url
	if strings.HasPrefix(url, "http") {
		newUrl = new(Http)
	} else if strings.HasPrefix(url, "ftp://") {
		newUrl = new(Ftp)
	} else {
		err := errors.New("Unsupported resource type:" + url)
		log.Println(err)
		return nil, err
	}
	newUrl.Init()
	newUrl.SetUrl(url)

	return newUrl, nil
}

func newUrls(urls []string) ([]Url, error) {
	resources := make([]Url, 0)
	for i, _ := range urls {
		url := urls[i]
		newu, err := newUrl(url)
		if err != nil {
			return nil, err
		}
		resources = append(resources, newu)
	}
	return resources, nil
}
