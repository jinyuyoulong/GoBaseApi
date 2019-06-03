package models

import (
	"project-api/src/bootstrap/service"

	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func Initialize() {
	if engine == nil {
		container := service.GetDi().Container
		container.Invoke(func(db *xorm.Engine) {
			engine = db
		})
	}
}
