package types

import (
	"strings"
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

type HomeData struct {
	Hot  *VideoChannel
	Best []*Group
	Rank []*Group
}

func ParseHomeData(json *simplejson.Json) (data *HomeData, err error) {
	var (
		hot  *VideoChannel
		best []*Group
		rank []*Group
	)

	blocks, err := json.Get("blocks").Array()
	if err != nil {
		return
	}

	for _, block := range blocks {
		json = simplejson.NewFromInterface(block)
		id := json.Get("id").MustString()
		switch {
		case strings.Contains(id, "hot.r"):
			hot, err = parseVideoChannel(json)
			if err != nil {
				return
			}
		case strings.Contains(id, "attention.recomm.r"):
			best, err = parseGroups(json)
			if err != nil {
				continue
			}
		case strings.Contains(id, "rank.r"):
			rank, err = parseGroups(json)
			if err != nil {
				return
			}
		}
	}
	data = &HomeData{
		Hot:  hot,
		Best: best,
		Rank: rank,
	}

	return
}
