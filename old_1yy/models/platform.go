package models

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

var (
	httpClient = &http.Client{}
	uaFlag     = map[string][]string{
		"UC": []string{"UCBrowser", "UCWEB", "UC AppleWebKit"},
	}
)

func checkRedirectFunc(req *http.Request, _ []*http.Request) error {
	if req.URL.Host == "m.iqiyi.com" {
		return fmt.Errorf("PHONE")
	}
	return nil
}
func determineBrowser(ua string) string {
	for browser, flags := range uaFlag {
		for _, flag := range flags {
			if strings.Index(ua, flag) > -1 {
				return browser
			}
		}
	}
	return "OTHER"
}
func determinePlatformByIqiyi(ua string) (platform string, err error) {
	platform = "PC"
	req, err := http.NewRequest("HEAD", "http://www.iqiyi.com/", nil)
	if err != nil {
		fmt.Printf("创建HEAD请求www.iqiyi.com出错: %s", err.Error())
		return
	} else {
		req.Header.Add("User-Agent", ua)
		httpClient.CheckRedirect = checkRedirectFunc
		_, err = httpClient.Do(req)
		if err != nil {
			if strings.Index(err.Error(), "PHONE") > 0 {
				platform = "PHONE"
			} else {
				fmt.Printf("发起HEAD请求www.iqiyi.com出错: %s", err.Error())
			}
		}
	}
	return
}

//判断用户终端平台及浏览器品牌
func DetermineTerminal(ua string) (platform string, browser string) {
	platform, browser = "PHONE", "OTHER"
	md5hash := md5.New()
	_, err := io.WriteString(md5hash, ua)
	if err != nil {
		fmt.Printf("获取UA md5出错: %s", err.Error())
		return
	}
	uaMd5 := hex.EncodeToString(md5hash.Sum(nil))
	key := fmt.Sprintf("ua/md5/%s", uaMd5)
	data, exists := cache.Get(key)
	if exists {
		info := data.([2]string)
		platform, browser = info[0], info[1]
	} else {
		row := db.QueryRow("SELECT platform,browser FROM ua WHERE md5=?", uaMd5)
		err = row.Scan(&platform, &browser)
		if err == sql.ErrNoRows {
			browser = determineBrowser(ua)
			platform, err = determinePlatformByIqiyi(ua)
			if err == nil {
				_, err = db.Exec("INSERT INTO ua(md5,ua,platform,browser) VALUES(?,?,?,?)", uaMd5, ua, platform, browser)
				if err != nil {
					fmt.Printf("增加UA平台记录出错: %s", err.Error())
				}
			}
		} else if err != nil {
			fmt.Printf("查询UA表出错: %s", err.Error())
		}
		cache.Set(key, [2]string{platform, browser}, time.Hour*24)
	}
	return
}
