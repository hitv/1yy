package enums

const (
	StatusEnabled  Status = "ENABLED"
	StatusDisabled        = "DISABLED"
	StatusDeleted         = "DELETED"
)

type Status string

func (s Status) String() string {
	return string(s)
}

func (s Status) Humanize() string {
	switch s {
	case StatusEnabled:
		return "启用"
	case StatusDisabled:
		return "禁用"
	case StatusDeleted:
		return "删除"
	}

	return "未知"
}
