package controllers

import (
	"1yy/models"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/hoisie/web"
	"net/url"
	"strconv"
	"time"
)

var channels = map[string]string{
	"movie":   "电影",
	"tv":      "电视剧",
	"anime":   "动漫",
	"variety": "综艺",
}

func ApiResInfo(ctx *web.Context, status int, message string, data interface{}) {
	ctx.SetHeader("content-type", "text/json", true)
	apiRes := map[string]interface{}{
		"Status":  status,
		"Message": message,
		"Data":    data,
	}
	resBytes, err := json.Marshal(apiRes)
	if err != nil {
		ctx.WriteHeader(500)
		ctx.WriteString(`{"Status:500,"Message":"序列化JSON出错!"`)
		return
	}

	ctx.Write(resBytes)
}
func GenreVideoApi(ctx *web.Context, pGenre, pPageSize, pPage string) {
	page, _ := strconv.Atoi(pPage)
	pageSize, _ := strconv.Atoi(pPageSize)

	videos, total, err := models.GetGenreVideoList(pGenre, page, pageSize)
	if err != nil {
		ApiResInfo(ctx, 500, err.Error(), nil)
		return
	}
	ApiResInfo(ctx, 200, "OK", map[string]interface{}{
		"Page":   page,
		"Total":  total,
		"Videos": videos,
	})
}
func SearchVideoApi(ctx *web.Context, pQuery, pPageSize, pPage string) {
	page, _ := strconv.Atoi(pPage)
	pageSize, _ := strconv.Atoi(pPageSize)
	queries, err := url.ParseQuery(pQuery)
	if err != nil {
		ServerErrorPage(ctx, err)
		return
	}
	keyword, pageSize := queries.Get("q"), 18
	if len(keyword) == 0 {
		keyword = queries.Get("a")
	}

	if len(keyword) == 0 {
		ApiResInfo(ctx, 400, "请输入查询关键字", nil)
		return
	}
	videos, total, err := models.SearchVideoList(keyword, page, pageSize)
	if err != nil {
		ApiResInfo(ctx, 500, err.Error(), nil)
		return
	}
	ApiResInfo(ctx, 200, "OK", map[string]interface{}{
		"Page":   page,
		"Total":  total,
		"Videos": videos,
	})
}
func OpenVideoApi(ctx *web.Context, pChannel string) {
	channel := channels[pChannel]

	if len(channel) == 0 {
		ApiResInfo(ctx, 400, "不支持的视频频道", nil)
		return
	}
	playType := ctx.Params["play_type"]
	startTimeStr, endTimeStr := ctx.Params["start_time"], ctx.Params["end_time"]
	pageSizeStr, pageStr := ctx.Params["page_size"], ctx.Params["page"]

	//缺省结束时间为当前时间
	endTS := time.Now().Unix()
	if len(endTimeStr) > 0 {
		tmpTS, err := strconv.Atoi(endTimeStr)
		if err != nil {
			ApiResInfo(ctx, 400, "end_time必须为整型数值", nil)
			return
		}
		endTS = int64(tmpTS)
	}

	startTS := endTS - 86400
	if len(startTimeStr) > 0 {
		tmpTS, err := strconv.Atoi(startTimeStr)
		if err != nil {
			ApiResInfo(ctx, 400, "start_time必须为整型数值", nil)
			return
		}
		startTS = int64(tmpTS)
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if pageSize == 0 {
		pageSize = 10
	}

	page := 1
	if len(pageStr) > 0 {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			ApiResInfo(ctx, 400, "page必须为整型数值", nil)
			return
		}
	}

	videos, total, err := models.GetTypeVideoList(playType, channel, startTS, endTS, page, pageSize)
	if err != nil {
		ApiResInfo(ctx, 500, err.Error(), nil)
		return
	}
	ApiResInfo(ctx, 200, "OK", map[string]interface{}{
		"Total":  total,
		"Videos": videos,
	})
}

func SearchRedirectApi(ctx *web.Context) {
	page, keyword := ctx.Params["p"], ctx.Params["q"]
	if len(page) == 0 {
		page = "1"
	}
	if len(keyword) == 0 {
		dftKwCfg, err := models.GetConfig("DefaultKeyword")
		if err != nil {
			ctx.Redirect(302, "/")
			return
		}
		keyword = dftKwCfg.Value
	}
	redirect := fmt.Sprintf("/search/q=%s/%s.html", url.QueryEscape(keyword), page)
	ctx.Redirect(302, redirect)
}

func PlayStatApi(ctx *web.Context, pVid, pType, pSn string) {
	vid, _ := strconv.Atoi(pVid)
	models.PVStatistics(vid)
	gifStr := "R0lGODlhAQABAIAAAP///wAAACH5BAEAAAAALAAAAAABAAEAAAICRAEAOw=="
	data, err := base64.StdEncoding.DecodeString(gifStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	ctx.SetHeader("content-type", "image/gif", true)
	ctx.Write(data)
}

func FlushTplCacheApi(ctx *web.Context) {
	cache.Flush()
	ctx.WriteString("Flush template cache: OK!\n")
}

func FlushConfigCacheApi(ctx *web.Context) {
	models.FlushConfigCache()
	ctx.WriteString("Flush config cache: OK!\n")
}

func FlushModelCacheApi(ctx *web.Context) {
	models.FlushModelCache()
	ctx.WriteString("Flush model cache: OK\n")
}
