package enums

type ChannelType uint16

const (
	ChannelTypeTop ChannelType = 0x0000
	ChannelTypeSub             = 0x0001 << (iota - 1)
	ChannelTypeRec
	ChannelTypeFilter
)

func (t ChannelType) IsSub() string {
	if t&ChannelTypeSub > 0 {
		return "Y"
	}
	return "N"
}

func (t ChannelType) IsRec() string {
	if t&ChannelTypeRec > 0 {
		return "Y"
	}
	return "N"
}

func (t ChannelType) IsFilter() string {
	if t&ChannelTypeFilter > 0 {
		return "Y"
	}
	return "N"
}
