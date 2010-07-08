package main

import (
    "flag"
    "fmt"
    "json"
    "github.com/hoisie/web.go"
    "./next"
)

var (
    http *string = flag.String("http", ":5939", "HTTP service address")
)

func htmlNext() string {
    thing := next.GetNext()

    return fmt.Sprintf(
            "<!DOCTYPE html>\n<html><head><title>Now.</title></head><body><p>%s</p></body></html>",
            thing)
}


func jsonNext(ctx *web.Context) string {
    out, _ := json.Marshal(next.GetNext())

    ctx.SetHeader("Content-Type", "application/json", true)
    return fmt.Sprintf("%s", out)
}

func jsonLater(ctx *web.Context) {
    if in, ok := ctx.Request.Params["thing"]; ok {
        for _, thing := range in {
            next.AddThing(thing)
        }
        ctx.StartResponse(204)
    } else {
        ctx.StartResponse(500)
    }
}

func main() {
    flag.Parse()
    web.Get("/next.json", jsonNext)
    web.Post("/later/", jsonLater)
    web.Get("/next", htmlNext)
    web.Run(*http)
}

