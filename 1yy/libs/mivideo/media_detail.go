package fetch

import (
	"github.com/hitv/1yy/app/models"
	"github.com/hitv/1yy/app/models/enums"
	"github.com/hitv/1yy/app/utils"
)

type MediaInfo struct {
	Actors        string `json:"actors"`
	Categorys     string `json:"allcategorys"`
	Area          string `json:"area"`
	Category      string `json:"category"`
	Desc          string `json:"desc"`
	ShortDesc     string `json:"shortdesc"`
	Director      string `json:"director"`
	Flag          int32  `json:"flag"`
	IsMultSet     int32  `json:"ismultset"`
	IssueDate     int32  `json:"issuedate"`
	Language      string `json:"language"`
	LastIssueDate string `json:"lastissuedate"`
	Md5           string `json:"md5"`
	DownloadSrc   []int  `json:"media_available_download_source"`
	MediaId       uint32 `json:"mediaid"`
	MediaName     string `json:"medianame"`
	MidType       int32  `json:"midtype"`
	NameIndex     struct {
		Capital      string `json:"capital"`
		FuzzyCapital string `json:"fuzzycapital"`
		Spelling     string `json:"spelling"`
		Symbol       string `json:"symbol"`
		ValidLength  uint32 `json:"validlength"`
	} `json:"nameindex"`
	PlayCount      uint32  `json:"playcount"`
	PlayLength     float32 `json:"playlength"`
	PosterURL      string  `json:"posterurl"`
	Resolution     int32   `json:"resolution"`
	Score          float64 `json:"score"`
	ScoreCount     uint32  `json:"scorecount"`
	SetCount       float32 `json:"setcount"`
	SetNow         int32   `json:"setnow"`
	SmallPosterMd5 string  `json:"smallpostermd5"`
	SmallPosterURL string  `json:"smallposterurl"`
	WebpPosterMd5  string  `json:"webpmd5"`
	WebpPosterURL  string  `json:"webpposterurl"`
}

func (c *MediaInfo) Save() (id uint32, err error) {
	v := &models.VideoModel{
		Title:             c.MediaName,
		Actors:            c.Actors,
		Director:          c.Director,
		Categorys:         c.Categorys,
		Category:          c.Category,
		Area:              c.Area,
		Description:       c.Desc,
		ShortDescription:  c.ShortDesc,
		Flag:              c.Flag,
		IsMultSet:         c.IsMultSet,
		IssueDate:         c.IssueDate,
		Language:          c.Language,
		LastIssueDate:     c.LastIssueDate,
		Md5:               c.Md5,
		DownloadSrc:       utils.JoinIntsToString(c.DownloadSrc, ","),
		MidType:           c.MidType,
		IndexCapital:      c.NameIndex.Capital,
		IndexFuzzyCapital: c.NameIndex.FuzzyCapital,
		IndexSpelling:     c.NameIndex.Spelling,
		IndexSymbol:       c.NameIndex.Symbol,
		IndexValidLength:  c.NameIndex.ValidLength,
		PlayCount:         c.PlayCount,
		PlayLength:        c.PlayLength,
		PosterURL:         c.PosterURL,
		Resolution:        c.Resolution,
		Score:             c.Score,
		ScoreCount:        c.ScoreCount,
		SetCount:          c.SetCount,
		SetNow:            c.SetNow,
		SmallPosterMd5:    c.SmallPosterMd5,
		SmallPosterURL:    c.SmallPosterURL,
		WebpPosterMd5:     c.WebpPosterMd5,
		WebpPosterURL:     c.WebpPosterURL,
		Status:            enums.StatusEnabled.String(),
		MiId:              c.MediaId,
	}
	video, err := models.VideoDao.GetByMid(c.MediaId)
	if err != nil {
		if err != models.ErrNotFound {
			return
		}
		id, err = models.VideoDao.Insert(video)
		return
	}

	id = video.Id
	err = models.VideoDao.Update(v)

	return
}

type MediaDetail struct {
	Status int `json:"status"`
	Data   struct {
		MediaInfo MediaInfo `json:"mediainfo"`
	} `json:"data"`
}

type MediaFilter struct {
	Status int         `json:"status"`
	Count  int         `json:"count"`
	Data   []MediaInfo `json:"data"`
}

func FetchMediaDetail(mediaId uint32) (data *MediaDetail, err error) {
	u := NewRequestURL(MI_HOST, "/tvservice/getmediadetail2")
	u.AddParam("mediaid", mediaId)
	u.AddParam("fee", 1)
	u.AddParam("userbehavdata", "{}")
	u.AddParam("pageno", 1)
	u.AddParam("orderby", -1)
	u.AddParam("pagesize", 1000)
	AddCommonParam(u)

	data = &MediaDetail{}
	err = DoPost(u, data)

	return
}

func FetchMediaFilter(channelId uint32, pageSize, page int) (data *MediaFilter, err error) {
	u := NewRequestURL(MI_HOST, "/tvservice/filtermediainfo")
	u.AddParam("channelids", channelId)
	u.AddParam("pageno", page)
	u.AddParam("pagesize", pageSize)
	u.AddParam("orderby", 1)
	u.AddParam("listtype", -1)
	u.AddParam("postertype", -1)
	u.AddParam("searchtype", -1)
	u.AddParam("fee", 1)
	AddCommonParam(u)

	data = &MediaFilter{}
	err = DoPost(u, data)

	return
}
