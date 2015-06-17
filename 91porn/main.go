package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"regexp"
)

const (
	VidPattern     = `so.addVariable('file','(\d+)');`
	MaxVidPattern  = `so.addVariable('max_vid','(\d+)');`
	SecCodePattern = `so.addVariable('seccode','([0-9a-f]+)');`
)

var (
	VidReg     = regexp.MustCompile(VidPattern)
	MaxVidReg  = regexp.MustCompile(MaxVidPattern)
	SecCodeReg = regexp.MustCompile(SecCodePattern)
	client     = &http.Client{}
)

type Params struct {
	Vid     string
	MaxVid  string
	SecCode string
}

func HttpGet(urlStr string) (body string, err error) {
	req, _ := http.NewRequest("GET", urlStr, nil)
	req.Header = http.Header{
		"Connection":      {"keep-alive"},
		"Cache-Control":   {"max-age=0"},
		"Accept":          {"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8"},
		"User-Agent":      {"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.152 Safari/537.36"},
		"Accept-Encoding": {"gzip, deflate, sdch"},
		"Accept-Language": {"zh-CN,zh;q=0.8"},
	}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	resp.Header
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	body = string(data)
	return
}

func GetParams(urlStr string) (params *Params, err error) {
	body, err := HttpGet(urlStr)
	if err != nil {
		return
	}
	url.URL.String()
	path.Clean()
	matches := VidReg.FindAllStringSubmatch(string(data))
	fmt.Printf("%#v", matches)
	return
}

func main() {
	http.HandleFunc("/91", func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`<!DOCTYPE HTML><html><head><meta charset="UTF-8"><meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0" /><title>91快达</title><body><form action="/91/redirect" method="post" enctype="application/x-www-form-urlencoded"><textarea name="url" rows="30"></textarea><input type="submit" value="直达"/></form></body></html>`))
	})
	http.HandleFunc("/91/redirect", func(rw http.ResponseWriter, req *http.Request) {
		err := req.ParseForm()
		if err != nil {
			rw.WriteHeader(500)
			rw.Write([]byte(err.Error()))
			return
		}
		params, err := GetParams(req.FormValue("url"))
		if err != nil {
			rw.WriteHeader(500)
			rw.Write([]byte(err.Error()))
			return
		}

		u, err := GetVideoURL(params)
		if err != nil {
			rw.WriteHeader(500)
			rw.Write([]byte(err.Error()))
			return
		}

		u.Path = path.Clean(u.Path)
		u.Host = "7xj0bm.com1.z0.glb.clouddn.com"

		rw.WriteHeader(302)
		rw.Header().Set("Location", u.String())
	})
}
