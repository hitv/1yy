// channel_info.go
package mivideo

/*
func (c *Channel) Save(parentId int64, channelType enums.ChannelType) (id int64, err error) {
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
*/
