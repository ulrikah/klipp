package main

import (
	"io/ioutil"
	"log"
	"os"
)

func makeTempFile(name string) *os.File {
	tmpfile, err := ioutil.TempFile("", name)
	if err != nil {
		log.Fatal(err)
	}
	return tmpfile
}

func main() {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		log.Println(file.Name(), "is directory?", file.IsDir())
	}
	tmpfile := makeTempFile("hush")
	defer os.Remove(tmpfile.Name()) // clean up

	message := []byte("\n\n\tHello World\n\n")
	if _, err := tmpfile.Write(message); err != nil {
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}

	content, err := ioutil.ReadFile(tmpfile.Name())
	log.Printf("File contents: %s", content)
	if err != nil {
		log.Fatal(err)
	}
}
