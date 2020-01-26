package tisp

import (
	"github.com/rs/xid"
)

type ID string

func generateNewID() ID {
	return ID(xid.New().String())
}
