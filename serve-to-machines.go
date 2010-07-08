package now

import (
    "fmt"
    "json"
    "github.com/hoisie/web.go"
)

func jsonNext(ctx *web.Context) string {
    out, _ := json.Marshal(GetNext())

    ctx.SetHeader("Content-Type", "application/json", true)
    return fmt.Sprintf("%s", out)
}

func jsonLater(ctx *web.Context) {
    if in, ok := ctx.Request.Params["thing"]; ok {
        for _, thing := range in {
            AddThing(thing)
        }
        ctx.StartResponse(204)  // Success, no comment
    } else {
        ctx.StartResponse(500)  // Vague error message
    }
}

