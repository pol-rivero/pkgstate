package groups

import (
	"os/user"

	"github.com/pol-rivero/pkgstate/lib/common"
	"github.com/pol-rivero/pkgstate/lib/common/config"
)

func (l *GroupsTool) GenerateConfig(cfg *config.Config) error {
	currentUser, err := user.Current()
	if err != nil {
		return err
	}

	userGroups, err := getUserGroups()
	if err != nil {
		return err
	}

	// Filter out the group with the same name as the user
	var filteredGroups []string
	for _, group := range userGroups {
		if group != currentUser.Username {
			filteredGroups = append(filteredGroups, group)
		}
	}

	cfg.UserGroups = common.Sorted(filteredGroups)
	return nil
}
