package main

import (
	"github.com/gofly/web"
	. "hi.tv/1yy/web/app/controllers"
)

func InitRoute(assetPath string) *web.Router {
	root := web.New(Base{}).
		Middleware((*Base).Init).
		Middleware(web.StaticMiddleware(assetPath)).
		Middleware((*Base).RequestLog)

	root.SubRouter(Index{}).
		Get("/", (*Index).Index).
		Get("/best", (*Index).Best).
		Get("/rank", (*Index).Rank)

	root.SubRouter(Channel{}).
		// 所有频道
		Get("/channels", (*Channel).All).
		// 频道页精选
		Get("/channel-(?P<channel>.+?)-best", (*Channel).Best).
		// 频道页二级分类视频列表
		Get("/channel-(?P<channel>.+?)-(?P<category>.+)", (*Channel).Category).
		// 频道页搜索视频列表
		Get("/channel-(?P<channel>.+?)", (*Channel).List)

	root.SubRouter(Video{}).
		Get("/video-(?P<vid>.+)", (*Video).Detail)

	return root
}
