package groups

import (
	"fmt"

	"github.com/pol-rivero/pkgstate/lib/common"
	"github.com/pol-rivero/pkgstate/lib/common/log"
	"github.com/pol-rivero/pkgstate/lib/common/prompt"
	. "github.com/pol-rivero/pkgstate/lib/types"
)

func (l *GroupsTool) ApplyFixes(requestConfirmation bool, _ bool) ApplyFixesResult {
	for _, group := range l.MissingGroups {
		if !group.Exists && !createUserGroup(group.Name, requestConfirmation) {
			continue
		}
		if !addUserToGroup(l.UserName, group.Name, requestConfirmation) {
			continue
		}
		fmt.Printf("-> Successfully added user '%s' to group '%s'\n", l.UserName, group.Name)
	}
	return Done
}

func createUserGroup(groupName string, requestConfirmation bool) bool {
	if requestConfirmation && prompt.RequestInput("Yn", "Do you want to create the missing group '%s'?", groupName) == 'n' {
		log.Info("Skipping group '%s'", groupName)
		return false
	}
	err := common.RunCommand("sudo", "groupadd", groupName)
	if err != nil {
		log.Error("Failed to create group '%s': %v", groupName, err)
		return false
	}
	return true
}

func addUserToGroup(userName, groupName string, requestConfirmation bool) bool {
	if requestConfirmation && prompt.RequestInput("Yn", "Do you want to add the user '%s' to the group '%s'?", userName, groupName) == 'n' {
		log.Info("Skipping group '%s'", groupName)
		return false
	}
	err := common.RunCommand("sudo", "usermod", "-aG", groupName, userName)
	if err != nil {
		log.Error("Failed to add user '%s' to group '%s': %v", userName, groupName, err)
		return false
	}
	return true
}
