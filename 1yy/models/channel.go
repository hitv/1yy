package models

import (
	"time"

	"hi.tv/1yy/models/enums"
)

var ChannelDao _channelDao

type ChannelModel struct {
	Id        int64  `xorm:"pk autoincr"`
	Code      string `xorm:"index"`
	Name      string
	Parent    int64
	IsSub     bool
	IsRec     bool
	IsFilter  bool
	Status    enums.Status `xorm:"index"`
	Type      int64
	MiId      int64     `xorm:"index"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}

func (m *ChannelModel) TableName() string {
	return "channel"
}

type _channelDao struct{}

func (*_channelDao) GetAllChannels() (data []ChannelModel, err error) {
	err = db.Where("`parent`=0 AND status=?", enums.StatusEnabled).Find(&data)
	return
}

func (*_channelDao) GetByCode(code string) (data *ChannelModel, err error) {
	data = &ChannelModel{}
	exist, err := db.Where("code=?", code).Get(data)
	if err == nil && !exist {
		err = ErrNotExist
	}
	return
}

func (*_channelDao) GetByMiId(miId int64) (data *ChannelModel, err error) {
	data = &ChannelModel{}
	exist, err := db.Where("mi_id=? AND status=?", miId, enums.StatusEnabled).Get(data)
	if err == nil && !exist {
		err = ErrNotExist
	}
	return
}

func (*_channelDao) Insert(data *ChannelModel) (int64, error) {
	return db.Insert(data)
}

func (*_channelDao) Update(data *ChannelModel) (int64, error) {
	return db.Update(data)
}
