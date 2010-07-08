package now

import (
    "github.com/hoisie/web.go"
)

func Serve(http string) {
    web.SetStaticDir("human-friendly")
    web.Get("/now.json", jsonNext)
    web.Get("/now", htmlNext)
    web.Post("/done", jsonPop)
    web.Post("/now", jsonPushNext)
    web.Post("/later/", jsonLater)

    web.Run(http)
}

