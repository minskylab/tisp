package tisp

import (
	"github.com/rs/xid"
)

type ID string

func GenerateNewID() ID {
	return ID(xid.New().String())
}
