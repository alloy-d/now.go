package main

import (
    "fmt"
    "../now"
)

func main() {
    thing := now.GetNext()

    fmt.Printf("%s\n", thing)
}
