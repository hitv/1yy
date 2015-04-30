package main

import (
	"github.com/gocraft/web"
	. "hi.tv/1yy/web/app/controllers"
)

func InitRoute(assetPath string) *web.Router {
	root := web.New(Base{}).
		Middleware(web.StaticMiddleware(assetPath)).
		Middleware(web.LoggerMiddleware).
		Middleware((*Base).Init)

	root.SubRouter(Index{}).
		Get("/", (*Index).Index).
		Get("/test", (*Index).Test)

	root.SubRouter(Video{}).
		Get("/:vid:v_(.+)", (*Video).List)

	return root
}
