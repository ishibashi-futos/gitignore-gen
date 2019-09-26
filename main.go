package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	url = "https://gitignore.io/api/%s"
	listFile = "/assets/list/list.txt"
	languageFile = "/assets/languages/%s.txt"
)

var (
	list = flag.Bool("ls", false, "view available template list")
	lang = flag.String("lang", "", "choose language")
)

 //go:generate go-assets-builder assets/ > assets.go
func main() {
	flag.Parse()
	if *list {
		showList()
	}
	if *lang == "" {
		log.Fatal("error")
	}
	// bytes, err := getGitignoreIo(*lang)
	f, err := Assets.Open(fmt.Sprintf(languageFile, *lang))
	if err != nil {
		log.Fatal(err)
	}
	if err := writeToGitignoreFile(read(f)); err != nil {
		log.Fatal(err)
	}
}

func read(file http.File) []byte {
	buf := new(bytes.Buffer)
	io.Copy(buf, file)
	return buf.Bytes()
}

func showList() {
	f, err := Assets.Open(listFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
	fmt.Println(fmt.Sprintf("templetes:\n%s", read(f)))
	os.Exit(0)
}

func getGitignoreIo(path string) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf(url, path))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func writeToGitignoreFile(bytes []byte) error {
	file, err := os.Create("./.gitignore")
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err = file.Write(bytes); err != nil {
		return err
	}
	return nil
}
