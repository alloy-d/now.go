package now

import (
    "github.com/hoisie/web.go"
)

func Serve(http string) {
    web.SetStaticDir("human-friendly")
    web.Get("/now.json", jsonNext)
    web.Post("/later/", jsonLater)
    web.Get("/now", htmlNext)

    web.Run(http)
}

