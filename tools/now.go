package main

import (
    "fmt"
    "../next"
)

func main() {
    thing := next.GetNext()

    fmt.Printf(thing)
}
