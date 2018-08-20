package main

import (
	"github.com/jlaffaye/ftp"
	//"github.com/secsy/goftp"
	//"io/ioutil"
	"log"
	"strings"
	"sync"
	"time"
)

var m string = "greengenes.microbio.me"

func ftpInfo(host, dir, file string) (exists bool, length int64, e error) {
	//func ftpInfo(host string) (exists bool, length int64, e error) {
	//log.Println("FTP [" + host + "]")
	log.Println("FTP START ", host)

	c, err := ftp.DialTimeout(host+":21", 10*time.Second)
	defer c.Quit()

	if err != nil {
		log.Println("Connect ", host, ":", err)
		//log.Println("FTP: Connect failed for host:", host)
		//return false, -1, err
	}

	err = c.Login("anonymous", "anonymous")
	if err != nil {
		log.Println("Login ", host, ":", err)
	}

	size, err := c.FileSize(dir + file)
	log.Println("size ", host, dir, file, size)

	//fi, err := c.Stat(file)
	// fi, err := c.Stat(dir)
	// if err != nil {
	// 	log.Println("FTP: Connect stat for host/file:", host, dir, file)
	// 	return false, -1, err
	// }
	//log.Println(fi)

	err = c.ChangeDir(dir)
	if err != nil {
		log.Println("ChangeDir ", host, ":", err)
	}

	entries, err := c.List(file)
	if err != nil {
		log.Println("List ", host, dir, file, " ", err)
	}
	if len(entries) != 1 {
		log.Fatal("File does not exist: ", host, dir, file)
	}
	for i, _ := range entries {
		log.Println("++++++++++++++   ", entries[i])
	}

	// Read without deadline
	r, err := c.Retr(file)
	defer r.Close()
	if err != nil {
		log.Println("Retrieve ", host, ":", file, " ", err)
	} else {
		err := writeFile(r, file)
		//_, err := ioutil.ReadAll(r)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = c.Logout()
	if err != nil {
		log.Println("Logout ", host, ":", err)
	}

	log.Println("FTP DONE ", host)

	return false, 0, nil
}

func getFtpInfo(c chan Url, wg *sync.WaitGroup) {
	defer wg.Done()
	//fmt.Println("*********")
	for u := range c {
		url := u.GetUrl()
		host, dir, file := ftpSplit(url)
		var elapsed time.Duration
		start := time.Now()
		_, _, err := ftpInfo(host, dir, file)
		if err != nil {
			log.Println("Failed FTP host=", host, err)
			log.Println(err)
		}
		elapsed = time.Since(start)
		log.Println("host  elapsed ", elapsed)

	}
}

func ftpSplit(url string) (host string, dir string, file string) {
	s := strings.TrimPrefix(url, "ftp://")
	parts := strings.SplitN(s, "/", 2)
	host, rest := parts[0], parts[1]
	n := strings.LastIndex(rest, "/")
	if n <= 0 {
		file = rest
		dir = "/"
	} else {
		dir = "/" + rest[0:n] + "/"
		file = rest[n+1:]

	}

	return host, dir, file
}
