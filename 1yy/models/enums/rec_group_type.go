package enums

type RecGroupType int

const (
	RecGroupTypeSlide RecGroupType = iota
	RecGroupTypeSpecial
	RecGroupTypeChannel
	RecGroupTypeBest
	RecGroupTypeRank
)
