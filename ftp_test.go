package main

import (
	"log"
	"testing"
)

func TestFtpSplit(t *testing.T) {
	urls := []string{"ftp://ftp.foo.com/a/file.tgz", "ftp://ftp.foo.com/a/b/file.tgz", "ftp://ftp.foo.com/file.tgz"}

	for i, _ := range urls {
		url := urls[i]
		host, dir, file := ftpSplit(url)
		log.Println("url", url)
		log.Println("host", host)
		log.Println("dir", dir)
		log.Println("file", file)
	}
	//if err != nil {
	//t.Error(err)
	//}

}
