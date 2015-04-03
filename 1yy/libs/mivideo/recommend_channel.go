package fetch

type RecommendChannel struct {
}

func (c *RecommendChannel) Fetch(channelId int) (err error) {
	u := NewRequestURL(MI_HOST, "/tvservice/getrecommendchannel")
	u.AddParam("channelid", channelId)
	return DoPost(u, c)
}
