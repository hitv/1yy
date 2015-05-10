package types

import "github.com/bitly/go-simplejson"

type VideoChannel struct {
	Title    string
	Slides   []*GroupItem
	Channels []*Channel
	*Time
}

func parseSlides(json *simplejson.Json) (slides []*GroupItem, err error) {
	items, err := json.Get("items").Array()
	if err == nil {
		for _, item := range items {
			groupItem, err := parseGroupItem(simplejson.NewFromInterface(item))
			if err != nil || groupItem.Entity != "pvideo" {
				continue
			}
			slides = append(slides, groupItem)
		}
	} /* else {
		log.Println(err.Error())
	}*/
	return
}

func parseVideoChannel(json *simplejson.Json) (hot *VideoChannel, err error) {
	title := json.Get("title").MustString()
	blocks, err := json.Get("blocks").Array()
	if err != nil {
		//log.Printf("parseHot, 1 error: %s\n", err)
		return
	}

	times := parseTime(json.Get("times"))

	var (
		slides   []*GroupItem
		channels []*Channel
	)

	// slides
	if len(blocks) > 0 {
		json := simplejson.NewFromInterface(blocks[0])
		slides, _ = parseSlides(json)
	}

	for _, block := range blocks {
		json := simplejson.NewFromInterface(block)
		channel, err := parseChannel(json)
		if err != nil {
			continue
		}
		channels = append(channels, channel)
	}

	hot = &VideoChannel{
		Title:    title,
		Slides:   slides,
		Channels: channels,
		Time:     times,
	}

	return
}
