package fetch

type BannerMedia struct {
}

func (c *BannerMedia) Fetch(channelId int) (err error) {
	u := NewRequestURL(MI_HOST, "/tvservice/getbannermedia")
	u.AddParam("channelid", channelId)
	u.AddParam("userbehavdata", "a")
	AddCommonParam(u)
	return DoPost(u, c)
}
