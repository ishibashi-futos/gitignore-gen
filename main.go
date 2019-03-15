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
	list = flag.Bool("ls", false, "view available template list")
	lang = flag.String("lang", "", "choose language")
)

func main() {
	flag.Parse()
	showList()
	if *lang == "" {
		log.Fatal("error")
	}
	bytes, err := getGitignoreIo(*lang)
	if err != nil {
		log.Fatal(err)
	}
	if err := writeToGitignoreFile(&bytes); err != nil {
		log.Fatal(err)
	}
}

func showList() {
	if *list {
		bytes, err := getGitignoreIo("list?format=lines")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(fmt.Sprintf("templetes:\n%s", bytes))
		os.Exit(0)
	}
}

func getGitignoreIo(path string) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf(url, path))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
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
