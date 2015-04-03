package fetion

import (
	"testing"
	"time"
	"wx-novel/app/models/cache"
)

var fetion *Fetion

func init() {
	config := &cache.RedisConfig{
		Host: "127.0.0.1:6379",
	}
	redisPool := cache.NewRedisPool(config)
	redisCache := cache.NewRedisCache(redisPool, 240*time.Hour)
	fetion = NewFetion("804268310", "zlx#747836", redisCache)
}
func TestGetCaptchaCode(t *testing.T) {
	body, err := fetion.GetCaptchaCode()
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("CaptchaCode data is: %+v", body)
	if len(body) == 0 {
		t.Errorf("CaptchaCode data length is 0")
	}
}

func TestFetionLogin(t *testing.T) {

}
func TestAllList(t *testing.T) {

}
func TestSearchFriendByPhone(t *testing.T) {
	idUser, err := fetion.SearchFriendByPhone("15021695001")
	t.Logf("idUser: %s, error: %s\n", idUser, err)
}
func TestGetGroupContacts(t *testing.T) {

}
func TestGetContactList(t *testing.T) {

}
func TestSendMsg(t *testing.T) {

}
func TestGetSelfInfo(t *testing.T) {

}
