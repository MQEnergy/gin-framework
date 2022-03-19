package base

import (
	"github.com/gin-gonic/gin"
	"mqenergy-go/pkg/response"
	"net/http"
)

type Controller struct{}

var Base = Controller{}

func (c Controller) Index(ctx *gin.Context) {
	response.ResponseJson(ctx, http.StatusOK, response.Success, "index", "")
}

func (c Controller) Create(ctx *gin.Context) {
	response.ResponseJson(ctx, http.StatusOK, response.Success, "create", "")
}

func (c Controller) Delete(ctx *gin.Context) {
	response.ResponseJson(ctx, http.StatusOK, response.Success, "delete", "")
}

func (c Controller) Update(ctx *gin.Context) {
	response.ResponseJson(ctx, http.StatusOK, response.Success, "update", "")
}

func (c Controller) View(ctx *gin.Context) {
	response.ResponseJson(ctx, http.StatusOK, response.Success, "view", "")
}
