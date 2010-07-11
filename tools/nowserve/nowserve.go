package main

import (
    "flag"
    "now"
)

var (
    http *string = flag.String("http", ":5939", "HTTP service address")
)

func main() {
    flag.Parse()
    now.Serve(*http)
}

