package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	jsonType = "application/json;charset=UTF-8"
)

type MemJar map[string][]*http.Cookie

func (j MemJar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	j[u.Host] = cookies
}

func (j MemJar) Cookies(u *url.URL) []*http.Cookie {
	return j[u.Host]
}

type Transport struct {
	client *http.Client
}

func NewTransport(client *http.Client) http.RoundTripper {
	return &Transport{
		client: client,
	}
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", "Android_4.3_yulong+Coolpad+8720L")
	req.Header.Set("Cookie2", "$Version=1")
	req.Header.Set("Accept-Encoding", "gzip")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("m-lng", "126.614407")
	req.Header.Set("m-lat", "42.080081")
	req.Header.Set("m-nw", "wifi")
	req.Header.Set("m-iv", "3.1.2")
	req.Header.Set("m-lt", "2")
	req.Header.Set("m-ct", "1")
	req.Header.Set("m-cv", "20")
	req.Header.Set("m-cw", "720")
	req.Header.Set("m-ch", "1280")

	return t.client.Do(req)
}

type Result struct {
	Code      int    `json:"Code"`
	IsSuccess bool   `json:"IsSuccess"`
	Desc      string `json:"Desc"`
}
type Advert struct {
	Id                    int  `json:"Id"`
	IsPublicServiceAdvert bool `json:"IsPublicServiceAdvert"`
}
type CategoryAdsResult struct {
	Result
	Data []struct {
		Adverts []Advert
	} `json:"Data"`
}

type IndexResult struct {
	Result
	Data struct {
		RecommendProductList []Advert `json:"RecommendProductList"`
	} `json:"Data"`
}

type AdvertResult struct {
	Result
	Id          int  `json:"Id"`
	IsAllowEarn bool `json:"IsAllowEarn"`
}

type GeneratedResult struct {
	Result
	Data int `json:"Data"`
}

type SilverAdvert struct {
	Imei     string `json:"Imei"`
	UserName string `json:"UserName"`
	Password string `json:"Password"`
	host     string
	client   *http.Client
}

func NewSilverAdvert(host, imei, username, password string) *SilverAdvert {
	client := &http.Client{
		Transport: NewTransport(&http.Client{}),
		Jar:       make(MemJar),
	}

	return &SilverAdvert{
		Imei:     imei,
		UserName: username,
		Password: password,
		host:     host,
		client:   client,
	}

}
func (s *SilverAdvert) Login() (err error) {
	buf := bytes.NewBuffer(nil)
	encoder := json.NewEncoder(buf)
	err = encoder.Encode(s)
	if err != nil {
		return
	}

	resp, err := s.client.Post(s.host+"/api/Auth/Login", jsonType, buf)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	result := &Result{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(result)
	if err != nil {
		return
	}
	if !result.IsSuccess {
		err = errors.New(result.Desc)
	}
	return
}

func (s *SilverAdvert) IndexAdverts() (adverts []Advert, err error) {
	resp, err := s.client.Get(s.host + "/api/CustomerHome/Index")
	if err != nil {
		return
	}
	defer resp.Body.Close()

	result := &IndexResult{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&result)
	if err != nil {
		return
	}
	adverts = result.Data.RecommendProductList
	return
}

func (s *SilverAdvert) PullCategoryAds(categoryId int) (adverts []Advert, err error) {
	r := strings.NewReader(`{"CategoryIds":[` + strconv.Itoa(categoryId) + `]}`)
	resp, err := s.client.Post(s.host+"/api/SilverAdvert/Pull", jsonType, r)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	result := &CategoryAdsResult{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(result)
	if err != nil {
		return
	}
	if !result.IsSuccess {
		err = errors.New(result.Desc)
		return
	}

	for _, data := range result.Data {
		adverts = append(adverts, data.Adverts...)
	}

	return
}

func (s *SilverAdvert) GetAdvertDetail(adverId int) (result *AdvertResult, err error) {
	path := "/api/SilverAdvert/Detail?Id=" + strconv.Itoa(adverId)
	resp, err := s.client.Get(s.host + path)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	result = &AdvertResult{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(result)
	if err != nil {
		return
	}
	if !result.IsSuccess {
		err = errors.New(result.Desc)
		return
	}

	return
}

func (s *SilverAdvert) GeneratedIntegral(adverId int) (earn int, err error) {
	r := strings.NewReader(`{"Id":` + strconv.Itoa(adverId) + `,"IsGame":true}`)
	resp, err := s.client.Post(s.host+"/api/SilverAdvert/GeneratedIntegral", jsonType, r)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	result := &GeneratedResult{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(result)
	if err != nil {
		return
	}
	if !result.IsSuccess {
		err = errors.New(result.Desc)
	}

	earn = result.Data
	return
}
