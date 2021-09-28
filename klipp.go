package main

import (
    "os"
    "io/ioutil"
    "log"
)

func main() {
    files, err := ioutil.ReadDir("./")
    if err != nil {
        log.Fatal(err)
    }
    for _, file := range files {
        log.Println(file.Name(), "is directory?", file.IsDir())
    }

	tmpfile, err := ioutil.TempFile("", "tempfile")
	if err != nil {
		log.Fatal(err)
	}
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
