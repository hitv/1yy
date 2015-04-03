package fetch

type MediaURL struct {
	Status    int    `json:"status"`
	VideoName string `json:"videoname"`
	Normal    []struct {
		Source  int    `json:"source"`
		IsHTML  int    `json:"ishtml"`
		PlayURL string `json:"playurl"`
	} `json:"normal"`
	High []struct {
		Source  int    `json:"source"`
		IsHTML  int    `json:"ishtml"`
		PlayURL string `json:"playurl"`
	} `json:"high"`
	Super []struct {
		Source  int    `json:"source"`
		IsHTML  int    `json:"ishtml"`
		PlayURL string `json:"playurl"`
	} `json:"super"`
}

func FetchMediaURL(mediaId, ci, source int) (data *MediaURL, err error) {
	u := NewRequestURL(MI_HOST, "/tvservice/getmediaurl")
	u.AddParam("mediaid", mediaId)
	u.AddParam("ci", ci)
	u.AddParam("source", source)
	AddCommonParam(u)

	data = &MediaURL{}
	err = DoPost(u, data)

	return
}
