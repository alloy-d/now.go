package main

import (
    "flag"
    "os"
    "../next"
)

func main() {
    if flag.NArg() == 0 { os.Exit(1) }

    thing := ""
    for i, w := range flag.Args() {
        if i == 0 {
            thing += w
        } else {
            thing += " " + w
        }
    }

    next.AddThing(thing)
}
