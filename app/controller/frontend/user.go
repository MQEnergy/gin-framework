package frontend

import (
	"github.com/MQEnergy/gin-framework/app/controller/base"
)

type UserController struct {
	*base.Controller
}

var User = &UserController{}
