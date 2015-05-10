package models

var VideoDao = &videoDao{}

type VideoModel struct {
	Id                int64
	Title             string
	Actors            string
	Director          string
	Categorys         string
	Category          string
	Area              string
	Description       string
	ShortDescription  string
	Flag              int64
	IsMultSet         int64
	IssueDate         int64
	Language          string
	LastIssueDate     string
	Md5               string
	DownloadSrc       string
	MidType           int64
	IndexCapital      string
	IndexFuzzyCapital string
	IndexSpelling     string
	IndexSymbol       string
	IndexValidLength  int64
	PlayCount         int64
	PlayLength        float64
	PosterURL         string
	Resolution        int64
	Score             float64
	ScoreCount        int64
	SetCount          float64
	SetNow            int64
	SmallPosterMd5    string
	SmallPosterURL    string
	WebpPosterMd5     string
	WebpPosterURL     string
	Status            string
	MiId              int64
}

func (m *VideoModel) TableName() string {
	return "video"
}

type videoDao struct{}

func (c *videoDao) GetByMid(miId int64) (video *VideoModel, err error) {
	video = &VideoModel{}
	has, err := db.Where("mi_id=?", miId).Get(video)
	if err == nil && !has {
		err = ErrNotExist
	}
	return
}

func (c *videoDao) Insert(video *VideoModel) (int64, error) {
	return db.Insert(video)
}

func (c *videoDao) Update(video *VideoModel) (int64, error) {
	return db.Update(video)
}
