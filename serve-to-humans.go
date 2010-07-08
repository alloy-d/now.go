package now

import (
    "github.com/hoisie/mustache.go"
    "github.com/hoisie/web.go"
)

func htmlNext(ctx *web.Context) string {
    if thing := GetNext(); thing != "" {
        out, err := mustache.RenderFile("human-friendly/now.mustache.html",
                map[string]string {"thing":thing})
        if err == nil {
            return out
        }
    }
    ctx.StartResponse(404)
    return ""
}

