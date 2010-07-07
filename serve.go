package main

import (
    "fmt"
    "json"
    "github.com/hoisie/web.go"
    "./next"
)

func jsonNext() string {
    out, _ := json.Marshal(next.GetNext())

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
    web.Get("/next", jsonNext)
    web.Post("/later/", jsonLater)
    web.Run("0.0.0.0:5939")
}

