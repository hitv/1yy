package main

import (
	"1yy/controllers"
	"1yy/models"
	"bitbucket.org/chai2010/pwdgen/ini"
	"flag"
	"github.com/hoisie/web"
	"runtime"
)

var port string

func init() {
	confFile := flag.String("conf", "./config.ini", "配置文件路径")
	flag.Parse()

	dict, err := ini.LoadFile(*confFile)
	if err != nil {
		panic(err)
	}
	err = models.Init(dict)
	if err != nil {
		panic(err)
	}

	port, err = dict.GetString("common", "listen")
	if err != nil {
		panic(err)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	web.Get("^/$", controllers.IndexPage)
	//详情页
	web.Get(`^/detail/(\d+)\.html$`, controllers.VideoDetailPage)
	//播放页
	web.Get(`^/play/(\d+)-(bdhd|qvod)-(\d+)\.html$`, controllers.VideoPlayPage)
	//频道页
	web.Get(`^/genre/(\S+)/(\d+)\.html$`, controllers.GenreVideoPage)
	//频道页分页数据
	web.Get(`^/genre/(\S+)/(\d+)/(\d+)\.json$`, controllers.GenreVideoApi)
	//搜索跳转接口
	web.Get(`^/search$`, controllers.SearchRedirectApi)
	//搜索结果页
	web.Get(`^/search/(\S+)/(\d+)\.html$`, controllers.SearchVideoPage)
	//站点地图
	web.Get(`^/sitemap/(\d+)\.xml$`, controllers.SitemapPage)
	//百度submit
	web.Get(`^/baidusubmit/(\d+)\.xml$`, controllers.BaiduSubmitPage)
	//搜索页分页数据
	web.Get(`^/search/(.+)/(\d+)/(\d+)\.json$`, controllers.SearchVideoApi)
	//同步数据API
	web.Get(`^/api/(.+).json$`, controllers.OpenVideoApi)
	//用于统计播放量的API
	web.Get(`^/stat/play/(\d+)/(\S+)/(\d+).gif$`, controllers.PlayStatApi)
	web.Delete(`^/admin/cache/tpl`, controllers.FlushTplCacheApi)
	web.Delete(`^/admin/cache/config`, controllers.FlushConfigCacheApi)
	web.Delete(`^/admin/cache/model`, controllers.FlushModelCacheApi)
	web.Run(port)
}
