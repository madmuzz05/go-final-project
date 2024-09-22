package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	sysresponse "github.com/madmuzz05/go-final-project/pkg/helper/sys_response"
	"github.com/madmuzz05/go-final-project/service/order/dto/request"
)

func (h OrderHandler) GetOrders(ctx *gin.Context) {

	res, err := h.OrderUsecase.GetOrders(ctx)
	if err != nil {
		sysresponse.GetResponseJson(ctx, err.GetStatusCode(), err.GetMessage(), err)
		return
	}

	sysresponse.GetResponseJson(ctx, http.StatusOK, "Found Data", res)
}

func (h OrderHandler) GetOrder(ctx *gin.Context) {

	id := ctx.Param("order_id")
	if id == "" {
		sysresponse.GetResponseJson(ctx, http.StatusServiceUnavailable, "data tidak sesuai", "ID is required")
	}

	res, err := h.OrderUsecase.GetOrder(ctx, id)
	if err != nil {
		sysresponse.GetResponseJson(ctx, err.GetStatusCode(), err.GetMessage(), err)
		return
	}

	sysresponse.GetResponseJson(ctx, http.StatusOK, "Found Data", res)
}
func (h OrderHandler) DeleteOrder(ctx *gin.Context) {

	id := ctx.Param("order_id")
	if id == "" {
		sysresponse.GetResponseJson(ctx, http.StatusServiceUnavailable, "data tidak sesuai", "ID is required")
	}

	err := h.OrderUsecase.DeleteOrder(ctx, id)
	if err != nil {
		sysresponse.GetResponseJson(ctx, err.GetStatusCode(), err.GetMessage(), err)
		return
	}

	sysresponse.GetResponseJson(ctx, http.StatusOK, "Delete data success", nil)
}
func (h OrderHandler) StoreOrder(ctx *gin.Context) {
	id := ctx.Param("order_id")
	req := request.OrderDtoRequest{}
	if valErr := ctx.ShouldBind(&req); valErr != nil {
		sysresponse.GetResponseJson(ctx, http.StatusBadRequest, "Bad Request", valErr.Error())
		return
	}
	if id == "" && req.OrderedAt == "" {
		sysresponse.GetResponseJson(ctx, http.StatusBadRequest, "Bad Request", "OrderedAt is required")
		return
	}

	if len(req.Items) == 0 {
		sysresponse.GetResponseJson(ctx, http.StatusBadRequest, "Bad Request", "Item is required")
		return
	}

	res, err := h.OrderUsecase.StoreOrder(ctx, req, id)
	if err != nil {
		sysresponse.GetResponseJson(ctx, err.GetStatusCode(), err.GetMessage(), err)
		return
	}
	if id != "" {
		sysresponse.GetResponseJson(ctx, http.StatusOK, "Update data success", res)
		return
	}
	sysresponse.GetResponseJson(ctx, http.StatusOK, "Create data success", nil)
}
