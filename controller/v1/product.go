package v1

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"OnlineShopServer/config"
	"OnlineShopServer/constant"
	"OnlineShopServer/proto/ShopServer"
)

func CreateProduct(c *gin.Context) {
	var params struct {
		Account string `form:"account" binding:"required"`
	}
	if err := c.ShouldBindQuery(&params); err != nil {
		constant.ResponseWithData(c, http.StatusOK, constant.INVALID_PARAMS, constant.ERROR_MSG, err.Error())
		return
	}

	var body struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		constant.ResponseWithData(c, http.StatusOK, constant.INVALID_PARAMS, constant.ERROR_MSG, err.Error())
		return
	}

	req := ShopServer.CreateProductReq{
		Account: params.Account,
		Data: &ShopServer.CreateProductInfo{
			Title:       body.Title,
			Description: body.Description,
		},
	}

	conn, err := grpc.Dial(config.ShopServerUrl, grpc.WithInsecure())
	if err != nil {
		constant.ResponseWithData(c, http.StatusOK, constant.ERROR, err.Error(), nil)
		return
	}

	shopResp, err := ShopServer.NewShopServerClient(conn).CreateProduct(context.Background(), &req)
	if err != nil {
		constant.ResponseWithData(c, http.StatusOK, constant.ERROR, err.Error(), nil)
		return
	}

	constant.ResponseWithData(c, http.StatusOK, int(shopResp.Code), shopResp.Msg, nil)
	return
}

func GetProductsList(c *gin.Context) {
	var params struct {
		Account string `form:"account" binding:"required"`
	}

	if err := c.ShouldBindQuery(&params); err != nil {
		constant.ResponseWithData(c, http.StatusOK, constant.INVALID_PARAMS, constant.ERROR_MSG, err.Error())
	}

	req := ShopServer.GetProductsListReq{
		Account: params.Account,
	}
	conn, err := grpc.Dial(config.ShopServerUrl, grpc.WithInsecure())
	if err != nil {
		constant.ResponseWithData(c, http.StatusOK, constant.ERROR, err.Error(), nil)
		return
	}

	shopResp, err := ShopServer.NewShopServerClient(conn).GetProductsList(context.Background(), &req)
	if err != nil {
		constant.ResponseWithData(c, http.StatusOK, constant.ERROR, err.Error(), nil)
		return
	}

	log.Println("shopResp", shopResp)

	constant.ResponseWithData(c, http.StatusOK, int(shopResp.Code), shopResp.Msg, shopResp.Data)
	return
}

func GetProduct(c *gin.Context) {
	var params struct {
		ProductID string `form:"product_id" binding:"required"`
	}

	if err := c.ShouldBindQuery(&params); err != nil {
		constant.ResponseWithData(c, http.StatusOK, constant.INVALID_PARAMS, constant.ERROR_MSG, err.Error())
	}

	req := ShopServer.GetProductReq{
		ProductID: params.ProductID,
	}
	conn, err := grpc.Dial(config.ShopServerUrl, grpc.WithInsecure())
	if err != nil {
		constant.ResponseWithData(c, http.StatusOK, constant.ERROR, err.Error(), nil)
		return
	}

	shopResp, err := ShopServer.NewShopServerClient(conn).GetProduct(context.Background(), &req)
	if err != nil {
		constant.ResponseWithData(c, http.StatusOK, constant.ERROR, err.Error(), nil)
		return
	}

	log.Println("shopResp", shopResp)

	constant.ResponseWithData(c, http.StatusOK, int(shopResp.Code), shopResp.Msg, shopResp.Data)
	return
}

func UpdateProduct(c *gin.Context) {
	var params struct {
		Account string `form:"account" binding:"required"`
	}

	if err := c.ShouldBindQuery(&params); err != nil {
		constant.ResponseWithData(c, http.StatusOK, constant.INVALID_PARAMS, constant.ERROR_MSG, err.Error())
	}

	var body struct {
		ProductID   string `json:"product_id" binding:"required"`
		Title       string `json:"title" binding:"required"`
		Description string `json:"description" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		constant.ResponseWithData(c, http.StatusOK, constant.INVALID_PARAMS, constant.ERROR_MSG, err.Error())
		return
	}

	req := ShopServer.UpdateProductReq{
		Account: params.Account,
		Data: &ShopServer.UpdateProductReqInfo{
			ProductID:   body.ProductID,
			Title:       body.Title,
			Description: body.Description,
		},
	}

	conn, err := grpc.Dial(config.ShopServerUrl, grpc.WithInsecure())
	if err != nil {
		constant.ResponseWithData(c, http.StatusOK, constant.ERROR, err.Error(), nil)
		return
	}

	shopResp, err := ShopServer.NewShopServerClient(conn).UpdateProduct(context.Background(), &req)
	if err != nil {
		constant.ResponseWithData(c, http.StatusOK, constant.ERROR, err.Error(), nil)
		return
	}

	constant.ResponseWithData(c, http.StatusOK, int(shopResp.Code), shopResp.Msg, nil)
	return
}

func DeleteProduct(c *gin.Context) {
	var params struct {
		Account string `form:"account" binding:"required"`
	}

	if err := c.ShouldBindQuery(&params); err != nil {
		constant.ResponseWithData(c, http.StatusOK, constant.INVALID_PARAMS, constant.ERROR_MSG, err.Error())
	}

	var body struct {
		ProductID   string `json:"product_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		constant.ResponseWithData(c, http.StatusOK, constant.INVALID_PARAMS, constant.ERROR_MSG, err.Error())
		return
	}

	req := ShopServer.DeleteProductReq{
		Account: params.Account,
		Data: &ShopServer.DeleteProductReqInfo{
			ProductID:   body.ProductID,
		},
	}

	conn, err := grpc.Dial(config.ShopServerUrl, grpc.WithInsecure())
	if err != nil {
		constant.ResponseWithData(c, http.StatusOK, constant.ERROR, err.Error(), nil)
		return
	}

	shopResp, err := ShopServer.NewShopServerClient(conn).DeleteProduct(context.Background(), &req)
	if err != nil {
		constant.ResponseWithData(c, http.StatusOK, constant.ERROR, err.Error(), nil)
		return
	}

	constant.ResponseWithData(c, http.StatusOK, int(shopResp.Code), shopResp.Msg, nil)
	return
}

