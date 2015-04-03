package models

var VideoDao = &videoDao{}

type VideoModel struct {
	Id                uint32
	Title             string
	Actors            string
	Director          string
	Categorys         string
	Category          string
	Area              string
	Description       string
	ShortDescription  string
	Flag              int32
	IsMultSet         int32
	IssueDate         int32
	Language          string
	LastIssueDate     string
	Md5               string
	DownloadSrc       string
	MidType           int32
	IndexCapital      string
	IndexFuzzyCapital string
	IndexSpelling     string
	IndexSymbol       string
	IndexValidLength  uint32
	PlayCount         uint32
	PlayLength        float32
	PosterURL         string
	Resolution        int32
	Score             float64
	ScoreCount        uint32
	SetCount          float32
	SetNow            int32
	SmallPosterMd5    string
	SmallPosterURL    string
	WebpPosterMd5     string
	WebpPosterURL     string
	Status            string
	MiId              uint32
}

type videoDao struct{}

func (c *videoDao) GetByMid(miId uint32) (data *VideoModel, err error) {
	data = &VideoModel{}

	row := DefaultDB.QueryRow("SELECT `id`,`title`,`actors`,`director`,`categorys`,`category`,`area`,`desc`,`short_desc`,`flag`,`is_multset`,`issue_date`,`language`,`last_issue_date`,`md5`,`download_src`,`mid_type`,`index_capital`,`index_fuzzycapital`,`index_spelling`,`index_symbol`,`index_validlength`,`play_count`,`play_length`,`poster_url`,`resolution`,`score`,`score_count`,`set_count`,`set_now`,`small_poster_md5`,`small_poster_url`,`webp_poster_md5`,`webp_poster_url`,`status`,`mi_id` FROM video WHERE mi_id=?", miId)
	err = row.Scan(&data.Id, &data.Title, &data.Actors, &data.Director, &data.Categorys, &data.Category, &data.Area, &data.Description, &data.ShortDescription, &data.Flag, &data.IsMultSet, &data.IssueDate, &data.Language, &data.LastIssueDate, &data.Md5, &data.DownloadSrc, &data.MidType, &data.IndexCapital, &data.IndexFuzzyCapital, &data.IndexSpelling, &data.IndexSymbol, &data.IndexValidLength, &data.PlayCount, &data.PlayLength, &data.PosterURL, &data.Resolution, &data.Score, &data.ScoreCount, &data.SetCount, &data.SetNow, &data.SmallPosterMd5, &data.SmallPosterURL, &data.WebpPosterMd5, &data.WebpPosterURL, &data.Status, &data.MiId)
	return
}

func (c *videoDao) Insert(data *VideoModel) (id uint32, err error) {
	rs, err := DefaultDB.Exec("INSERT INTO video(`title`,`actors`,`director`,`categorys`,`category`,`area`,`desc`,`short_desc`,`flag`,`is_multset`,`issue_date`,`language`,`last_issue_date`,`md5`,`download_src`,`mid_type`,`index_capital`,`index_fuzzycapital`,`index_spelling`,`index_symbol`,`index_validlength`,`play_count`,`play_length`,`poster_url`,`resolution`,`score`,`score_count`,`set_count`,`set_now`,`small_poster_md5`,`small_poster_url`,`webp_poster_md5`,`webp_poster_url`,`status`,`mi_id`) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", data.Title, data.Actors, data.Director, data.Categorys, data.Category, data.Area, data.Description, data.ShortDescription, data.Flag, data.IsMultSet, data.IssueDate, data.Language, data.LastIssueDate, data.Md5, data.DownloadSrc, data.MidType, data.IndexCapital, data.IndexFuzzyCapital, data.IndexSpelling, data.IndexSymbol, data.IndexValidLength, data.PlayCount, data.PlayLength, data.PosterURL, data.Resolution, data.Score, data.ScoreCount, data.SetCount, data.SetNow, data.SmallPosterMd5, data.SmallPosterURL, data.WebpPosterMd5, data.WebpPosterURL, data.Status, data.MiId)
	if err != nil {
		return
	}
	insertId, err := rs.LastInsertId()
	if err != nil {
		return
	}
	id = uint32(insertId)
	return
}

func (c *videoDao) Update(data *VideoModel) (err error) {
	_, err = DefaultDB.Exec("UPDATE video SET `title`=?,`actors`=?,`director`=?,`categorys`=?,`category`=?,`area`=?,`desc`=?,`short_desc`=?,`flag`=?,`is_multset`=?,`issue_date`=?,`language`=?,`last_issue_date`=?,`md5`=?,`download_src`=?,`mid_type`=?,`index_capital`=?,`index_fuzzycapital`=?,`index_spelling`=?,`index_symbol`=?,`index_validlength`=?,`play_count`=?,`play_length`=?,`poster_url`=?,`resolution`=?,`score`=?,`score_count`=?,`set_count`=?,`set_now`=?,`small_poster_md5`=?,`small_poster_url`=?,`webp_poster_md5`=?,`webp_poster_url`=?,`status`=?,`mi_id`=? WHERE id=?", data.Title, data.Actors, data.Director, data.Categorys, data.Category, data.Area, data.Description, data.ShortDescription, data.Flag, data.IsMultSet, data.IssueDate, data.Language, data.LastIssueDate, data.Md5, data.DownloadSrc, data.MidType, data.IndexCapital, data.IndexFuzzyCapital, data.IndexSpelling, data.IndexSymbol, data.IndexValidLength, data.PlayCount, data.PlayLength, data.PosterURL, data.Resolution, data.Score, data.ScoreCount, data.SetCount, data.SetNow, data.SmallPosterMd5, data.SmallPosterURL, data.WebpPosterMd5, data.WebpPosterURL, data.Status, data.MiId, data.Id)

	return
}
