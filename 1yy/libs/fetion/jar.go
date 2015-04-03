package fetion

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
	"wx-novel/app/models/cache"
)

type Jar struct {
	Cache      cache.Cache
	Expiration time.Duration
	KeyPrefix  string
}

func (jar *Jar) getJarKey(u *url.URL) string {
	return fmt.Sprintf("%s_jar_%s", jar.KeyPrefix, u.Host)
}
func (jar *Jar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	var (
		existsCookies []*http.Cookie
		jarKey        = jar.getJarKey(u)
	)
	//从缓存中获取站点的cookies
	err := jar.Cache.Get(jarKey, &existsCookies)
	if err != nil {
		log.Printf("Get cookie error: %s\n", err)
	}
	//建立cookie名的反向索引
	cookieMap := make(map[string]int)
	for i, cookie := range cookies {
		cookieMap[cookie.Name] = i
	}
	//找出原cookies中不存在新cookies里的cookie，追加到新cookies里
	for _, cookie := range existsCookies {
		if _, exists := cookieMap[cookie.Name]; !exists {
			cookies = append(cookies, cookie)
		}
	}
	//将cookies存入缓存中
	err = jar.Cache.Set(jarKey, cookies, jar.Expiration)
	if err != nil {
		log.Printf("Set cookie error: %s\n", err)
	}
}

func (jar *Jar) Cookies(u *url.URL) (cookies []*http.Cookie) {
	//从缓存中获取站点的cookies
	jarKey := jar.getJarKey(u)
	err := jar.Cache.Get(jarKey, &cookies)
	if err != nil {
		log.Printf("Get cookie error: %s\n", err)
	}
	return cookies
}
