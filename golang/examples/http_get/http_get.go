package main

import (
    "fmt"
    "net/http"
    "os"
    "io/ioutil"
)

func main() {
    url := "http://httpbin.org/get"
    fmt.Println(url)

    resp, err := http.Get(url)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    fmt.Println(string(body))
}

