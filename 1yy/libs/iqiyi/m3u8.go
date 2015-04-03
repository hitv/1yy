package iqiyi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
)

var (
	m3u8PartReg = regexp.MustCompile(``)
)

type m3u8Resp struct {
	Tvs struct {
		Tv struct {
			Vid    string `json:"vid"`
			Length int    `json:"len"`
			Res    []struct {
				Vid    string `jsno:"vid"`
				Length int    `json:"len"`
				Type   string `json:"t"`
			} `json:"res"`
			TsRes []struct {
				TsURL  string `json:"_tsurl"`
				Length int    `json:"len"`
				Type   string `json:"t"`
			}
		} `json:"0"`
	} `json:"tv"`
}

var (
	tr     = NewTransport("boW$%gqmHh-8JF[P", "10020202ddf238a3ed4b7fbac0e1c989", "4.9.1", 1001479717)
	client = http.Client{
		Transport: tr,
	}
)

func getM3u8URL() {

}
func GetM3u8(playURL string) (str string, err error) {
	var (
		resp1, resp2 *http.Response
		decoder      *json.Decoder
	)

	infoURL := "http://iface2.iqiyi.com/php/xyz/entry/nebula.php?key=10020202ddf238a3ed4b7fbac0e1c989&version=4.9.1&os=4.4.4&ua=MI+2S&network=1&resolution=1280*720&udid=863738029982338&openudid=5cd436af4856be7f&ppid=&uniqid=2f904dae9c18a3fa0ff466bf848b17f1&device_id=863738029982338&cpu=1728000&idfv=CAF252BCD263CB1429B645E85B36D08D&platform=GPhone_sdk_xiaomi&block=0&w=1&compat=1&other=1&v5=1&ad_str=1&api=1.1.2&ts=128%2C4%2C8%2C16&playurl=" + url.QueryEscape(playURL) + "&user_res=0&ad=2&js=1&vs=0&vt=0&xbm=0&x=0&y=3&z=0&cts=1426316029&lts=0&wts=128%2C4%2C8%2C16&wtsh=-1&v_m=2.2_005&qyid=imei"
	resp1, err = client.Get(infoURL)
	if err != nil {
		return
	}
	defer resp1.Body.Close()

	m3u8RespObj := &m3u8Resp{}

	decoder = json.NewDecoder(resp1.Body)
	err = decoder.Decode(m3u8RespObj)
	if err != nil {
		return
	}

	resp2, err = client.Get(m3u8RespObj.Tvs.Tv.Vid)
	if err != nil {
		return
	}
	defer resp2.Body.Close()

	body, _ := ioutil.ReadAll(resp2.Body)
	str = string(body)
	return
}
