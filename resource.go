package main

import (
	"errors"
	"log"
	"strings"
	"time"
)

type Url interface {
	GetUrl() string
	SetUrl(string)
	GetRemoteSize() int64
	GetSampleTime() time.Duration
}

type UrlBase struct {
	url        string
	remoteSize int64
	sampleTime time.Duration
}

func (u *UrlBase) GetUrl() string {
	return u.url
}

func (u *UrlBase) SetUrl(url string) {
	u.url = url
}

func (u *UrlBase) GetRemoteSize() int64 {
	return u.remoteSize
}

func (u *UrlBase) GetSampleTime() time.Duration {
	return u.sampleTime
}

type Ftp struct {
	UrlBase
}

type Http struct {
	UrlBase
}

func newUrls(urls []string) ([]Url, error) {
	resources := make([]Url, 0)
	for i, _ := range urls {
		url := urls[i]
		if strings.HasPrefix(url, "http") {
			newr := new(Http)
			newr.SetUrl(url)
			resources = append(resources, newr)
		} else if strings.HasPrefix(url, "ftp://") {
			newr := new(Ftp)
			newr.SetUrl(url)
			resources = append(resources, newr)
		} else {
			err := errors.New("Unsupported resource type:" + url)
			log.Println(err)
			return nil, err
		}

	}
	return resources, nil
}
