package common

import (
	"github.com/gin-gonic/gin"
	"mqenergy-go/app/controller/base"
	"mqenergy-go/pkg/response"
	"net/http"
)

type Controller struct {
	base.Controller
}

var Common = Controller{}

// Ping 心跳
func (c Controller) Ping(ctx *gin.Context) {
	response.ResponseJson(ctx, http.StatusOK, response.Success, "Pong!", "")
}
