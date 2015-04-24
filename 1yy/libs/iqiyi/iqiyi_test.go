package iqiyi

import (
	"testing"
)

func TestGetM3u8(t *testing.T) {
	str, err := GetM3u8("http://www.iqiyi.com/v_19rro1vd70.html")
	if err != nil {
		t.Errorf("GetM3u8 error: %s", err)
	}
	t.Log(str)
}
