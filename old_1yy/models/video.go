package models

import (
	"database/sql"
	"fmt"
)

type Video struct {
	Id              uint64
	Channel         string
	Title           string
	OriginTitle     string
	Alias           string
	Poster          string
	Directors       string
	Actors          string
	Year            uint64
	Lang            string
	Genres          string
	Area            string
	Tags            string
	Summary         string
	Weight          int64
	HourlyPlayCount uint64
	TotalPlayCount  uint64
	Status          string
	UpdateTime      string
	CreateTime      string
	ModifiedTime    string
	MergeVid        uint64
	Items           []*Url
}

type ChannelVideos struct {
	Code   string
	Name   string
	Videos []*Video
}

func GetRankVideos(channels []string, num int) (map[string][]*Video, error) {
	var videos map[string][]*Video
	key := fmt.Sprintf("RankVideos/%d", num)
	data, exists := cache.Get(key)
	if exists {
		videos = data.(map[string][]*Video)
	} else {
		videos = make(map[string][]*Video)
		stmt, err := db.Prepare("SELECT id,title,poster,total_play_count FROM `video` WHERE status='NORMAL' AND channel=? ORDER BY total_play_count DESC LIMIT ?")
		if err != nil {
			return nil, err
		}
		for _, channel := range channels {
			videos[channel] = make([]*Video, 0, num)
			rows, err := stmt.Query(channel, num)
			if err != nil {
				return nil, err
			}
			defer rows.Close()
			for rows.Next() {
				v := &Video{}
				err = rows.Scan(&v.Id, &v.Title, &v.Poster, &v.TotalPlayCount)
				if err != nil {
					return nil, err
				}
				videos[channel] = append(videos[channel], v)
			}
		}
		cache.Set(key, videos, 0)
	}
	return videos, nil
}

func GetVideo(id uint64) (*Video, error) {
	var video *Video
	key := fmt.Sprintf("video/id:%d", id)
	data, exists := cache.Get(key)
	if exists {
		video = data.(*Video)
	} else {
		row := db.QueryRow("SELECT id,channel,title,alias,directors,actors,year,lang,genres,area,tags,summary,poster,merge_vid,DATE(update_time),total_play_count,status FROM `video` WHERE `id`=? AND `status` IN ('NORMAL','MERGED')", id)
		video = &Video{}
		err := row.Scan(&video.Id, &video.Channel, &video.Title, &video.Alias, &video.Directors, &video.Actors, &video.Year, &video.Lang, &video.Genres, &video.Area, &video.Tags, &video.Summary, &video.Poster, &video.MergeVid, &video.UpdateTime, &video.TotalPlayCount, &video.Status)
		if err != nil {
			return nil, err
		}
		cache.Set(key, video, 0)
	}
	return video, nil
}

//分页获取分类视频列表
func GetGenreVideoList(genre string, page, num int) ([]*Video, int, error) {
	var (
		list  []*Video
		total int
	)
	key := [2]string{
		fmt.Sprintf("genre:%s/videos/page:%d,num:%d", genre, page, num),
		fmt.Sprintf("genre:%s/total/page:%d,num:%d", genre, page, num),
	}
	videoData, vExists := cache.Get(key[0])
	countData, cExists := cache.Get(key[1])
	if vExists && cExists {
		list = videoData.([]*Video)
		total = countData.(int)
	} else {
		list = make([]*Video, 0, num)
		rows, err := db.Query("SELECT  v.id,v.title,g.`name`,v.alias,v.directors,v.actors,v.year,v.lang,v.genres,v.area,v.summary,v.poster,v.status,update_time,create_time,modified_time FROM `video` v JOIN `genre` g ON g.name=v.channel OR g.name=v.genres WHERE g.code=? AND v.status='NORMAL' ORDER BY `update_time` DESC LIMIT ?,?", genre, (page-1)*num, num)
		if err != nil {
			return nil, 0, err
		}
		defer rows.Close()
		for rows.Next() {
			v := &Video{}
			err = rows.Scan(&v.Id, &v.Title, &v.Channel, &v.Alias, &v.Directors, &v.Actors, &v.Year, &v.Lang, &v.Genres, &v.Area, &v.Summary, &v.Poster, &v.Status, &v.UpdateTime, &v.CreateTime, &v.ModifiedTime)
			if err != nil {
				return nil, 0, err
			}
			list = append(list, v)
		}

		row := db.QueryRow("SELECT COUNT(*) FROM `video` v JOIN `genre` g ON g.name=v.channel OR g.name=v.genres WHERE g.code=? AND v.status='NORMAL'", genre)
		err = row.Scan(&total)
		if err != nil {
			return nil, total, err
		}
		cache.Set(key[0], list, 0)
		cache.Set(key[1], total, 0)
	}
	return list, total, nil
}

//获取首页频道视频
func GetIndexChannelVideos(num int) ([]*ChannelVideos, error) {
	var (
		list []*ChannelVideos
		key  = fmt.Sprintf("index/videos/num:%d", num)
	)
	data, exists := cache.Get(key)
	if exists {
		list = data.([]*ChannelVideos)
	} else {
		list = make([]*ChannelVideos, 0, num)
		stmt, err := db.Prepare("SELECT id,channel,title,alias,directors,actors,year,lang,genres,area,poster FROM `video` WHERE `channel`=? AND status='NORMAL' ORDER BY update_time DESC LIMIT ?")
		if err != nil {
			return list, err
		}
		genres, err := getNavGenres()
		if err != nil {
			return list, err
		}
		for _, genre := range genres {
			rows, err := stmt.Query(genre.Name, num)
			if err != nil {
				return list, err
			}
			defer rows.Close()
			videos := make([]*Video, 0)

			for rows.Next() {
				v := &Video{}
				err := rows.Scan(&v.Id, &v.Channel, &v.Title, &v.Alias, &v.Directors, &v.Actors, &v.Year, &v.Lang, &v.Genres, &v.Area, &v.Poster)
				if err != nil {
					return nil, err
				}
				videos = append(videos, v)
			}

			video := &ChannelVideos{
				Code:   genre.Code,
				Name:   genre.Name,
				Videos: videos,
			}
			list = append(list, video)
		}
		cache.Set(key, list, 0)
	}
	return list, nil
}

//根据关键字获取视频列表
func SearchVideoList(keyword string, page, num int) ([]*Video, int, error) {
	keyword = fmt.Sprintf("%%%s%%", keyword)

	//获取所有影片
	list := make([]*Video, 0, num)
	key := fmt.Sprintf("search/%s/videos/%d/%d", keyword, page, num)
	data, exists := cache.Get(key)
	if exists {
		list = data.([]*Video)
	} else {
		list = make([]*Video, 0, num)
		rows, err := db.Query("SELECT  id,channel,title,alias,directors,actors,year,lang,genres,area,poster FROM `video` WHERE (title LIKE ? OR actors LIKE ? OR tags LIKE ? OR alias LIKE ?) AND status='NORMAL' ORDER BY `update_time` DESC,`total_play_count` DESC LIMIT ?,?", keyword, keyword, keyword, keyword, (page-1)*num, num)
		if err != nil {
			return nil, 0, err
		}
		defer rows.Close()
		for rows.Next() {
			v := &Video{}
			err = rows.Scan(&v.Id, &v.Channel, &v.Title, &v.Alias, &v.Directors, &v.Actors, &v.Year, &v.Lang, &v.Genres, &v.Area, &v.Poster)
			if err != nil {
				return nil, 0, err
			}
			list = append(list, v)
		}

		cache.Set(key, list, 0)
	}

	//获取影片总数
	total := 0
	key = fmt.Sprintf("search/%s/total", keyword, page, num)
	data, exists = cache.Get(key)
	if exists {
		total = data.(int)
	} else {
		row := db.QueryRow("SELECT COUNT(*) FROM `video` WHERE (title LIKE ? OR actors LIKE ? OR tags LIKE ? OR alias LIKE ?) AND status='NORMAL'", keyword, keyword, keyword, keyword)
		err := row.Scan(&total)
		if err != nil {
			return nil, 0, err
		}
		cache.Set(key, total, 0)
	}
	return list, total, nil
}

//获取分类视频列表
func GetTypeVideoList(urlType, channel string, startTS, endTS int64, page, num int) ([]*Video, int, error) {
	videos := make([]*Video, 0, num)
	total := 0
	rows, err := db.Query("SELECT  v.id,v.channel,v.title,v.alias,v.directors,v.actors,v.year,v.lang,v.genres,v.area,v.summary,v.origin_poster,v.update_time,v.create_time,v.modified_time,v.status FROM video v JOIN url u ON u.vid=v.id and u.sn=1 WHERE v.status IN ('NORMAL','DELETED','MERGED') AND u.type=? AND channel=? AND v.modified_time>=FROM_UNIXTIME(?) AND v.modified_time<FROM_UNIXTIME(?) limit ?,?", urlType, channel, startTS, endTS, (page-1)*num, num)
	if err != nil {
		return videos, 0, err
	}
	defer rows.Close()
	for rows.Next() {
		v := &Video{}
		err = rows.Scan(&v.Id, &v.Channel, &v.Title, &v.Alias, &v.Directors, &v.Actors, &v.Year, &v.Lang, &v.Genres, &v.Area, &v.Summary, &v.Poster, &v.UpdateTime, &v.CreateTime, &v.ModifiedTime, &v.Status)
		if err != nil {
			return videos, 0, err
		}
		urls, err := GetVideoTypeUrls(v.Id, urlType)
		if err != nil {
			fmt.Println(err)
			continue
		}
		//v.Poster = PosterUrl(v.Poster)
		v.Items = urls
		videos = append(videos, v)
	}

	row := db.QueryRow("SELECT COUNT(*) FROM video v JOIN url u ON u.vid=v.id and u.sn=1 WHERE v.status IN ('NORMAL','DELETED','MERGED') AND u.type=? AND channel=? AND v.modified_time>=FROM_UNIXTIME(?) AND v.modified_time<FROM_UNIXTIME(?)", urlType, channel, startTS, endTS)
	err = row.Scan(&total)
	if err != nil {
		return videos, 0, err
	}
	return videos, total, nil
}

func GetVideoCount() ([2]int, error) {
	var count [2]int
	key := "VideoCount"
	data, exists := cache.Get(key)
	if exists {
		count = data.([2]int)
	} else {
		row := db.QueryRow("SELECT (SELECT COUNT(*) FROM video WHERE status='NORMAL' AND update_time BETWEEN CURDATE() AND ADDDATE(CURDATE(), 1)),(SELECT COUNT(*) FROM video WHERE status='NORMAL')")
		err := row.Scan(&count[0], &count[1])
		if err != nil {
			return count, err
		}
		cache.Set(key, count, 0)
	}
	return count, nil
}

func GetRecommendVideos(video *Video, num int) ([]*Video, error) {
	var videos []*Video
	key := fmt.Sprintf("RecommendVideos/%d/%d", num, video.Id)
	data, exists := cache.Get(key)
	if exists {
		videos = data.([]*Video)
	} else {
		videos = make([]*Video, 0, num)
		rows, err := db.Query("SELECT id,title,poster FROM video WHERE genres=? ORDER BY RAND() LIMIT ?", video.Genres, num)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			v := &Video{}
			err = rows.Scan(&v.Id, &v.Title, &v.Poster)
			if err != nil {
				return nil, err
			}
			videos = append(videos, v)
		}
		cache.Set(key, videos, 0)
	}
	return videos, nil
}

//获取分类热门视频，若genre为空串，则返回所有分类热门视频
func GetHotVideos(genre string, num int) ([]*Video, error) {
	var videos []*Video
	key := fmt.Sprintf("HotVideos/%s/%d", genre, num)
	data, exists := cache.Get(key)
	if exists {
		videos = data.([]*Video)
	} else {
		videos = make([]*Video, 0, num)
		var (
			err  error
			rows *sql.Rows
		)
		if genre == "" {
			rows, err = db.Query("SELECT id,title,poster,hourly_play_count FROM `video` WHERE status='NORMAL' ORDER BY hourly_play_count DESC LIMIT ?", num)
		} else {
			rows, err = db.Query("SELECT id,title,poster,hourly_play_count FROM `video` WHERE status='NORMAL' AND genres=? OR channel=? ORDER BY hourly_play_count DESC LIMIT ?", genre, genre, num)
		}
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			v := &Video{}
			err = rows.Scan(&v.Id, &v.Title, &v.Poster, &v.HourlyPlayCount)
			if err != nil {
				return nil, err
			}
			videos = append(videos, v)
		}
		cache.Set(key, videos, 0)
	}
	return videos, nil
}

func GetSitemapVideos(page, num int) (videos []*Video, err error) {
	key := fmt.Sprintf("SiteMapVideos/%d/%d", page, num)
	data, exists := cache.Get(key)
	if exists {
		videos = data.([]*Video)
	} else {
		var rows *sql.Rows
		offset := (page - 1) * num
		videos = make([]*Video, 0, num)
		rows, err = db.Query("SELECT id,title,summary,create_time,modified_time,hourly_play_count,total_play_count FROM `video` WHERE status='NORMAL' ORDER BY modified_time DESC LIMIT ?,?", offset, num)
		if err == nil {
			for rows.Next() {
				video := &Video{}
				rows.Scan(&video.Id, &video.Title, &video.Summary, &video.CreateTime, &video.ModifiedTime, &video.HourlyPlayCount, &video.TotalPlayCount)
				videos = append(videos, video)
			}
		}
	}
	return
}
