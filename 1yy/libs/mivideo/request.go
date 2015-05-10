package mivideo

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"path"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const UserAgent = "Dalvik/2.1.0 (Linux; U; Android 5.0.2; MI 2S MIUI/5.4.17)"

type Request struct {
	host        string
	api         string
	path        string
	key         string
	token       string
	paramKeys   []string
	paramValues map[string]string
	client      *http.Client
}

func (r *Request) AddCommonParam() {
	r.AddParam("_locale", "zh_CN")
	r.AddParam("_res", "hd720")
	r.AddParam("_devid", "863583023262760")
	r.AddParam("_md5", "")
	r.AddParam("_model", "MI+2S")
	r.AddParam("_miuiver", "5.4.17")
	r.AddParam("_nonce", rand.Int())
	r.AddParam("_appver", "2015050890")
	r.AddParam("_ts", time.Now().Unix())
	r.AddParam("_ver", "2015.05.08-小米视频")
	r.AddParam("_devtype", 1)
	r.AddParam("_cam", "5e3a40cde6af860a24f1e091510fd6a9")
	r.AddParam("_diordna", "5ade50edbef0d2e2c49c5748e3d9e5a3")
	r.AddParam("token", r.token)

	opaque := r.Opaque()
	r.AddParam("opaque", opaque)
}

func (r *Request) AddParam(key string, value interface{}) {
	var str string
	if key != "" {
		val := reflect.ValueOf(value)
		switch val.Kind() {
		case reflect.String:
			str = reflect.ValueOf(value).String()
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			str = strconv.FormatInt(val.Int(), 10)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			str = strconv.FormatUint(val.Uint(), 10)
		case reflect.Float32, reflect.Float64:
			str = strconv.FormatFloat(val.Float(), 'f', -1, 64)
		}

		r.paramKeys = append(r.paramKeys, key)
		r.paramValues[key] = key + "=" + url.QueryEscape(str)
	}
}

func (r *Request) Opaque() string {
	str := path.Join(r.api, r.path) + "?" + r.Query()
	buf := bytes.NewBufferString(str)
	h := hmac.New(sha1.New, []byte(r.key))
	io.Copy(h, buf)
	return hex.EncodeToString(h.Sum(nil))
}

func (r *Request) Query() string {
	params := make([]string, len(r.paramValues))
	for i, paramKey := range r.paramKeys {
		if param, ok := r.paramValues[paramKey]; ok {
			params[i] = param
		}
	}
	return strings.Join(params, "&")
}

func (r *Request) Get(fn func(io.Reader) error) (err error) {
	reqURL := r.host + path.Join(r.api, r.path) + "?" + r.Query()

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return
	}
	req.Header.Set("User-Agent", UserAgent)

	res, err := r.client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	err = fn(res.Body)
	return
}

func (r *Request) Post(fn func(io.Reader) error) (err error) {
	var (
		reqURL  = r.host + path.Join(r.api, r.path)
		bodyBuf = bytes.NewBufferString(r.Query())
	)

	req, err := http.NewRequest("GET", reqURL, bodyBuf)
	if err != nil {
		return
	}
	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := r.client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	err = fn(res.Body)
	return
}
