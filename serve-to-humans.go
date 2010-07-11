package now

import (
    "github.com/hoisie/mustache.go"
    "github.com/hoisie/web.go"
    "path"
)

func htmlNext(ctx *web.Context) string {
    if thing := GetNext(); thing != "" {
        template := path.Join(ResourceDir(), "now.mustache.html")
        out, err := mustache.RenderFile(template,
                map[string]string {"thing":thing})
        if err == nil {
            return out
        }
    }
    ctx.StartResponse(404)
    return ""
}

func htmlLater(ctx *web.Context) string {
    template := path.Join(ResourceDir(), "later.mustache.html")
    out, err := mustache.RenderFile(template, nil)
    if err == nil {
        return out
    }
    ctx.StartResponse(404)
    return ""
}

