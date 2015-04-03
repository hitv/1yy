// channel_info.go
package fetch

import (
	"fmt"

	"github.com/hitv/1yy/app/models"
	"github.com/hitv/1yy/app/models/enums"
)

type Channel struct {
	Id        uint32    `json:"id"`
	Name      string    `json:"name"`
	Count     uint32    `json:"count"`
	Type      int32     `json:"type"`
	RecSub    []Channel `json:"recsub"`
	Sub       []Channel `json:"sub"`
	SubFilter []Channel `json:"subfilter"`
}

func (c *Channel) Save(parentId uint32, channelType enums.ChannelType) (id uint32, err error) {
	channel, err := models.ChannelDao.GetByMiId(c.Id)
	if err != nil {
		if err != models.ErrNotFound {
			return
		}
		channel = &models.ChannelModel{
			Name:     c.Name,
			Parent:   parentId,
			IsSub:    channelType.IsSub(),
			IsRec:    channelType.IsRec(),
			IsFilter: channelType.IsFilter(),
			Type:     c.Type,
			MiId:     c.Id,
		}
		id, err = models.ChannelDao.Insert(channel)
	} else {
		id = channel.Id
	}

	for _, sub := range c.RecSub {
		_, err = sub.Save(id, enums.ChannelTypeRec)
		if err != nil {
			return
		}
	}

	for _, sub := range c.Sub {
		_, err = sub.Save(id, enums.ChannelTypeSub)
		if err != nil {
			return
		}
	}

	for _, filter := range c.SubFilter {
		_, err = filter.Save(id, enums.ChannelTypeFilter)
		if err != nil {
			return
		}
	}

	return
}

type ChannelInfo struct {
	Status int       `json:"status"`
	Data   []Channel `json:"data"`
}

func FetchChannelInfo(channelId int) (info *ChannelInfo, err error) {
	info = &ChannelInfo{}
	u := NewRequestURL(MI_HOST, "/tvservice/getchannelinfo3")
	u.AddParam("channelid", channelId)
	AddCommonParam(u)
	err = DoPost(u, info)
	return
}

func (c *ChannelInfo) Save() (err error) {
	if c.Status != 0 {
		err = fmt.Errorf("response status is: %d", c.Status)
		return
	}

	for _, channel := range c.Data {
		_, err = channel.Save(0, enums.ChannelTypeTop)
		if err != nil {
			return
		}
	}

	return
}
