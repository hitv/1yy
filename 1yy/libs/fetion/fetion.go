package fetion

import (
	"errors"
	"fmt"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
	"wx-novel/app/models/cache"
)

var (
	NOT_LOGIN_ERROR = errors.New("尚未登录飞信")
)

type Fetion struct {
	UserName   string
	Password   string
	client     *http.Client
	NickName   string
	LoginState string
	IdUser     string
}

func NewFetion(username, password string, cache cache.Cache) *Fetion {
	return &Fetion{
		UserName: username,
		Password: password,
		client: &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				if strings.Contains(req.URL.Path, "login.action") {
					return errors.New("logout_page")
				}
				return nil
			},
			Jar: &Jar{cache, time.Hour * 24, "fetion"},
		},
	}
}
func isRedirectLogout(err error) bool {
	if strings.HasSuffix(err.Error(), "logout_page") {
		return true
	}
	return false
}
func checkLoginRedirect(res *http.Response, err error) (*simplejson.Json, error) {
	if err != nil {
		if isRedirectLogout(err) {
			return nil, nil
		}
		return nil, err
	}
	return simplejson.NewFromReader(res.Body)
}
func (f *Fetion) httpGet(u string) (*simplejson.Json, error) {
	return checkLoginRedirect(f.client.Get(u))
}

func (f *Fetion) httpPost(u string, kv map[string]string) (*simplejson.Json, error) {
	form := url.Values{}
	for k, v := range kv {
		form.Set(k, v)
	}
	return checkLoginRedirect(f.client.PostForm(u, form))
}

//获取登录验证码
func (f *Fetion) GetCaptchaCode() ([]byte, error) {
	res, err := f.client.Get(fmt.Sprintf("http://f.10086.cn/im5/systemimage/verifycode%d.png?tp=im5", time.Now().UnixNano()))
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

//登录
func (f *Fetion) Login(captchaCode string) error {
	form := map[string]string{
		"m":            f.UserName,
		"pass":         f.Password,
		"captchaCode":  captchaCode,
		"checkCodeKey": "null",
	}

	json, err := f.httpPost("http://f.10086.cn/im5/login/loginHtml5.action", form)
	if err != nil {
		return err
	}
	if json == nil {
		return NOT_LOGIN_ERROR
	}
	loginState, _ := json.Get("loginstate").String()
	if loginState == "" {
		tip, _ := json.Get("tip").String()
		return fmt.Errorf(tip)
	}
	f.NickName, _ = json.Get("nickname").String()
	f.LoginState, _ = json.Get("loginstate").String()
	f.IdUser, _ = json.Get("idUser").String()
	return nil
}

// 获取即时消息列表
func (f *Fetion) GetAllList() ([]byte, error) {
	res, err := f.client.Get("http://f.10086.cn/im5/box/alllist.action")
	if err != nil {
		if isRedirectLogout(err) {
			return nil, NOT_LOGIN_ERROR
		}
		return nil, err
	}
	return ioutil.ReadAll(res.Body)
}

//获取联系人分组数据
func (f *Fetion) GetGroupContacts() error {
	json, err := f.httpGet("http://f.10086.cn/im5/index/loadGroupContactsAjax.action")
	if err != nil {
		return err
	}
	if json == nil {
		return NOT_LOGIN_ERROR
	}
	fmt.Printf("json: %+v, error: %s\n", json, err)
	return nil
}

//获取某个分组的联系人列表
func (f *Fetion) GetContactList(listId int) error {
	json, err := f.httpGet(fmt.Sprintf("http://f.10086.cn/im5/index/contactlistView.action?idContactList=%d", listId))
	if err != nil {
		return err
	}
	if json == nil {
		return NOT_LOGIN_ERROR
	}
	fmt.Printf("json: %+v, error: %s\n", json, err)
	return nil
}

// 添加好友响应
func (f *Fetion) HandleContactRequest(reqId string) (string, error) {

	form := map[string]string{
		"idAddContactRequest": reqId,
		"requesterIdFetion":   "990789512",
		"requesterMobileNo":   "0",
		"response1":           "1",
		"requesterIdUser":     "726218956",
		"requestNickname":     "zlx",
	}
	json, err := f.httpPost("http://f.10086.cn/im5/box/handleContactRequest.action", form)
	if err != nil {
		return "", err
	}
	return json.Get("tip").String()
}

//按手机号搜索好友
func (f *Fetion) SearchFriendByPhone(phoneNum string) (string, error) {
	uid := 0
	//首先查询该手机号是否已加为好友
	u := fmt.Sprintf("http://f.10086.cn/im5/index/searchFriendsByQueryKey.action?queryKey=%s", phoneNum)
	json, err := f.httpGet(u)
	fmt.Printf("json: %+v, err: %+v\n", json, err)
	//若没加好友，响应内容为空，此处err不为nil
	if err != nil {
		//查找该手机号对应的id
		u := fmt.Sprintf("http://f.10086.cn/im5/user/searchFriendByPhone.action?number=%s", phoneNum)
		json, err := f.httpGet(u)
		if err != nil {
			return "", err
		}
		if json == nil {
			return "", NOT_LOGIN_ERROR
		}
		tip, _ := json.Get("tip").String()
		if tip != "" {
			return "", errors.New(tip)
		}
		uid, err = json.GetPath("userinfo", "idUser").Int()
	} else {
		uid, err = json.GetPath("contacts", "idContact").Int()
	}
	return fmt.Sprintf("%d", uid), err
}

//群发消息
func (f *Fetion) SendGroupMsg(uids, msg string) error {
	//{"sendCode":"200","info":"发送成功"}
	form := map[string]string{
		"touserid": uids,
		"msg":      msg,
	}
	json, err := f.httpPost("http://f.10086.cn/im5/chat/sendNewGroupShortMsg.action", form)
	if err != nil {
		return err
	}
	if json == nil {
		return NOT_LOGIN_ERROR
	}
	fmt.Printf("json: %+v, error: %s\n", json, err)
	return nil
}

//发消息给好友
func (f *Fetion) SendMsg(uid, msg string) error {
	form := map[string]string{
		"touserid": uid,
		"msg":      msg,
	}
	json, err := f.httpPost("http://f.10086.cn/im5/chat/sendNewMsg.action", form)
	if err != nil {
		return err
	}
	if json == nil {
		return NOT_LOGIN_ERROR
	}
	status, err := json.Get("sendCode").String()
	if err != nil {
		return err
	}
	info, err := json.Get("info").String()
	if err != nil {
		return err
	}
	if status != "200" {
		return fmt.Errorf("Send message error: %s", info)
	}
	return nil
}

//发送短信
func (f *Fetion) SendShortMsg(uids, msg string) error {
	form := map[string]string{
		"touserid": uids,
		"msg":      msg,
	}
	json, err := f.httpPost("http://f.10086.cn/im5/chat/sendNewShortMsg.action", form)
	if err != nil {
		return err
	}
	if json == nil {
		return NOT_LOGIN_ERROR
	}
	status, err := json.Get("sendCode").String()
	if err != nil {
		return err
	}
	info, err := json.Get("info").String()
	if err != nil {
		return err
	}
	if status != "200" {
		return fmt.Errorf("Send message error: %s", info)
	}
	return nil
}

//获取个人信息
func (f *Fetion) GetSelfInfo() (int, error) {
	json, err := f.httpGet("http://f.10086.cn/im5/user/selfInfo.action")
	if err != nil {
		return 0, err
	}
	if json == nil {
		return 0, NOT_LOGIN_ERROR
	}
	return json.GetPath("userinfo", "idUser").Int()
}

//退出登录
func (f *Fetion) Logout() error {
	_, err := f.httpGet("http://f.10086.cn/im5/index/logoutsubmit.action")
	if err != nil {
		return err
	}
	return nil
}
