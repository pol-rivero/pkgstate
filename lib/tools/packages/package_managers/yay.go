package packagemanagers

import "github.com/pol-rivero/pkgstate/lib/common"

type Yay struct{}

func (y *Yay) GetAllInstalledPackages() ([]string, error) {
	return common.RunCommandGetLines("yay", "-Qq")
}

func (y *Yay) GetExplicitlyInstalledPackages() ([]string, error) {
	return common.RunCommandGetLines("yay", "-Qqe")
}
