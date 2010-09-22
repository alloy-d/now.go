package now

import (
    "github.com/hoisie/mustache.go"
    "github.com/hoisie/web.go"
    "path"
)

// TODO: return the proper response code if something goes wrong with a
// template.

func htmlNext(ctx *web.Context) string {
    if thing := GetNext(); thing != "" {
        template := path.Join(ResourceDir(), "now.mustache.html")
        out := mustache.RenderFile(template,
                map[string]string {"thing":thing})
        return out
    }
    ctx.StartResponse(404)
    return ""
}

func htmlLater(ctx *web.Context) string {
    template := path.Join(ResourceDir(), "later.mustache.html")
    out := mustache.RenderFile(template, nil)
    return out
}

