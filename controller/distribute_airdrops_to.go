package controller

import (
	"Lotso_Airdrop_Server/model"
	"Lotso_Airdrop_Server/model/base"
	"Lotso_Airdrop_Server/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// DistributeAirdropsTo 接收地址和空投数量，并立刻对其发送空投，不对mysql记录做任何修改
func DistributeAirdropsTo(c *gin.Context) {
	var param model.DistributesParam
	err := c.ShouldBind(&param)
	if err != nil {
		c.JSON(http.StatusOK, base.NewErrorResponse(err, base.WrongParams))
		return
	}

	if err = param.Validate(); err != nil {
		response := base.NewErrorResponse(err, base.InvalidAddress)
		c.JSON(http.StatusOK, response)
		return
	}
	c.JSON(http.StatusOK, service.DistributeAirdropsTo(&param))
}
