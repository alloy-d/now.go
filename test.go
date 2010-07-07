package main

import (
    "./next"
    "fmt"
)

func doWrite() {
    next.AddThing("make now.go better.")
}

func doRead() string {
    thing := next.GetNext()
    return thing
}

func main() {
    doWrite()
    thing := doRead()

    fmt.Printf("%v\n", thing);
}
