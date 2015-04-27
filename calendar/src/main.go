package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var (
	file  *os.File
	plans = make(map[string]int)
)

func ReadPlans() error {
	buf := bytes.NewBuffer([]byte{})
	_, err := io.Copy(buf, file)
	if err != nil {
		return err
	}
	lines := strings.Split(buf.String(), "\n")

	plans = make(map[string]int, len(lines))
	for _, plan := range lines {
		parts := strings.Split(plan, ":")
		if len(parts) > 1 {
			plans[parts[0]], _ = strconv.Atoi(parts[1])
		}
	}
	return nil
}

func SavePlans() error {
	if len(plans) == 0 {
		return file.Truncate(0)
	}

	lines := make([]string, 0)
	for day, value := range plans {
		lines = append(lines, fmt.Sprintf("%s:%d", day, value))
	}

	err := file.Truncate(0)
	if err != nil {
		return err
	}

	buf := bytes.NewBufferString(strings.Join(lines, "\n"))
	_, err = io.Copy(file, buf)
	return err
}

func GetPlan(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Header().Set("Content-type", "json")
	encoder := json.NewEncoder(w)
	encoder.Encode(plans)
}

func SetPlan(w http.ResponseWriter, r *http.Request) {
	var (
		day   = r.FormValue("day")
		value = r.FormValue("value")
		res   = map[string]interface{}{
			"success": true,
			"message": "保存成功",
		}
	)

	day = strings.TrimSpace(day)
	value = strings.TrimSpace(value)

	defer func() {
		encoder := json.NewEncoder(w)
		encoder.Encode(res)
	}()

	val, err := strconv.Atoi(value)
	if err != nil || day == "" {
		w.WriteHeader(400)
		res["success"] = false
		res["message"] = "参数错误"
		return
	}

	if val > -1 {
		plans[day] = val
	} else {
		delete(plans, day)
	}

	err = SavePlans()
	if err != nil {
		fmt.Printf("SavePlans error: %s", err)
		w.WriteHeader(500)
		res["success"] = false
		res["message"] = "保存出错"
	}

}

func main() {
	var err error
	file, err = os.OpenFile("plans.txt", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = ReadPlans()
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/plan", GetPlan)
	mux.HandleFunc("/plan/set", SetPlan)
	mux.Handle("/", http.FileServer(http.Dir("./public")))

	fmt.Println("start listen on :8001 port")
	http.ListenAndServe(":8001", mux)
}
