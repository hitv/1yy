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

	root.Subrouter(Index{}, "").
		Get("/", (*Index).Index)

	return root
}
