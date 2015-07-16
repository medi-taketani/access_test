package main

import (
    "net/http"
    "fmt"
    "sync"
)

func fetch () {
    response,err := http.Get("http://192.168.66.10/")

    if err != nil {

    }

    fmt.Println(response.Status)

}

func main () {
    var wg sync.WaitGroup
    for i := 0; i < 1000; i ++ {
        wg.Add(1)
        go func () {
            fetch()
            wg.Done()
        }()
    }
    wg.Wait()
}
