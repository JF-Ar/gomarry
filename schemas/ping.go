package schemas

import (
	"gorm.io/gorm"
)

type PingPong struct {
	gorm.Model
	Ping string
}
