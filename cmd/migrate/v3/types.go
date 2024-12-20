package v3

import (
	"github.com/forbole/callisto/v4/modules/actions"
	v3 "github.com/forbole/juno/v6/cmd/migrate/v3"
)

type Config struct {
	v3.Config `yaml:"-,inline"`

	// The following are there to support modules which config are present if they are enabled

	Actions *actions.Config `yaml:"actions"`
}
