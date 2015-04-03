package models

import (
	"database/sql"
	"fmt"
)

type MtAlbum struct {
	Id          int                //图册id
	Title       string             //图册标题
	PrevAlbumId int                //上一个图册的id
	NextAlbumId int                //下一个图册的id
	CategoryId  int                //所属分类id
	PlayUri     string             //快播播放URI
	Images      []*MtImage         //包含的图片
	Recommend   []*MtCategoryAlbum //推荐图册
}

type MtCategoryAlbum struct {
	Id         int    //图册id
	Title      string //图册标题
	Cover      string //封面图
	Width      int    //封面图宽度
	Height     int    //封面图高度
	ImageCount int    //包含的图片个数
	Link       string
}

func GetAlbum(id int) (album *MtAlbum, err error) {
	key := fmt.Sprintf("Album/%d", id)
	data, exists := cache.Get(key)
	if exists {
		album = data.(*MtAlbum)
		return
	}

	row := db.QueryRow("SELECT id,cid,title,play_uri FROM `album` WHERE `id`=? AND `status`='NORMAL'", id)
	var (
		albumId, categoryId int
		albumTitle, playUri string
	)
	err = row.Scan(&albumId, &categoryId, &albumTitle, &playUri)
	if err != nil {
		return
	}
	//获取相册的上下个图册id
	prevId, nextId, err := getAlbumContextId(categoryId, albumId)
	if err != nil {
		return
	}
	recommend, err := getRecommend(albumId, 4)
	if err != nil {
		return
	}
	images, err := GetAlbumImages(albumId)
	if err != nil {
		return
	}

	album = &MtAlbum{
		Id:          albumId,
		Title:       albumTitle,
		CategoryId:  categoryId,
		PrevAlbumId: prevId,
		NextAlbumId: nextId,
		PlayUri:     playUri,
		Images:      images,
		Recommend:   recommend,
	}
	cache.Set(key, album, 0)
	return
}

func GetAlbums(categoryId, pageSize, page int) (albums []*MtCategoryAlbum, err error) {
	key := fmt.Sprintf("Albums/%d,%d,%d", categoryId, pageSize, page)
	data, exists := cache.Get(key)
	if exists {
		albums = data.([]*MtCategoryAlbum)
		return
	}
	offset := (page - 1) * pageSize
	albums = make([]*MtCategoryAlbum, 0, pageSize)
	rows, err := db.Query("SELECT a.id,a.title,a.cover,a.width,a.height,(SELECT COUNT(*) FROM image WHERE aid=a.id AND status='NORMAL') FROM album a JOIN category c ON c.id=a.cid WHERE c.id=? AND a.status='NORMAL' AND c.status='NORMAL' ORDER BY a.create_time DESC LIMIT ?,?", categoryId, offset, pageSize)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		album := &MtCategoryAlbum{}
		err = rows.Scan(&album.Id, &album.Title, &album.Cover, &album.Width, &album.Height, &album.ImageCount)
		if err != nil {
			return
		}
		albums = append(albums, album)
	}
	cache.Set(key, albums, 0)
	return
}
func getAlbumContextId(categoryId, albumId int) (prevId, nextId int, err error) {
	key := fmt.Sprintf("AlbumContextId/%d/%d", categoryId, albumId)
	data, exists := cache.Get(key)
	if exists {
		context := data.([2]int)
		prevId = context[0]
		nextId = context[1]
		return
	}
	partSqlFmt := "(SELECT a.id FROM album a JOIN category c ON c.id=a.cid WHERE c.id=? AND c.status='NORMAL' AND a.status='NORMAL' AND a.id%s? ORDER BY a.id DESC LIMIT 1)"
	sqlStr := fmt.Sprintf("SELECT %s,%s", fmt.Sprintf(partSqlFmt, ">"), fmt.Sprintf(partSqlFmt, "<"))
	row := db.QueryRow(sqlStr, categoryId, albumId, categoryId, albumId)
	prev, next := &sql.NullInt64{}, &sql.NullInt64{}
	err = row.Scan(prev, next)
	if err != nil {
		return
	}
	prevId = int(prev.Int64)
	nextId = int(next.Int64)
	cache.Set(key, [2]int{prevId, nextId}, 0)
	return
}
func getRecommend(albumId, num int) (albums []*MtCategoryAlbum, err error) {
	key := fmt.Sprintf("Recommend/%d/%d/%d", albumId, num)
	data, exists := cache.Get(key)
	if exists {
		albums = data.([]*MtCategoryAlbum)
		return
	}
	albums = make([]*MtCategoryAlbum, 0, num)
	rows, err := db.Query("SELECT a.id,a.title,a.cover,a.width,a.height FROM album a JOIN category c ON c.id=a.cid WHERE c.status='NORMAL' AND a.status='NORMAL' AND a.id!=? ORDER BY rand() ASC LIMIT ?", albumId, num)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		album := &MtCategoryAlbum{}
		err = rows.Scan(&album.Id, &album.Title, &album.Cover, &album.Width, &album.Height)
		if err != nil {
			return
		}
		albums = append(albums, album)
	}
	cache.Set(key, albums, 0)
	return
}
