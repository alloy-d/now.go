package main

import (
    "flag"
    "github.com/alloy-d/now.go"
)

var (
    http *string = flag.String("http", ":5939", "HTTP service address")
)

func main() {
    flag.Parse()
    now.Serve(*http)
}

