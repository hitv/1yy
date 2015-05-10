package models

import (
	"hi.tv/1yy/models/enums"
)

var ChannelDao = &channelDao{}

type ChannelModel struct {
	Id       int64 `xorm:"pk autoincr"`
	Code     string
	Name     string
	Parent   int64
	IsSub    string
	IsRec    string
	IsFilter string
	Status   string
	Type     int64
	MiId     int64
}

func (m *ChannelModel) TableName() string {
	return "channel"
}

type channelDao struct{}

func (_ *channelDao) GetAllChannels() (data []ChannelModel, err error) {
	err = db.Where("`parent`=0 AND status=?", enums.StatusEnabled.String()).Find(&data)
	return
}

func (_ *channelDao) GetByMiId(miId int64) (data *ChannelModel, err error) {
	has, err := db.Where("mi_id=? AND status=?", miId, enums.StatusEnabled.String()).Get(data)
	if err == nil && !has {
		err = ErrNotExist
	}
	return
}

func (_ *channelDao) Insert(data *ChannelModel) (int64, error) {
	return db.Insert(data)
}
