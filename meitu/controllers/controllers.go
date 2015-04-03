package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/hoisie/web"
	gocache "github.com/pmylund/go-cache"
	"github.com/ziutek/kview"
	"meitu/models"
	"strconv"
	"time"
)

var cache *gocache.Cache

const (
	CATEGORY_PAGE_SIZE = "CATEGORY_PAGE_SIZE"
)

type apiRes struct {
	Status  int
	Message string
	Data    interface{}
}

func apiError(ctx *web.Context, status int, msg string) {
	ctx.ContentType("text/json")
	jsonBytes, err := json.Marshal(apiRes{
		Status:  status,
		Message: msg,
		Data:    nil,
	})
	if err != nil {
		fmt.Fprint(ctx, `{Status: 500, Message: "服务器错误"}`)
		return
	}
	ctx.Write(jsonBytes)
}
func pageError(ctx *web.Context, msg string) {
	ctx.ContentType("text/html")
	fmt.Fprintf(ctx, `服务器出错: %s&nbsp;&nbsp;<a href="/">返回首页</a>`, msg)
}
func apiEcho(ctx *web.Context, data interface{}) {
	jsonBytes, err := json.Marshal(apiRes{
		Status:  200,
		Message: "OK",
		Data:    data,
	})
	if err != nil {
		fmt.Fprint(ctx, `{Status: 500, Message: "服务器错误"}`)
		return
	}
	ctx.Write(jsonBytes)
}
func getView(tpl string) (view *kview.KView) {
	key := fmt.Sprintf("tpl/%s", tpl)
	data, exists := cache.Get(key)
	if exists {
		view = data.(*kview.KView)
		return
	}
	view = kview.New(tpl + ".tpl")
	cache.Set(key, view, time.Minute*5)
	return
}
func init() {
	cache = gocache.New(time.Minute*5, time.Minute)
}
func CategoryPage(ctx *web.Context, pCategoryId string) {
	categoryId, _ := strconv.Atoi(pCategoryId)
	categories, err := models.GetCategories()
	if err != nil {
		pageError(ctx, err.Error())
		return
	}
	view := getView("category")
	view.Exec(ctx, map[string]interface{}{
		"categoryId": categoryId,
		"categories": categories,
	})
}
func AlbumPage(ctx *web.Context, pAlbumId string) {
	view := getView("album")
	view.Exec(ctx, map[string]string{"albumId": pAlbumId})
}
func CategoryAlbumsApi(ctx *web.Context, pCategoryId, pPage string) {
	categoryId, _ := strconv.Atoi(pCategoryId)
	page, _ := strconv.Atoi(pPage)
	pageSizeConf, err := models.GetConfig(CATEGORY_PAGE_SIZE)
	if err != nil {
		apiError(ctx, 500, err.Error())
		return
	}
	pageSize, err := strconv.Atoi(pageSizeConf)
	if err != nil {
		apiError(ctx, 500, err.Error())
		return
	}
	albums, err := models.GetAlbums(categoryId, pageSize, page)
	if err != nil {
		apiError(ctx, 500, err.Error())
		return
	}
	apiEcho(ctx, albums)
}
func AlbumImagesApi(ctx *web.Context, pAlbumId string) {
	albumId, _ := strconv.Atoi(pAlbumId)
	album, err := models.GetAlbum(albumId)
	if err != nil {
		apiError(ctx, 500, err.Error())
		return
	}

	apiEcho(ctx, album)
}
