package main

import(
	"log"
	"io"
	"time"
	"net/http"
	"errors"
)

type Http struct {
	UrlBase
}


func (h *Http) Get() (io.ReadCloser, error){
	resp, err := getHead(h.url)
	h.remoteSize = resp.ContentLength
	h.remoteSizeFromConnect = true
	return resp.Body, err
}		

func (h *Http) GetRemoteSize() (int64,error) {
	if h.remoteSizeFromConnect{
		return h.remoteSize, nil
	}
	head, err := getHead(h.url)
	if err != nil{
		return 0, err
	}
	h.remoteSize = head.ContentLength
 	h.remoteSizeFromConnect = true
	return h.remoteSize, nil
}

const numHttpSamples = 5

func (h *Http) SampleTime() error{
	log.Println("SmapleTime: Starting", h.url)
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
			err = errors.New("Http response is not 200; is: "+resp.Status + " for url="+h.url )
		}
		elapsed = time.Since(start)
	}
	h.sampleTime = elapsed
	log.Println(elapsed)
	return err
}
