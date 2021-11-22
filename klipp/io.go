package klipp

import (
	"io/ioutil"
	"log"
	"os"

	"golang.design/x/clipboard"
)

type Klipp struct {
	HomeDir string
}

func (k Klipp) pathTo(relPath string) string {
	// TODO: probably some util in ioutils to concat paths
	return k.HomeDir + "/" + relPath
}

func (k Klipp) Read(noteName string) string {
	// Reads from the specified note and copies it to the clipboard
	content, err := ioutil.ReadFile(k.pathTo(noteName))
	if err != nil {
		return ""
	} else {
		copyToClipboard(string(content))
		return string(content)
	}
}

func (k Klipp) Write(note string) string {
	// Writes to a specified note from the clipboard

	// TODO: properly handle errors
	absPath := k.pathTo(note)
	_, err := ioutil.ReadFile(absPath)
	if err != nil {
		ioutil.WriteFile(absPath, []byte(pasteFromClipboard()), 0644)
		pasteFromClipboard()
		return "success"
	}
	return "failure"
}

func makeTempFile(name string) *os.File {
	tmpfile, err := ioutil.TempFile("", name)
	if err != nil {
		log.Fatal(err)
	}
	return tmpfile
}

func pasteFromClipboard() string {
	return string(clipboard.Read(clipboard.FmtText))
}

func copyToClipboard(msg string) <-chan struct{} {
	return clipboard.Write(clipboard.FmtText, []byte(msg))
}

// k := Klipp{HomeDir: "/Users/ulrikah/.klipp"}
// k.Write("test")
// log.Printf("In the clipboard: %s", pasteFromClipboard())
// for i := 0; i < 10; i++ {
// 	copyToClipboard(fmt.Sprintf("%s%d", "Hello: ", i))
// 	log.Printf("In the clipboard: %s", pasteFromClipboard())
// }
// files, err := ioutil.ReadDir("./")
// if err != nil {
// 	log.Fatal(err)
// }
// for _, file := range files {
// 	log.Println(file.Name(), "is directory?", file.IsDir())
// }
// tmpfile := makeTempFile("hush")
// defer os.Remove(tmpfile.Name()) // clean up

// message := []byte("\n\n\tHello World\n\n")
// if _, err := tmpfile.Write(message); err != nil {
// 	log.Fatal(err)
// }
// if err := tmpfile.Close(); err != nil {
// 	log.Fatal(err)
// }

// content, err := ioutil.ReadFile(tmpfile.Name())
// log.Printf("File contents: %s", content)
// if err != nil {
// 	log.Fatal(err)
// }
