package groups

import (
	"os/user"

	"github.com/pol-rivero/pkgstate/lib/common"
	"github.com/pol-rivero/pkgstate/lib/common/config"
)

func (l *GroupsTool) GatherData(config *config.Config) error {
	desiredGroups := common.Sorted(config.UserGroups)
	unsortedCurrentGroups, err := getUserGroups()
	if err != nil {
		return err
	}
	currentGroups := common.Sorted(unsortedCurrentGroups)
	missingGroups, err := toMissingGroups(common.DifferenceOfOrderedSlices(desiredGroups, currentGroups))
	if err != nil {
		return err
	}
	l.MissingGroups = missingGroups

	currentUser, err := user.Current()
	if err != nil {
		return err
	}
	l.UserName = currentUser.Username
	return nil
}

func toMissingGroups(groupNames []string) ([]MissingGroup, error) {
	var missingGroups []MissingGroup
	for _, name := range groupNames {
		missingGroups = append(missingGroups, MissingGroup{
			Name:   name,
			Exists: userGroupExists(name),
		})
	}
	return missingGroups, nil
}

func getUserGroups() ([]string, error) {
	output, err := common.RunCommandGetOutput("groups")
	if err != nil {
		return nil, err
	}
	groups := common.SplitAndTrim(output, " ")
	return groups, nil
}

func userGroupExists(groupName string) bool {
	_, err := common.RunCommandGetOutput("getent", "group", groupName)
	return err == nil
}
