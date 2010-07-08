package now

import (
    "fmt"
    "json"
    "github.com/hoisie/web.go"
)

func jsonNext(ctx *web.Context) string {
    if thing := GetNext(); thing != "" {
        out, _ := json.Marshal(thing)

        ctx.SetHeader("Content-Type", "application/json", true)
        return fmt.Sprintf("%s", out)
    }
    ctx.StartResponse(418)  // I'm a teapot
    return ""
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

func jsonPop(ctx *web.Context) string {
    if _, ok := ctx.Request.Params["done"]; ok {
        Pop()
        return jsonNext(ctx)
    }

    ctx.StartResponse(400)
    return ""
}

func jsonPushNext(ctx *web.Context) string {
    if thing, ok := ctx.Request.Params["thing"]; ok {
        if len(thing) > 1 {
            ctx.StartResponse(400)
            return ""
        }

        PushNext(thing[0])
    }
    return jsonNext(ctx)
}
