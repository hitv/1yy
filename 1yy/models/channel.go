package models

import "github.com/hitv/1yy/models/enums"

var ChannelDao = &channelDao{}

type ChannelModel struct {
	Id       uint32
	Code     string
	Name     string
	Parent   uint32
	IsSub    string
	IsRec    string
	IsFilter string
	Status   string
	Type     int32
	MiId     uint32
}

type channelDao struct{}

func (_ *channelDao) GetAllFilter() (data []*ChannelModel, err error) {
	rows, err := DefaultDB.Query("SELECT `id`,`code`,`name`,`parent`,`is_sub`,`is_rec`,`is_filter`,`status`,`type`,`mi_id` FROM `channel` WHERE `parent`=0 AND status=?", enums.StatusEnabled.String())
	if err != nil {
		return
	}
	defer rows.Close()

	data = make([]*ChannelModel, 0)
	for rows.Next() {
		channel := &ChannelModel{}
		err = rows.Scan(&channel.Id, &channel.Code, &channel.Name, &channel.Parent, &channel.IsSub, &channel.IsRec, &channel.IsFilter, &channel.Status, &channel.Type, &channel.MiId)
		data = append(data, channel)
	}

	return
}

func (_ *channelDao) GetByMiId(miId uint32) (data *ChannelModel, err error) {
	data = &ChannelModel{}
	row := DefaultDB.QueryRow("SELECT `id`,`code`,`name`,`parent`,`is_sub`,`is_rec`,`is_filter`,`status`,`type`,`mi_id` FROM `channel` WHERE mi_id=?", miId)
	err = row.Scan(&data.Id, &data.Code, &data.Name, &data.Parent, &data.IsSub, &data.IsRec, &data.IsFilter, &data.Status, &data.Type, &data.MiId)
	return
}

func (_ *channelDao) Insert(data *ChannelModel) (id uint32, err error) {
	rs, err := DefaultDB.Exec("INSERT INTO `channel`(code,name,parent,is_sub,is_rec,is_filter,type,mi_id) VALUES(?,?,?,?,?,?,?,?)", data.Code, data.Name, data.Parent, data.IsSub, data.IsRec, data.IsFilter, data.Type, data.MiId)
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
