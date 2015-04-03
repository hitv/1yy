package controllers

import (
	"1yy/models"
	"fmt"
	"github.com/hoisie/web"
	gocache "github.com/pmylund/go-cache"
	"github.com/ziutek/kview"
	"math"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var (
	cache *gocache.Cache
	utils = map[string]interface{}{
		"QueryEscape": url.QueryEscape,
		"PosterUrl":   models.PosterUrl,
		"StaticUrl":   models.StaticUrl,
		"ActorLink":   actorLink,
		"Now": func() string {
			return time.Now().Format("2006-01-02 15:04:05")
		},
	}
)

func actorLink(str string) string {
	if len(str) == 0 {
		return ""
	}
	actors := make([]string, 0, 20)
	for _, actor := range strings.Split(str, ",") {
		if len(actor) > 0 {
			actor = fmt.Sprintf(`<a href="/search/q=%s/1.html" title="点击查看%s演的影片">%s</a>`, url.QueryEscape(actor), actor, actor)
			actors = append(actors, actor)
		}
	}
	return strings.Join(actors, ",")
}
func init() {
	cache = gocache.New(-1, -1)
}

//获取通用配置
func getCommonData() (string, []string, [2]int, error) {
	var count [2]int
	dftKwCfg, err := models.GetConfig("DefaultKeyword")
	if err != nil {
		return "", nil, count, err
	}
	recKwCfg, err := models.GetConfig("RecommendKeywords")
	if err != nil {
		return "", nil, count, err
	}
	recKeywords := strings.Split(recKwCfg.Value, "|")
	count, err = models.GetVideoCount()
	if err != nil {
		return "", nil, count, err
	}
	return dftKwCfg.Value, recKeywords, count, nil
}
func getView(platform, tplName string) *kview.KView {
	var view *kview.KView
	platform = strings.ToLower(platform)
	key := fmt.Sprintf("template/%s/%s", platform, tplName)
	data, exists := cache.Get(key)
	if exists {
		view = data.(*kview.KView)
	} else {
		tplName = fmt.Sprintf("%s/%s.tpl", platform, tplName)
		view = kview.New(tplName)
		cache.Set(key, view, 0)

	}
	return view
}
func pageContent(pageSize, page, total, num int) (int, []int) {
	pageNumbers := make([]int, 0, num*2+1)
	pageCount := int(math.Ceil(float64(total) / float64(pageSize)))
	min := page - num
	max := page + num
	if max > pageCount {
		max = pageCount
	}
	if min < 1 {
		min = 1
	}
	for i := min; i <= max; i++ {
		pageNumbers = append(pageNumbers, i)
	}
	return pageCount, pageNumbers
}
func NotFoundPage(ctx *web.Context, message string) {
	ctx.WriteHeader(404)
	ctx.WriteString(message)
}
func ServerErrorPage(ctx *web.Context, err error) {
	ctx.WriteHeader(500)
	ctx.WriteString(err.Error())
}

func IndexPage(ctx *web.Context) {
	ua := ctx.Request.Header.Get("User-Agent")
	platform, browser := models.DetermineTerminal(ua)

	hotVideos, err := models.GetHotVideos("", 6)
	if err != nil {
		ServerErrorPage(ctx, err)
		return
	}
	rankVideos, err := models.GetRankVideos([]string{"电影", "电视剧"}, 13)
	if err != nil {
		ServerErrorPage(ctx, err)
		return
	}

	videoNum := 6
	if platform == "PC" {
		videoNum = 30
	}

	channelVideos, err := models.GetIndexChannelVideos(videoNum)
	if err != nil {
		ServerErrorPage(ctx, err)
		return
	}
	activeItem, navItems, err := models.GetNavGenres("index")
	if err != nil {
		ServerErrorPage(ctx, err)
		return
	}
	dftKeyword, recKeywords, count, err := getCommonData()
	if err != nil {
		ServerErrorPage(ctx, err)
		return
	}
	layoutTpl := getView(platform, "layout")
	layoutTpl.Div("headTpl", getView(platform, "head/index"))
	layoutTpl.Div("contentTpl", getView(platform, "content/index"))
	layoutTpl.Exec(ctx, map[string]interface{}{
		"head": map[string]interface{}{},
		"search": map[string]interface{}{
			"keyword":        "",
			"defaultKeyword": dftKeyword,
			"recKeywords":    recKeywords,
		},
		"count": map[string]interface{}{
			"todayUpdate": count[0],
			"total":       count[1],
		},
		"nav": map[string]interface{}{
			"items":       navItems,
			"activeItem":  activeItem,
			"activeGenre": "index",
		},
		"content": map[string]interface{}{
			"hotVideos":     hotVideos,
			"channelVideos": channelVideos,
			"rankVideos":    rankVideos,
			"utils":         utils,
		},
		"browser":        browser,
		"showShare":      true,
		"showFriendLink": true,
		"utils":          utils,
	})
}

func SearchVideoPage(ctx *web.Context, pQuery, pPage string) {
	ua := ctx.Request.Header.Get("User-Agent")
	platform, browser := models.DetermineTerminal(ua)

	dftKeyword, recKeywords, count, err := getCommonData()
	if err != nil {
		ServerErrorPage(ctx, err)
		return
	}

	queries, err := url.ParseQuery(pQuery)
	if err != nil {
		ServerErrorPage(ctx, err)
		return
	}
	page, _ := strconv.Atoi(pPage)
	keyword, pageSize := strings.Trim(queries.Get("q"), " "), 18
	if len(keyword) == 0 {
		keyword = queries.Get("a")
	}
	if platform == "PC" {
		pageSize = 10
	}

	activeItem, navItems, err := models.GetNavGenres("")
	if err != nil {
		ServerErrorPage(ctx, err)
		return
	}

	videos, videoCount, err := models.SearchVideoList(keyword, page, pageSize)
	if err != nil {
		ServerErrorPage(ctx, err)
		return
	}

	hotVideos, err := models.GetHotVideos("", 30)
	if err != nil {
		ServerErrorPage(ctx, err)
		return
	}

	pageCount, pageNumbers := pageContent(pageSize, page, videoCount, 4)

	layoutTpl := getView(platform, "layout")
	layoutTpl.Div("headTpl", getView(platform, "head/search"))
	layoutTpl.Div("contentTpl", getView(platform, "content/search"))
	layoutTpl.Exec(ctx, &map[string]interface{}{
		"head": map[string]interface{}{
			"keyword": keyword,
			"page":    page,
		},
		"search": map[string]interface{}{
			"keyword":        keyword,
			"defaultKeyword": dftKeyword,
			"recKeywords":    recKeywords,
		},
		"count": map[string]interface{}{
			"todayUpdate": count[0],
			"total":       count[1],
		},
		"nav": map[string]interface{}{
			"items":       navItems,
			"activeItem":  activeItem,
			"activeGenre": "",
		},
		"content": map[string]interface{}{
			"activeNavItem":  activeItem,
			"currentKeyword": keyword,
			"pager": map[string]interface{}{
				"recordCount": videoCount,
				"currentPage": page,
				"pageSize":    pageSize,
				"pageCount":   pageCount,
				"pageNumbers": pageNumbers,
				"hasNextPage": (page < pageCount),
				"hasPrevPage": (page > 1),
				"prevPage":    page - 1,
				"nextPage":    page + 1,
			},
			"videos":    videos,
			"hotVideos": hotVideos,
			"utils":     utils,
		},
		"browser": browser,
		"utils":   utils,
	})
}

func VideoDetailPage(ctx *web.Context, pVid string) {
	ua := ctx.Request.Header.Get("User-Agent")
	platform, browser := models.DetermineTerminal(ua)

	vid, _ := strconv.Atoi(pVid)
	video, err := models.GetVideo(uint64(vid))
	if err != nil {
		ServerErrorPage(ctx, err)
		return
	}
	if video == nil {
		NotFoundPage(ctx, `Sorry!该视频已下线 <a href="/"><<返回首页</a>`)
		return
	}
	if video.Status == "MERGED" && video.MergeVid > 0 {
		ctx.Redirect(301, fmt.Sprintf("%d.html", video.MergeVid))
		return
	}
	activeItem, navItems, err := models.GetNavGenres(video.Channel)
	if err != nil {
		ServerErrorPage(ctx, err)
		return
	}
	urls, err := models.GetSortedVideoTypeUrls(vid, "qvod")
	if err != nil {
		ServerErrorPage(ctx, err)
		return
	}

	dftKeyword, recKeywords, count, err := getCommonData()
	if err != nil {
		ServerErrorPage(ctx, err)
		return
	}

	recNum := 8
	if platform == "PHONE" {
		recNum = 9
	}
	recVideos, err := models.GetRecommendVideos(video, recNum)
	if err != nil {
		ServerErrorPage(ctx, err)
		return
	}

	layoutTpl := getView(platform, "layout")
	layoutTpl.Div("headTpl", getView(platform, "head/detail"))
	layoutTpl.Div("contentTpl", getView(platform, "content/detail"))
	layoutTpl.Exec(ctx, map[string]interface{}{
		"head": map[string]interface{}{
			"video": video,
		},
		"search": map[string]interface{}{
			"keyword":        "",
			"defaultKeyword": dftKeyword,
			"recKeywords":    recKeywords,
		},
		"count": map[string]interface{}{
			"todayUpdate": count[0],
			"total":       count[1],
		},
		"nav": map[string]interface{}{
			"items":       navItems,
			"activeItem":  activeItem,
			"activeGenre": video.Genres,
		},
		"content": map[string]interface{}{
			"video":       video,
			"typeUrlList": urls.Slice,
			"channel":     activeItem,
			"recVideos":   recVideos,
			"utils":       utils,
		},
		"browser": browser,
		"utils":   utils,
	})
}

func GenreVideoPage(ctx *web.Context, pGenre, pPage string) {
	ua := ctx.Request.Header.Get("User-Agent")
	platform, browser := models.DetermineTerminal(ua)

	page, _ := strconv.Atoi(pPage)
	pageSize := 18
	if platform == "PC" {
		pageSize = 10
	}
	genre, err := models.GetGenre(pGenre)
	if err != nil {
		ServerErrorPage(ctx, err)
		return
	}
	if genre == nil {
		NotFoundPage(ctx, `Sorry!没找到该类型 <a href="/"><<返回首页</a>`)
		return
	}
	activeItem, navItems, err := models.GetNavGenres(pGenre)
	if err != nil {
		ServerErrorPage(ctx, err)
		return
	}
	videos, videoCount, err := models.GetGenreVideoList(pGenre, page, pageSize)
	if err != nil {
		ServerErrorPage(ctx, err)
		return
	}

	dftKeyword, recKeywords, updateCount, err := getCommonData()
	if err != nil {
		ServerErrorPage(ctx, err)
		return
	}

	hotVideos, err := models.GetHotVideos(genre.Name, 30)
	if err != nil {
		ServerErrorPage(ctx, err)
		return
	}

	pageCount, pageNumbers := pageContent(pageSize, page, videoCount, 4)

	layoutTpl := getView(platform, "layout")
	layoutTpl.Div("contentTpl", getView(platform, "content/genre"))
	layoutTpl.Div("headTpl", getView(platform, "head/genre"))
	layoutTpl.Exec(ctx, map[string]interface{}{
		"head": map[string]interface{}{
			"activeNavItem": activeItem,
			"genre":         genre,
			"page":          page,
		},
		"search": map[string]interface{}{
			"keyword":        "",
			"defaultKeyword": dftKeyword,
			"recKeywords":    recKeywords,
		},
		"count": map[string]interface{}{
			"todayUpdate": updateCount[0],
			"total":       updateCount[1],
		},
		"nav": map[string]interface{}{
			"items":       navItems,
			"activeItem":  activeItem,
			"activeGenre": pGenre,
		},
		"content": map[string]interface{}{
			"genre":         genre,
			"activeNavItem": activeItem,
			"pager": map[string]interface{}{
				"recordCount": videoCount,
				"currentPage": page,
				"pageSize":    pageSize,
				"pageCount":   pageCount,
				"pageNumbers": pageNumbers,
				"hasNextPage": (page < pageCount),
				"hasPrevPage": (page > 1),
				"prevPage":    page - 1,
				"nextPage":    page + 1,
			},
			"videos":    videos,
			"hotVideos": hotVideos,
			"showMore":  (page < pageSize),
			"utils":     utils,
		},
		"browser": browser,
		"utils":   utils,
	})
}

func VideoPlayPage(ctx *web.Context, pVid, pType, pSn string) {
	ua := ctx.Request.Header.Get("User-Agent")
	platform, browser := models.DetermineTerminal(ua)

	vid, _ := strconv.Atoi(pVid)
	sn, _ := strconv.Atoi(pSn)

	video, err := models.GetVideo(uint64(vid))
	if err != nil {
		ServerErrorPage(ctx, err)
		return
	}
	if video == nil {
		NotFoundPage(ctx, `Sorry!该视频已下线 <a href="/"><<返回首页</a>`)
		return
	}
	activeItem, navItems, err := models.GetNavGenres("")
	if err != nil {
		ServerErrorPage(ctx, err)
		return
	}
	urls, err := models.GetSortedVideoTypeUrls(vid, pType)
	if err != nil {
		ServerErrorPage(ctx, err)
		return
	}

	if urls.Map[pType] == nil {
		msg := fmt.Sprintf("Sorry!视频播放出错 <a href=\"/detail/%d.html\"><<返回</a>", vid)
		NotFoundPage(ctx, msg)
		return
	}

	currentUrl := urls.Map[pType].Map[int64(sn)]
	if currentUrl == nil {
		NotFoundPage(ctx, `Sorry!没找到该剧集 <a href="/"><<返回首页</a>`)
		return
	}

	dftKeyword, recKeywords, count, err := getCommonData()
	if err != nil {
		ServerErrorPage(ctx, err)
		return
	}

	recNum := 8
	if platform == "PHONE" {
		recNum = 9
	}
	recVideos, err := models.GetRecommendVideos(video, recNum)
	if err != nil {
		ServerErrorPage(ctx, err)
		return
	}

	layoutTpl := getView(platform, "layout")
	layoutTpl.Div("headTpl", getView(platform, "head/play"))
	layoutTpl.Div("contentTpl", getView(platform, "content/play"))
	layoutTpl.Exec(ctx, map[string]interface{}{
		"head": map[string]interface{}{
			"video":      video,
			"currentUrl": currentUrl,
		},
		"search": map[string]interface{}{
			"keyword":        "",
			"defaultKeyword": dftKeyword,
			"recKeywords":    recKeywords,
		},
		"count": map[string]interface{}{
			"todayUpdate": count[0],
			"total":       count[1],
		},
		"nav": map[string]interface{}{
			"items":       navItems,
			"activeItem":  activeItem,
			"activeGenre": video.Channel,
		},
		"content": map[string]interface{}{
			"video":       video,
			"typeUrlList": urls.Slice,
			"currentUrl":  currentUrl,
			"recVideos":   recVideos,
			"utils":       utils,
		},
		"browser": browser,
		"utils":   utils,
	})
}

func SitemapPage(ctx *web.Context, pPage string) {
	page, _ := strconv.Atoi(pPage)
	videos, err := models.GetSitemapVideos(page, 100)
	if err != nil {
		ServerErrorPage(ctx, err)
		return
	}
	sitemapTpl := getView("pc", "sitemap")
	mTimeFmt := "2006-01-02T15:04:05+08:00"
	ctx.SetHeader("Content-type", "text/xml", true)
	sitemapTpl.Exec(ctx, map[string]interface{}{
		"page":          page,
		"sitemapVideos": videos,
		"timeFmt": map[string]interface{}{
			"Now": func() string {
				return time.Now().Format(mTimeFmt)
			},
			"Format": func(timeStr string) string {
				mTime, err := time.Parse("2006-01-02 15:04:05", timeStr)
				if err != nil {
					mTime = time.Now()
				}
				return mTime.Format(mTimeFmt)
			},
		},
	})
}

func BaiduSubmitPage(ctx *web.Context, pPage string) {
	page, _ := strconv.Atoi(pPage)
	videos, err := models.GetSitemapVideos(page, 50)
	if err != nil {
		ServerErrorPage(ctx, err)
		return
	}
	sitemapTpl := getView("pc", "baidusubmit")
	mTimeFmt := "2006-01-02T15:04:05+08:00"
	ctx.SetHeader("Content-type", "text/xml", true)
	sitemapTpl.Exec(ctx, map[string]interface{}{
		"page":          page,
		"sitemapVideos": videos,
		"timeFmt": map[string]interface{}{
			"Now": func() string {
				return time.Now().Format(mTimeFmt)
			},
			"Format": func(timeStr string) string {
				mTime, err := time.Parse("2006-01-02 15:04:05", timeStr)
				if err != nil {
					mTime = time.Now()
				}
				return mTime.Format(mTimeFmt)
			},
		},
	})
}
