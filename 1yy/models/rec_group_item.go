package models

import (
	"time"

	"hi.tv/1yy/models/enums"
)

var RecGroupItemDao *_recGroupItemDao

type RecGroupItemModel struct {
	Id        int64
	GroupId   int64 `xorm:"index"`
	Title     string
	SubTitle  string
	Poster    string
	Entity    string
	TargetURL string `xorm:"'target_url'"`
	Hint      string
	Status    enums.Status `xorm:"index"`
	MiURL     string       `xorm:"index 'mi_url'"`
	CreatedAt time.Time    `xorm:"created"`
	UpdatedAt time.Time    `xorm:"updated"`
}

type _recGroupItemDao struct{}

func (*_recGroupItemDao) GetByGroupIdAndMiURL(groupId int64, miURL string) (item *RecGroupItemModel, err error) {
	item = &RecGroupItemModel{}
	exist, err := db.Where("group_id=? AND mi_url=?", groupId, miURL).Get(item)
	if err == nil && !exist {
		err = ErrNotExist
	}
	return
}
func (*_recGroupItemDao) Insert(item *RecGroupItemModel) (int64, error) {
	return db.Insert(item)
}

func (*_recGroupItemDao) Update(item *RecGroupItemModel) (int64, error) {
	return db.Update(item)
}
