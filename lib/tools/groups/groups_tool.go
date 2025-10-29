package groups

type MissingGroup struct {
	Name   string
	Exists bool
}

type GroupsTool struct {
	UserName      string
	MissingGroups []MissingGroup
}

func (l *GroupsTool) FriendlyProcessName() string {
	return "get user groups"
}

func NewGroupsTool() *GroupsTool {
	return &GroupsTool{}
}
