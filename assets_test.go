package main

import (
	"io"
	"os"

	"testing"
)

func open(file string, t *testing.T) { 
	f, err := Assets.Open(file)
	if err != nil {
		t.Fatal(err)
	}
	io.Copy(os.Stdout, f)
}

func TestListAssets(t *testing.T) {
	open("/assets/list/list.txt", t)
}

func TestLanguageAssets(t *testing.T) {
	open("/assets/languages/java.txt", t)
	open("/assets/languages/csharp.txt", t)
	open("/assets/languages/go.txt", t)
}