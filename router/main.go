package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

var (
	lastTime   time.Time
	routerIp   string
	routerPort string
	routerUser string
	routerPwd  string

	tokenReg = regexp.MustCompile(`(sysauth=[0-9a-f]{32});.*;stok=([0-9a-f]{32})`)
)

func routerHost() string {
	if routerIp == "" || routerPort == "" {
		return ""
	}
	return "http://" + routerIp + ":" + routerPort
}

func login(host, username, password string) (cookie, token string, err error) {
	res, err := http.PostForm(host+"/cgi-bin/luci", url.Values{
		"username": {username},
		"password": {password},
	})
	if err != nil {
		return
	}
	cookies := res.Header.Get("Set-Cookie")
	matches := tokenReg.FindAllStringSubmatch(cookies, -1)
	if len(matches) > 0 {
		cookie = matches[0][1]
		token = matches[0][2]
	}
	return
}

func reconnectRouterInterface(network, cookie, token string) error {
	host := routerHost()
	reqUrl := host + "/cgi-bin/luci/;stok=" + token + "/admin/network/iface_reconnect/" + network
	req, _ := http.NewRequest("GET", reqUrl, nil)
	req.Header.Set("Cookie", cookie)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode == 200 && strings.Contains(res.Status, "Reconnected") {
		return nil
	}
	return fmt.Errorf("reconnect %s failed", network)
}

func setRouterIpHandler(res http.ResponseWriter, req *http.Request) {
	routerIp = req.Header.Get("X-Forwarded-For")
	lastTime = time.Now()
	res.Header().Set("Content-Type", "text/plain")
	res.Write([]byte("time: " + lastTime.Format("2006-01-02 15:04:05") + "\nip: " + routerIp))
}

func getRouterIpHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/plain")
	res.Write([]byte("time: " + lastTime.Format("2006-01-02 15:04:05") + "\nip: " + routerIp))
}

func reconnectHandler(res http.ResponseWriter, req *http.Request) {
	clientIp := req.Header.Get("X-Forwarded-For")
	if clientIp != routerIp {
		res.WriteHeader(http.StatusUnauthorized)
		res.Write([]byte("您无权限执行此操作"))
		return
	}
	host := routerHost()
	cookie, token, err := login(host, routerUser, routerPwd)
	if err != nil {
		res.Write([]byte("登录出错"))
		log.Printf("login error: %s\n", err)
		return
	}
	if token == "" || cookie == "" {
		res.Write([]byte("登录失败"))
		log.Println("login error: token is empty")
		return
	}
	err = reconnectRouterInterface("wan", cookie, token)
	if err != nil {
		res.Write([]byte("重连失败"))
		log.Printf("reconnect %s error: %s", "wan", err)
		return
	}
	res.Write([]byte("重连成功"))
	routerIp = ""
}

func main() {
	var (
		listen string
	)
	flag.StringVar(&listen, "listen", ":8800", "The port to listen on")
	flag.StringVar(&routerIp, "routerIp", "", "The ip of router")
	flag.StringVar(&routerPort, "routerPort", "80", "The port router web server listen on")
	flag.StringVar(&routerUser, "routerUser", "root", "The username of router")
	flag.StringVar(&routerPwd, "routerPwd", "root", "The password of router")
	flag.Parse()

	http.HandleFunc("/reconnect", reconnectHandler)
	http.HandleFunc("/get_ip", getRouterIpHandler)
	http.HandleFunc("/set_ip", setRouterIpHandler)

	fmt.Printf("Start listen on %s", listen)
	err := http.ListenAndServe(listen, nil)
	if err != nil {
		log.Printf("Service listen error: %s", err)
	}
}
