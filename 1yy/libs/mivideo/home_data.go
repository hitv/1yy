package mivideo

import (
	"strings"

	"github.com/bitly/go-simplejson"
)

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
		return nil, ErrFormatInvalid
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
			best, err = parseGroups(json, 0, 1, 2)
			if err != nil {
				continue
			}
		case strings.Contains(id, "rank.r"):
			rank, err = parseGroups(json, 0, 2, 1)
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
