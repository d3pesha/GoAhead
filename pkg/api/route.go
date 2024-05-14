package api

import (
	"GoAhead/pkg/service"
	"github.com/gin-gonic/gin"
)

type CentralBankRoute struct {
	uc service.CentralBankUseCase
}

func NewCentralBankRoute(uc service.CentralBankUseCase) *CentralBankRoute {
	return &CentralBankRoute{
		uc: uc,
	}
}

func (r CentralBankRoute) Register(router *gin.Engine) {
	router.GET("/currency", r.GetCursByValue)
}
