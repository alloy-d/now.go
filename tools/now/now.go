package main

import (
    "bytes"
    "flag"
    "fmt"
    "http"
    "json"
    "os"
    "strings"
)

var (
    addr *string = flag.String("http", ":5939", "HTTP service address")
    done *bool = flag.Bool("done", false, "Current thing is finished")
)

func pop() {
    _, err := http.Post("http://" + *addr + "/done", "text/plain", strings.NewReader("done"))
    if err != nil {
        fmt.Fprintf(os.Stderr, "http.Post: %v", err)
        os.Exit(1)
    }
}

func getNext() string {
    resp, _, err := http.Get("http://" + *addr + "/now.json")
    if err != nil {
        fmt.Fprintf(os.Stderr, "http.Get: %v", err)
        os.Exit(1)
    }

    buf := new(bytes.Buffer)
    buf.ReadFrom(resp.Body)

    var out interface{}
    json.Unmarshal(buf.Bytes(), &out)

    return out.(string)
}

func main() {
    flag.Parse()

    if *done {
        pop()
    }

    thing := getNext()

    fmt.Println(thing)
}
