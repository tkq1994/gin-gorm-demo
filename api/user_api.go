package api

import (
	"github.com/gin-gonic/gin"
	"iTypeService/service"
)

type UserApi struct {
	BaseApi
	Service *service.UserService
}

func NewUserApi() UserApi {
	return UserApi{
		BaseApi: NewBaseApi(),
		Service: service.NewUserService(),
	}
}

func (m UserApi) Login(c *gin.Context) {
	return
}

func (m UserApi) GetUserById(c *gin.Context) {
	return
}
