package mivideo

import (
	"time"

	"github.com/bitly/go-simplejson"
)

type Time struct {
	CreateTime time.Time
	UpdateTime time.Time
}

func parseTime(json *simplejson.Json) (t *Time) {
	var (
		createTime = time.Now()
		updateTime = createTime
	)

	createTs, err := json.Get("created").Int64()
	if err == nil {
		createTime = time.Unix(createTs, 0)
	}

	updateTs, err := json.Get("updated").Int64()
	if err != nil {
		updateTime = time.Unix(updateTs, 0)
	}

	return &Time{createTime, updateTime}
}
