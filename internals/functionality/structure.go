package functionality

import (
	"sync"

	utils "github.com/Yadiiiig/blue-stats/internals/utils"
)

type Collection struct {
	sync.RWMutex
	Saved      *utils.ActiveUsers
	Connection *utils.Connection
}

type Body struct {
	Stream string `json:"stream"`
}

func NewCollection(c *utils.Utilities) *Collection {
	return &Collection{
		Saved:      c.ActiveUsers,
		Connection: c.Connection,
	}
}
