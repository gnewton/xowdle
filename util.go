package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"time"
)

func writeFile(r io.Reader, filename string, timeCreated time.Time) error {
	//file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0600) // For read access.
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0600) // For read access.
	defer func() {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
		//err = os.Chtimes(filename, time, time)
		err = os.Chtimes(filename, time.Now(), timeCreated)
		if err != nil {
			log.Println(err)
		}
	}()
	if err != nil {
		log.Println(err)
		return err
	}

	w := bufio.NewWriter(file)

	if _, err := io.Copy(w, r); err != nil {
		log.Println("Error writing:", filename)
		log.Println(err)
		return err
	}

	return nil
}
