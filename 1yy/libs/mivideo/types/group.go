package types

import (
	"log"

	"github.com/bitly/go-simplejson"
	"hi.tv/1yy/libs/mivideo/errors"
)

type Group struct {
	Title     string
	SubTitle  string
	TargetURL string
	Special   *GroupItem
	Items     []*GroupItem
	*Time
}

func parseGroup(json *simplejson.Json) (group *Group, err error) {
	blocks, err := json.Get("blocks").Array()
	if err != nil {
		log.Printf("111")
		return
	}

	if len(blocks) < 3 {
		err = errors.ErrFormatInvalid
		return
	}

	block := simplejson.NewFromInterface(blocks[0])
	title := block.Get("title").MustString()
	subTitle := block.Get("sub_title").MustString()

	block = simplejson.NewFromInterface(blocks[1])
	special, err := parseGroupItem(block)
	if err != nil {
		return
	}

	block = simplejson.NewFromInterface(blocks[2])
	items, err := block.Get("items").Array()
	if err != nil {
		log.Printf("222")
		return
	}

	groupItems := make([]*GroupItem, 0, len(items))
	for _, item := range items {
		groupItem, err := parseGroupItem(simplejson.NewFromInterface(item))
		if err != nil || groupItem.Entity != "play" {
			continue
		}
		groupItems = append(groupItems, groupItem)
	}

	times := parseTime(json.Get("times"))

	group = &Group{
		Title:    title,
		SubTitle: subTitle,
		Special:  special,
		Items:    groupItems,
		Time:     times,
	}

	return
}

func parseGroups(json *simplejson.Json) (groups []*Group, err error) {
	blocks, err := json.Get("blocks").Array()
	if err != nil {
		return
	}
	log.Printf("len(blocks): %d\n", len(blocks))
	for _, block := range blocks {
		json = simplejson.NewFromInterface(block)
		group, err := parseGroup(json)
		if err != nil {
			continue
		}
		groups = append(groups, group)
	}
	return
}
