package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"regexp"
	"time"
)

const (
	ViewKeyPattern = `viewkey=([0-9a-z]+)`
	ParamsPattern  = `<title>Free\ porn\ video\ \-\ (.*)\s*<\/title>(\s|\S)*so\.addVariable\('file','(\d+)'\);(\s|\S)*so\.addVariable\('max_vid','(\d+)'\);(\s|\S)*so\.addVariable\('seccode','([0-9a-f]+)'\);`
	Mp4UrlPattern  = `file=(http:\/\/.*)&domainUrl`
)

var (
	ViewKeyReg = regexp.MustCompile(ViewKeyPattern)
	ParamsReg  = regexp.MustCompile(ParamsPattern)
	Mp4UrlReg  = regexp.MustCompile(Mp4UrlPattern)
	client     = &http.Client{}

	file        *os.File
	viewData    = make([]*Params, 0)
	viewDataMap = make(map[string]int)
)

type Params struct {
	Is91Porn bool
	ViewKey  string
	Title    string
	Vid      string
	MaxVid   string
	SecCode  string
	Mp4URL   string
}

func ReadData() error {
	bufReader := bufio.NewReader(file)
	for {
		line, err := bufReader.ReadString('\n')
		if err == io.EOF {
			break
		}
		parts := strings.Split(line, "\t")
		if len(parts) > 1 {
			params := &Params{
				ViewKey:  strings.TrimSpace(parts[0]),
				Title:    strings.TrimSpace(parts[1]),
				Is91Porn: true,
			}
			if len(parts) > 2 {
				params.Is91Porn = false
				params.Mp4URL = parts[2]
			}
			viewDataMap[params.ViewKey] = len(viewData)
			viewData = append(viewData, params)
		}
	}
	return nil
}

func SaveData() error {
	if len(viewData) == 0 {
		return nil
	}

	err := file.Truncate(0)
	if err != nil {
		return err
	}
	_, err = file.Seek(0, 0)
	if err != nil {
		return err
	}

	for _, params := range viewData {
		var data = []string{params.ViewKey, params.Title}
		if !params.Is91Porn {
			data = append(data, params.Mp4URL)
		}
		line := strings.Join(data, "\t")
		line += "\n"
		_, err = file.Write([]byte(line))
		if err != nil {
			return err
		}
	}
	return nil
}

func GenIp() string {
	src := rand.NewSource(time.Now().Unix())
	r := rand.New(src)
	return fmt.Sprintf("%d.%d.%d.%d", 1+r.Intn(253), 1+r.Intn(253), 1+r.Intn(253), 1+r.Intn(253))
}

func init() {
	var err error
	file, err = os.OpenFile("data.txt", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}

	err = ReadData()
	if err != nil {
		panic(err)
	}
}

func HttpGet(urlStr, ip string) (body string, err error) {
	req, _ := http.NewRequest("GET", urlStr, nil)
	req.Header = http.Header{
		"Connection":    {"close"},
		"Cache-Control": {"max-age=0"},
		"Accept":        {"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8"},
		"User-Agent":    {"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.152 Safari/537.36"},
		"X-Forwarded-For": {ip},
		"Accept-Language": {"zh-CN,zh;q=0.8"},
	}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = fmt.Errorf("status code: %d", resp.StatusCode)
		return
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	body = string(data)
	return
}

func GetParams(viewKey, ip string) (params *Params, err error) {
	urlStr := "http://www.91porn.com/view_video.php?viewkey=" + viewKey
	body, err := HttpGet(urlStr, ip)
	if err != nil {
		return
	}
	matches := ParamsReg.FindStringSubmatch(body)
	if len(matches) != 8 {
		err = fmt.Errorf("body not match: %s", body)
		return
	}
	params = &Params{
		ViewKey: viewKey,
		Title:   strings.TrimSpace(matches[1]),
		Vid:     matches[3],
		MaxVid:  matches[5],
		SecCode: matches[7],
	}
	return
}

func GetVideoURL(params *Params, ip string) (u *url.URL, err error) {
	urlStr := fmt.Sprintf("http://www.91porn.com/getfile.php?VID=%s&mp4=0&seccode=%s&max_vid=%s", params.Vid, params.SecCode, params.MaxVid)
	body, err := HttpGet(urlStr, ip)
	if err != nil {
		return
	}
	matches := Mp4UrlReg.FindStringSubmatch(body)
	if len(matches) != 2 {
		err = fmt.Errorf("url:"+urlStr+", body not match: %s", body)
		return
	}
	u, err = url.Parse(matches[1])
	if err != nil {
		return
	}
	u.Path = path.Clean(u.Path)
	return
}

func GetViewKey(urlStr string) string {
	matches := ViewKeyReg.FindStringSubmatch(urlStr)
	if len(matches) != 2 {
		return ""
	}
	return matches[1]
}

func main() {
	http.HandleFunc("/av", func(rw http.ResponseWriter, req *http.Request) {
		err := req.ParseForm()
		if err != nil {
			rw.WriteHeader(500)
			rw.Write([]byte(err.Error()))
			return
		}
		var (
			n        = req.FormValue("n")
			clientIp = req.Header.Get("X-Forwarded-For")
			userAgent = req.Header.Get("User-Agent")
			now = time.Now().Format("2006-01-02 15:03:04")
		)
		rw.Write([]byte(`<!DOCTYPE HTML><html><head><meta charset="UTF-8"><meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0" /><title>精彩视频</title><body><h3>精彩视频</h3><hr/><ol style="line-height: 2em">`))
		for _, params := range viewData {
			rw.Write([]byte(`<li><a href="/av/play?id=` + params.ViewKey + `&amp;n=` + n + `">` + params.Title + `</a></li>`))
		}
		rw.Write([]byte(`</ol></body></html>`))
		fmt.Printf("%s IP: %s, UA: %s, GET /av, n: %s\n", now, clientIp,  userAgent, n)
	})

	http.HandleFunc("/av/add", func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`<!DOCTYPE HTML><html><head><meta charset="UTF-8"><meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0" /><title>91直达号</title><body><form action="/av/redirect" method="post" enctype="application/x-www-form-urlencoded"><textarea name="url" cols="50" rows="10"></textarea><br/><input type="submit" value="直达"/>&nbsp;&nbsp;<input type="reset" value="清空"/></form></body></html>`))
	})

	http.HandleFunc("/av/play", func(rw http.ResponseWriter, req *http.Request) {
		err := req.ParseForm()
		if err != nil {
			rw.WriteHeader(500)
			rw.Write([]byte(err.Error()))
			return
		}
		var (
			viewKey  = req.FormValue("id")
			clientIp = req.Header.Get("X-Forwarded-For")
			userAgent = req.Header.Get("User-Agent")
			now = time.Now().Format("2006-01-02 15:03:04")
			n        = req.FormValue("n")
			ip       = GenIp()
			params   *Params
		)

		index, ok := viewDataMap[viewKey]
		params = viewData[index]
		if !ok || (ok && params.Is91Porn) {
			var err error
			params, err = GetParams(viewKey, ip)
			if err != nil {
				rw.WriteHeader(500)
				rw.Write([]byte(err.Error()))
				return
			}

			u, err := GetVideoURL(params, ip)
			if err != nil {
				rw.WriteHeader(500)
				rw.Write([]byte(err.Error()))
				return
			}
			params.Mp4URL = u.String()

			if _, ok := viewDataMap[viewKey]; !ok {
				viewDataMap[viewKey] = len(viewData)
				viewData = append(viewData, params)
				SaveData()
			}
		}

		rw.Write([]byte(`<!DOCTYPE HTML><html><head><meta charset="UTF-8"><meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0" /><title>` + params.Title + `</title><body><h1 style="font-size: 16px">` + params.Title + `</h1><hr/><video  controls="controls" autoplay="autoplay" src="` + params.Mp4URL + `" style="width: 100%;height: 240px"></video><p><a href="` + params.Mp4URL + `">下载视频</a></p><h3>更多精彩视频</h3><hr/><ol style="line-height: 2em">`))

		for _, params := range viewData {
			rw.Write([]byte(`<li><a href="/av/play?id=` + params.ViewKey + `&amp;n=` + n + `">` + params.Title + `</a></li>`))
		}
		rw.Write([]byte(`</ol></body></html>`))

		fmt.Printf("%s IP: %s, UA: %s, GET /av/play?id=%s&n=%s\n", now, clientIp, userAgent, viewKey, n)
	})

	http.HandleFunc("/av/redirect", func(rw http.ResponseWriter, req *http.Request) {
		err := req.ParseForm()
		if err != nil {
			rw.WriteHeader(500)
			rw.Write([]byte(err.Error()))
			return
		}
		var (
			urlStr   = req.FormValue("url")
			viewKey  = GetViewKey(urlStr)
			clientIp = req.Header.Get("X-Forwarded-For")
		)
		if viewKey == "" {
			rw.WriteHeader(400)
			rw.Write([]byte("Bad Request"))
			return
		}

		rw.Header().Set("Location", "/av/play?id="+viewKey)
		rw.WriteHeader(302)

		fmt.Printf("IP: %s, POST /av/redirect url=%s\n", clientIp, urlStr)
	})

	fmt.Println("Start listen to :8090")
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		panic(err)
	}
}
