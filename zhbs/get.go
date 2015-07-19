package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"
)

type Code struct {
	ReCode int    `json:"recode"`
	Msg    string `json:"msg"`
	Result string `json:"result"`
}

type Result struct {
	ReCode int    `json:"recode"`
	Msg    string `json:"msg"`
}

func shakeGet(url string) (r io.Reader, err error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("user-agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1;SV1)")
	req.Header.Set("connection", "Keep-Alive")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, res.Body)

	return buf, err
}

func step1(phone, id, contentId string) (code *Code, err error){
	u1 := fmt.Sprintf("http://cms.zhcs0439.com/client/robConfigure/robConfigureCode.do?phoneNumber=%s&id=%s&content_id=%s", phone, id, contentId)
	r, err := shakeGet(u1)
	if err != nil {
		err = fmt.Errorf("get configure code error: %s\n", err)
		return
	}
	//var r io.Reader = bytes.NewReader([]byte(`{"recode":1,"msg":"验证成功，可返回验证码！","result":"9db17da1314fbdfb5f00ee30b94e89b1"}`))
	code = &Code{}
	decoder := json.NewDecoder(r)
	err = decoder.Decode(code)
	if err != nil {
		err = fmt.Errorf("decode configure code error: %s\n", err)
		return
	}
	if code.ReCode == 0 {
		err = fmt.Errorf("configure code return: %d, not 1, msg: %s\n", code.ReCode, code.Msg)
		return
	}
	return
}

func step2(phone, code, id, contentId string) (result *Result, err error){
	u2 := fmt.Sprintf("http://cms.zhcs0439.com/client/robConfigure/robConfigure.do?phoneNumber=%s&code=%s&id=%s&content_id=%s", phone, code, id, contentId)
	r, err := shakeGet(u2)
	if err != nil {
		err = fmt.Errorf("get robconfig error: %s\n", err)
		return
	}
	result = &Result{}
	decoder := json.NewDecoder(r)
	err = decoder.Decode(result)
	if err != nil {
		err = fmt.Errorf("decode robconfig error: %s\n", err)
		return
	}

	return
}

func main() {
	// get 158xxxxxxxx 38 3888
	runtime.GOMAXPROCS(runtime.NumCPU() * 2)
	if len(os.Args) < 4 {
		log.Printf("Usage: %s phone id content_id\n", os.Args[0])
		return
	}
	lock := make(chan bool)
	for m := 0; m < 500; m++ {
		code, err := step1(os.Args[1], os.Args[2], os.Args[3])
		if err != nil {
			log.Printf("[step1:fail:%d] get code error: %s\n", m, err)
			time.Sleep(time.Millisecond * 50)
			continue
		}
		log.Printf("[step1:info:%d] get code: %s\n", m, code.Result)

		for i := 0; i < 10; i++ {
			go func(n int) {
				result, err := step2(os.Args[1], code.Result, os.Args[2], os.Args[3])
				if err != nil {
					log.Printf("[step2:fail:%d] get robconfig error: %s\n", n, err)
					return
				}

				if result.ReCode == 1 {
					log.Printf("[step2:success:%d] get robconfig return: %d, msg: %s\n", i, result.ReCode, result.Msg)
					lock <- true
					return
				}
				log.Printf("[step2:fail:%d] get robconfig return: %d, not 1, msg: %s\n", i, result.ReCode, result.Msg)
			}(i)
		}
		break
	}
	<- lock
}
