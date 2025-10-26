package fix

import (
	"github.com/pol-rivero/pkgstate/lib/common/config"
)

func Fix(noConfirm bool) {
	config := config.GetConfig()
	_ = config
}
