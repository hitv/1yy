package mivideo

import "github.com/bitly/go-simplejson"

type GroupItem struct {
	Title     string
	SubTitle  string
	Poster    string
	Entity    string
	TargetURL string
	Hint      string
	*Time
}

func parseGroupItem(json *simplejson.Json) (item *GroupItem, err error) {
	title := json.Get("title").MustString()
	subTitle := json.Get("sub_title").MustString()
	hint := json.GetPath("hint", "left").MustString()
	entity := json.GetPath("target", "entity").MustString()
	targetURL := json.GetPath("target", "url").MustString()
	poster := json.GetPath("images", "poster", "url").MustString()
	times := parseTime(json.Get("times"))

	item = &GroupItem{
		Title:     title,
		SubTitle:  subTitle,
		Poster:    poster,
		Hint:      hint,
		Entity:    entity,
		TargetURL: targetURL,
		Time:      times,
	}
	return
}
