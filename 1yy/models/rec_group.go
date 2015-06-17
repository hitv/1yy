package models

import (
	"time"

	"hi.tv/1yy/models/enums"
)

var RecGroupDao *_recGroupDao

type RecGroupModel struct {
	Id        int64
	ChannelId int64
	Title     string
	SubTitle  string
	TargetURL string               `xorm:"'target_url'"`
	Special   *RecGroupItemModel   `xorm:"-"`
	Items     []*RecGroupItemModel `xorm:"-"`
	Status    enums.Status         `xorm:"index"`
	CreatedAt time.Time            `xorm:"created"`
	UpdatedAt time.Time            `xorm:"updated"`
}

func (*RecGroupModel) TableName() string {
	return "rec_group"
}

type _recGroupDao struct{}

func (*_recGroupDao) GetByChannelId(channelId int64) (group *RecGroupModel, err error) {
	group = &RecGroupModel{}
	exist, err := db.Where("channel_id=?", channelId).Get(group)
	if err == nil && !exist {
		err = ErrNotExist
	}
	return
}

func (*_recGroupDao) GetByTargetURL(targetURL string) (group *RecGroupModel, err error) {
	group = &RecGroupModel{}
	exist, err := db.Where("target_url=?", targetURL).Get(group)
	if err == nil && !exist {
		err = ErrNotExist
	}
	return
}

func (*_recGroupDao) Insert(group *RecGroupModel) (int64, error) {
	return db.Insert(group)
}

func (*_recGroupDao) Update(group *RecGroupModel) (int64, error) {
	return db.Update(group)
}
