package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    fmt.Println("Hello, World!")
    files, err := ioutil.ReadDir("./")
    if err != nil {
        fmt.Println(err)
    }
    for _, file := range files {
        fmt.Println(file.Name(), "is directory?", file.IsDir())
    }
}
