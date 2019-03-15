package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const url = "https://gitignore.io/api/%s"

var (
	lang = flag.String("lang", "", "choose language")
)

func main() {
	flag.Parse()
	if *lang == "" {
		log.Fatal("error")
	}
	resp, err := http.Get(fmt.Sprintf(url, *lang))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if err := writeToGitignoreFile(&bytes); err != nil {
		log.Fatal(err)
	}
}

func writeToGitignoreFile(bytes *[]byte) error {
	file, err := os.Create("./.gitignore")
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err = file.Write(*bytes); err != nil {
		return err
	}
	return nil
}
