package daily_refetch

import (
	callistodb "github.com/forbole/callisto/v4/database"
	"github.com/forbole/juno/v6/modules"
	"github.com/forbole/juno/v6/node"
)

var (
	_ modules.Module                   = &Module{}
	_ modules.PeriodicOperationsModule = &Module{}
)

type Module struct {
	node     node.Node
	database *callistodb.Db
}

// NewModule builds a new Module instance
func NewModule(
	node node.Node,
	database *callistodb.Db,
) *Module {
	return &Module{
		node:     node,
		database: database,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "daily refetch"
}
