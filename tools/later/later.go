package main

import (
    "flag"
    "fmt"
    "http"
    "os"
    "strings"
)

var addr *string = flag.String("http", ":5939", "HTTP service address")

func addThing(thing string) {
    _, err := http.Post("http://" + *addr + "/later/", "text/plain", strings.NewReader("thing=" + thing))
    if err != nil {
        fmt.Fprintf(os.Stderr, "http.Post: %v", err)
        os.Exit(1)
    }
}

func main() {
    flag.Parse()
    if flag.NArg() == 0 { os.Exit(1) }

    thing := ""
    for i, w := range flag.Args() {
        if i == 0 {
            thing += w
        } else {
            thing += " " + w
        }
    }

    addThing(thing)
}
