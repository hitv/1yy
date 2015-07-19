package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"
)

const (
	SHAKE_CODE_URL_FMT = "http://cms.zhcs0439.com:8686/client/app/shake/shake.do?cityCode=101060901&IMSI=3&code=&phoneNumber=%s&IMEI=863583035931452"
	SHAKE_URL_FMT      = "http://cms.zhcs0439.com:8686/client/app/shake/shake.do?phoneNumber=%s&IMSI=3&IMEI=863583035931452&code=%s&cityCode=101060901"
)

var (
	codeReg   = regexp.MustCompile(`"code":"(.+?)"`)
	resultReg = regexp.MustCompile(`"result":"((\d|\.)+)(慧币)"`)
)

type ShakeResult struct {
	Sn    int
	Phone string
	Award [2]string
	Err   error
}

func (r *ShakeResult) AwardByteNum() int {
	n, _ := strconv.ParseFloat(r.Award[0], 32)
	switch r.Award[1] {
	case "GB":
		n *= (1 << 30)
	case "MB":
		n *= (1 << 20)
	case "KB":
		n *= (1 << 10)
	}
	return int(n)
}

func shakeGet(url string) (body []byte, err error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("user-agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1;SV1)")
	req.Header.Set("connection", "Keep-Alive")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	buf := bytes.NewBuffer([]byte{})
	_, err = io.Copy(buf, res.Body)

	body = buf.Bytes()
	return
}

func getShakeCode(phone string) (code string, err error) {
	url := fmt.Sprintf(SHAKE_CODE_URL_FMT, phone)

	body, err := shakeGet(url)
	if err != nil {
		return
	}
	matches := codeReg.FindSubmatch(body)
	if 2 != len(matches) {
		err = fmt.Errorf("Can't find code")
		return
	}

	md5Hash := md5.New()
	_, err = md5Hash.Write(matches[1])
	if err != nil {
		return
	}
	code = hex.EncodeToString(md5Hash.Sum(nil))
	return
}

func shake(ch chan *ShakeResult, phone string, sn int) {
	code, err := getShakeCode(phone)
	if err != nil {
		ch <- &ShakeResult{
			Sn:    sn,
			Phone: phone,
			Err:   err,
		}
		return
	}
	url := fmt.Sprintf(SHAKE_URL_FMT, phone, code)
	body, err := shakeGet(url)
	if err != nil {
		ch <- &ShakeResult{
			Sn:    sn,
			Phone: phone,
			Err:   err,
		}
		return
	}
	matches := resultReg.FindSubmatch(body)
	if 4 != len(matches) {
		ch <- &ShakeResult{
			Sn:    sn,
			Phone: phone,
			Err:   fmt.Errorf("Response body(%s) format error", body),
		}
		return
	}

	ch <- &ShakeResult{
		Sn:    sn,
		Phone: phone,
		Award: [2]string{string(matches[1]), string(matches[3])},
	}
}

func byteNumToUnit(byteNum int) string {
	var (
		unitNums    = []int{1 << 30, 1 << 20, 1 << 10, 0}
		unitSymbols = []string{"GB", "MB", "KB", "B"}
		unitNum     = float32(byteNum)
		unitSymbol  string
	)
	for i, v := range unitNums {
		if byteNum > v {
			unitNum = float32(byteNum) / float32(v)
			unitSymbol = unitSymbols[i]
			break
		}
	}
	return fmt.Sprintf("%.2f%s", unitNum, unitSymbol)
}

func parseArgs(dftNum int) map[string]int {
	var (
		nums      []int
		phones    []string
		phoneNums = make(map[string]int)
	)
	if len(os.Args) > 1 {
		phones = strings.Split(os.Args[1], ",")
	}
	if len(os.Args) > 2 {
		argNums := strings.Split(os.Args[2], ",")
		for _, n := range argNums {
			num, err := strconv.Atoi(n)
			if err == nil {
				nums = append(nums, num)
			}
		}
	}
	for i, v := range phones {
		if i < len(nums) {
			phoneNums[v] = nums[i]
		} else {
			phoneNums[v] = dftNum
		}
	}
	return phoneNums
}

func notifySignal(fn func(os.Signal)) {
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGINT, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGTERM)

	go func() {
		fn(<-sigCh)
	}()
}

func main() {
	var (
		phoneNums = parseArgs(100)
		awardMap  = make(map[string]float64)
		resultCh  = make(chan *ShakeResult, 100)

		report = func() {
			fmt.Println("== End shake ==")
			for phone, num := range awardMap {
				fmt.Printf("-- Phone %s shake %d times, total award %.01f\n", phone, phoneNums[phone], num)
			}
		}
	)

	if len(phoneNums) == 0 {
		fmt.Println("Usage: shake phone1[,phone2[,...]] num ")
		return
	}

	cpuNum := runtime.NumCPU()
	fmt.Printf("== CPU Num is %d ==\n", cpuNum)
	runtime.GOMAXPROCS(cpuNum)

	notifySignal(func(sig os.Signal) {
		report()
		os.Exit(0)
	})
	for phone, num := range phoneNums {
		fmt.Printf("== Start shake %s %d times ==\n", phone, num)
		go func(){
			for i := 0; i < num; i++ {
				shake(resultCh, phone, i)
				sleepDuration := time.Second * 3
				fmt.Printf("-- Sleep %s... --\n", sleepDuration)
				time.Sleep(sleepDuration)

			}
		}()
	}

	for _, num := range phoneNums {
		for i := 0; i < num; i++ {
			result := <-resultCh
			fmt.Printf("-- Shake %s[%d/%d], ", result.Phone, result.Sn, phoneNums[result.Phone])
			if result.Err == nil {
				fmt.Printf("award: %s%s", result.Award[0], result.Award[1])
				award, _ := strconv.ParseFloat(result.Award[0], 64)
				awardMap[result.Phone] += award
			} else {
				fmt.Printf("error: %s", result.Err)

			}

			fmt.Println(" --")

		}
	}

	report()
}
