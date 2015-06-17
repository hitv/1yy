package mivideo

import "github.com/bitly/go-simplejson"

type Group struct {
	Title     string
	SubTitle  string
	TargetURL string
	Special   *GroupItem
	Items     []*GroupItem
	*Time
}

func parseGroup(json *simplejson.Json, titleIndex, specialIndex, itemsIndex int) (*Group, error) {
	var (
		title, subTitle string
		special         *GroupItem
		groupItems      []*GroupItem
	)

	blocks, err := json.Get("blocks").Array()
	if err != nil {
		return nil, ErrFormatInvalid
	}

	if len(blocks) < 3 {
		return nil, ErrFormatInvalid
	}

	if titleIndex > -1 {
		block := simplejson.NewFromInterface(blocks[titleIndex])
		title = block.Get("title").MustString()
		subTitle = block.Get("sub_title").MustString()
	}
	if specialIndex > -1 {
		block := simplejson.NewFromInterface(blocks[specialIndex])
		special, err = parseGroupItem(block)
		if err != nil {
			return nil, err
		}
	}
	if itemsIndex > -1 {
		var items []interface{}
		block := simplejson.NewFromInterface(blocks[itemsIndex])
		items, err = block.Get("items").Array()
		if err != nil {
			return nil, ErrFormatInvalid
		}
		for _, item := range items {
			groupItem, err := parseGroupItem(simplejson.NewFromInterface(item))
			if err != nil {
				return nil, err
			}
			if groupItem.Entity != "play" && groupItem.Entity != "pvideo" {
				continue
			}
			groupItems = append(groupItems, groupItem)
		}
	}

	if len(groupItems) == 0 {
		return nil, ErrFormatInvalid
	}

	return &Group{
		Title:    title,
		SubTitle: subTitle,
		Special:  special,
		Items:    groupItems,
		Time:     parseTime(json.Get("times")),
	}, nil
}

func parseGroups(json *simplejson.Json, titleIndex, specialIndex, itemsIndex int) (groups []*Group, err error) {
	blocks, err := json.Get("blocks").Array()
	if err != nil {
		err = ErrFormatInvalid
		return
	}

	for _, block := range blocks {
		json = simplejson.NewFromInterface(block)
		group, err := parseGroup(json, titleIndex, specialIndex, itemsIndex)
		if err != nil {
			continue
		}
		groups = append(groups, group)
	}

	return
}
