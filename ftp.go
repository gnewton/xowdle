package main

import (
	"github.com/jlaffaye/ftp"
	"log"
	"time"
)

var m string = "greengenes.microbio.me"

func ftpInfo(host, dir, file string) (exists bool, length int64, e error) {
	//func ftpInfo(host string) (exists bool, length int64, e error) {
	//log.Println("FTP [" + host + "]")

	c, err := ftp.DialTimeout(host+":21", 10*time.Second)
	//_, err := ftp.DialTimeout("greengenes.microbio.me"+":21", 10*time.Second)

	if err != nil {
		log.Println("FTP: Connect failed for host:", host)
		return false, -1, err
	}

	err = c.Login("anonymous", "anonymous")
	if err != nil {
		log.Println("FTP: Anonymnous login failed for host host:", host)
		return false, -1, err
	}

	err = c.ChangeDir(dir)
	if err != nil {
		log.Println("FTP: chdir host:[", host, "] Dir:[", dir, "]")
		return false, -1, err
	}

	return false, 0, nil
}
