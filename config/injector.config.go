package config

import (
	"github.com/google/wire"
)

var db = wire.NewSet(InitDB)
