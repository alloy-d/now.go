package now

import (
    "github.com/hoisie/web.go"
)

func Serve(http string) {
    web.SetStaticDir(ResourceDir())
    web.Get("/now.json", jsonNext)
    web.Get("/now", htmlNext)
    web.Get("/later/", htmlLater)
    web.Get("/later", func (c *web.Context) { c.Redirect(301, "/later/") })
    web.Post("/done", jsonPop)
    web.Post("/now", jsonPushNext)
    web.Post("/later/", jsonLater)
    web.Delete("/now", jsonPop)
    web.Get("/", func (c *web.Context) { c.Redirect(301, "/now") })

    web.Run(http)
}

