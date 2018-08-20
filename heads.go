package main

import (
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

func getHeadInfo(urls []Url) error {
	nftp := defaultNumFtpRoutines
	nhttp := defaultNumHttpRoutines

	var wg sync.WaitGroup

	httpChannel := make(chan Url, nhttp)
	ftpChannel := make(chan Url, nftp)

	for i := 0; i < nhttp; i++ {
		wg.Add(1)
		go getHttpHeads(httpChannel, &wg)
	}

	for i := 0; i < nftp; i++ {
		wg.Add(1)
		go getFtpInfo(ftpChannel, &wg)
	}

	for i, _ := range urls {
		url := urls[i]
		u := url.GetUrl()

		if strings.HasPrefix(u, "ftp://") {
			ftpChannel <- url
		} else {
			//httpChannel <- url
		}
	}
	close(httpChannel)
	close(ftpChannel)
	wg.Wait()

	return nil
}

func getHttpHeads(c chan Url, wg *sync.WaitGroup) {
	defer wg.Done()

	var resp *http.Response
	var err error

	for url := range c {
		u := url.GetUrl()
		log.Println(u)

		var elapsed time.Duration
		for j := 0; j < 5; j++ {
			start := time.Now()
			resp, err = http.Head(u)
			if err != nil {
				log.Println("Failed http HEAD on", u, err)
				continue // or stop
			}
			log.Println("Http status:", resp.Status, u)
			if resp.StatusCode >= 403 { //Forbidden
				break
			}

			if resp.StatusCode >= 400 {
				//err404 := "NOT FOUND 404: " + url
				log.Println("404 Failed http HEAD on", u, err)
				break
				//log.Fatal(errors.New("NOT FOUND 404: " + url))
			}
			elapsed = time.Since(start)

		}

		if resp != nil {
			log.Printf("%s  %d %d %d W=%d  Time: %s", elapsed, elapsed/1000000, resp.ContentLength, resp.ContentLength/10000, int64(elapsed/1000000)*(resp.ContentLength/1000), u)
		}
		//fmt.Println(resp.ContentLength)
		//fmt.Printf("%+v\n", resp)

	}

}
