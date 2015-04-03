package iqiyi

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"io"
	"net/http"
	"strconv"
	"time"
)

type Transport struct {
	key       string
	n         int64
	mKey      string
	clientVer string
}

func NewTransport(key, mKey, clientVer string, n int64) *Transport {
	return &Transport{
		key:       key,
		mKey:      mKey,
		clientVer: clientVer,
		n:         n,
	}
}

func (c *Transport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	var (
		ts     = time.Now().Unix()
		h      = md5.New()
		t      = strconv.FormatInt(ts^c.n, 10)
		buf    = bytes.NewBufferString("")
		header = make(http.Header)
	)

	buf.WriteString(strconv.FormatInt(ts, 10))
	buf.WriteString(c.key)
	buf.WriteString(c.mKey)
	buf.WriteString(c.clientVer)
	_, err = io.Copy(h, buf)

	if err != nil {
		return
	}

	header.Set("t", t)
	header.Set("sign", hex.EncodeToString(h.Sum(nil)))
	req.Header = header

	resp, err = http.DefaultTransport.RoundTrip(req)

	return
}
