package fetch

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (
	MI_REQUEST_KEY   = "581582928c881b42eedce96331bff5d3"
	MI_REQUEST_TOKEN = "0f9dfa001cba164d7bda671649c50abf"
	MI_HOST          = "http://mobile.duokanbox.com"
)

type RequestURL struct {
	host   string
	path   string
	params map[string]string
	keys   []string
}

func NewRequestURL(host, path string) *RequestURL {
	return &RequestURL{
		host:   host,
		path:   path,
		params: make(map[string]string),
		keys:   make([]string, 0),
	}
}

func (p *RequestURL) AddParam(key string, value interface{}) {
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

		p.keys = append(p.keys, key)
		p.params[key] = key + "=" + url.QueryEscape(str)
	}
}

func (p *RequestURL) Opaque(key, suffix string) string {
	str := p.path + "?" + p.Query() + suffix
	buf := bytes.NewBufferString(str)
	h := hmac.New(sha1.New, []byte(key))
	io.Copy(h, buf)
	return hex.EncodeToString(h.Sum(nil))
}

func (p *RequestURL) Query() string {
	params := make([]string, 0, len(p.params))
	for _, key := range p.keys {
		if param, ok := p.params[key]; ok {
			params = append(params, param)
		}
	}
	return strings.Join(params, "&")
}

func (p *RequestURL) Host() string {
	return p.host
}

func (p *RequestURL) Path() string {
	return p.path
}

func (p *RequestURL) String() string {
	return p.host + p.path + "?" + p.Query()
}

func AddCommonParam(u *RequestURL) {
	u.AddParam("deviceid", "786ed36a5b9a7b543248e7ff32bf40ec")
	u.AddParam("apiver", "4.3")
	u.AddParam("ver", "5.1.22")
	u.AddParam("miuiver", 6)
	u.AddParam("devicetype", 1)
	u.AddParam("ptf", 201)
	u.AddParam("ts", time.Now().Unix())
	u.AddParam("nonce", rand.Int())

	opaque := u.Opaque(MI_REQUEST_KEY, "&token="+MI_REQUEST_TOKEN)
	u.AddParam("opaque", opaque)
}

func DoPost(u *RequestURL, data interface{}) (err error) {
	bodyBuf := bytes.NewBufferString(u.Query())
	reqURL := u.Host() + u.Path()
	res, err := http.Post(reqURL, "application/x-www-form-urlencoded", bodyBuf)
	if err != nil {
		return
	}
	defer res.Body.Close()

	//	b, _ := ioutil.ReadAll(res.Body)
	//	fmt.Println(string(b))

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(data)
	if err != nil {
		fmt.Printf("Url: %s, data: %s\n", reqURL, u.Query())
	}
	return
}
