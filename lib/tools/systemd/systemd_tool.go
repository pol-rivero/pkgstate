package systemd

type SystemdTool struct {
	SystemScope    bool
	UnitMismatches []SystemdUnitMismatch
}

func (l *SystemdTool) FriendlyProcessName() string {
	if l.SystemScope {
		return "get systemd units (system scope)"
	} else {
		return "get systemd units (user scope)"
	}
}

func NewSystemdTool(systemScope bool) *SystemdTool {
	return &SystemdTool{
		SystemScope: systemScope,
	}
}
