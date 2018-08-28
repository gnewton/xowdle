package main

import (
	//	"log"
	"errors"
	"io"
	"net/http"
	"time"
)

type Http struct {
	UrlBase
}

func (h *Http) Get() (io.ReadCloser, error) {
	resp, err := getHead(h.url)
	if err != nil {
		if resp != nil {
			h.statusCode = resp.StatusCode
		}
		return nil, err
	}
	h.remoteSize = resp.ContentLength
	h.remoteSizeFromConnect = true
	return resp.Body, err
}

func (h *Http) GetRemoteSize() (int64, error) {
	if h.remoteSizeFromConnect {
		return h.remoteSize, nil
	}
	head, err := getHead(h.url)
	if head != nil {
		h.statusCode = head.StatusCode
	}
	if err != nil {
		return 0, err
	}
	if head.StatusCode != 200 {
		return -1, errors.New("Non 200 http response; response=" + head.Status + "; " + h.url)
	}
	h.remoteSize = head.ContentLength
	h.remoteSizeFromConnect = true
	return h.remoteSize, nil
}

const numHttpSamples = 5

func (h *Http) SampleTime() error {
	//log.Println("SmapleTime: Starting", h.url)
	var err error
	var resp *http.Response
	var elapsed time.Duration
	for j := 0; j < numHttpSamples; j++ {
		start := time.Now()
		resp, err = http.Head(h.url)
		if err != nil {
			continue // we will try again
		}
		if resp.StatusCode != 200 {
			return errors.New("Http response is not 200; is: " + resp.Status + " for url=" + h.url)
		}
		elapsed = time.Since(start)
	}
	h.sampleTime = elapsed
	//log.Println(elapsed)
	return err
}
