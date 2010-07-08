package main

import (
    "./now"
    "fmt"
    "json"
)

func doWrite() {
    now.AddThing("make now.go better.")
}

func doRead() string {
    thing := now.GetNext()
    return thing
}

func main() {
    thing := now.Thing{"test now.go"}
    jsonThing, _ := json.Marshal(thing)
    fmt.Printf("%s", jsonThing)
}
