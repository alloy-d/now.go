package main

import (
    "flag"
    "fmt"
    "../now"
)

var done *bool = flag.Bool("done", false, "Current thing is finished")

func main() {
    flag.Parse()
    if *done {
        now.Pop()
    }

    thing := now.GetNext()

    fmt.Printf("%s\n", thing)
}
