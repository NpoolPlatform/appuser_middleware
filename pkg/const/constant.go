package constant

import (
	"time"
)

type DTMAction struct {
	Action string
	Revert string
}

const (
	DBTimeout = 5 * time.Second

	CreateUser       = "CreateUser"
	CreateUserRevert = "CreateUserRevert"

	DefaultRowLimit = int32(100)
)
