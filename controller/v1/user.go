package v1

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"OnlineShopServer/config"
	"OnlineShopServer/constant"
	"OnlineShopServer/proto/UserServer"
)

func CreateUser(c *gin.Context) {
	var params struct {
		Account  string `form:"account" binding:"required"`
		Password string `form:"password" binding:"required"`
	}

	if err := c.ShouldBindQuery(&params); err != nil {
		constant.ResponseWithData(c, http.StatusOK, constant.INVALID_PARAMS, constant.ERROR_MSG, err.Error())
		return
	}

	req := UserServer.CreateUserReq{
		Account:  params.Account,
		Password: params.Password,
	}
	log.Println("user url", config.UserServerUrl)
	conn, err := grpc.Dial(config.UserServerUrl, grpc.WithInsecure())
	if err != nil {
		constant.ResponseWithData(c, http.StatusOK, constant.ERROR, err.Error(), nil)
		return
	}

	userResp, err := UserServer.NewUserServerClient(conn).CreateUser(context.Background(), &req)
	if err != nil {
		constant.ResponseWithData(c, http.StatusOK, constant.ERROR, err.Error(), nil)
		return
	}

	constant.ResponseWithData(c, http.StatusOK, int(userResp.Code), userResp.Msg, nil)
	return
}

func GetUser(c *gin.Context) {
	log.Println("get request")
	var params struct {
		Account  string `form:"account" binding:"required"`
		Password string `form:"password" binding:"required"`
	}
	if err := c.ShouldBindQuery(&params); err != nil {
		constant.ResponseWithData(c, http.StatusOK, constant.INVALID_PARAMS, constant.ERROR_MSG, err.Error())
		return
	}
	req := UserServer.GetUserReq{
		Account:  params.Account,
		Password: params.Password,
	}
	log.Println("user url", config.UserServerUrl)
	conn, err := grpc.Dial(config.UserServerUrl, grpc.WithInsecure())
	if err != nil {
		log.Println("conn err")
		constant.ResponseWithData(c, http.StatusOK, constant.ERROR, err.Error(), nil)
		return
	}
	log.Println("conn success", err, conn.GetState().String())
	userResp, err := UserServer.NewUserServerClient(conn).GetUser(context.Background(), &req)
	if err != nil {
		constant.ResponseWithData(c, http.StatusOK, constant.ERROR, err.Error(), nil)
		return
	}

	constant.ResponseWithData(c, http.StatusOK, int(userResp.Code), userResp.Msg, nil)
	return
}
