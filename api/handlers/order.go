package handlers

import (
	"Projects/Car24/car24_api_gateway/api/http"
	"Projects/Car24/car24_api_gateway/genproto/order_service"
	"Projects/Car24/car24_api_gateway/models"
	"Projects/Car24/car24_api_gateway/pkg/helper"
	"Projects/Car24/car24_api_gateway/pkg/util"
	"context"

	"github.com/gin-gonic/gin"
)

// CreateOrder godoc
// @ID create_order
// @Router /order [POST]
// @Summary Create Order
// @Description  Create Order
// @Tags Order
// @Accept json
// @Produce json
// @Param profile body order_service.CreateOrder true "CreateOrderRequestBody"
// @Success 200 {object} http.Response{data=order_service.Order} "GetOrderBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateOrder(c *gin.Context) {
	var order order_service.CreateOrder
	err := c.ShouldBindJSON(&order)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}
	resp, err := h.services.OrderService().Create(
		c.Request.Context(),
		&order,
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}
	h.handleResponse(c, http.Created, resp)
}

// GetOrderByID godoc
// @ID get_order_by_id
// @Router /order/{id} [GET]
// @Summary Get Order By ID
// @Description Get Order By ID
// @Tags Order
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=order_service.Order} "OrderBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetOrderByID(c *gin.Context) {
	orderId := c.Param("id")
	if !util.IsValidUUID(orderId) {
		h.handleResponse(c, http.InvalidArgument, "order id is an invalid uuid")
		return
	}
	resp, err := h.services.OrderService().GetByID(
		context.Background(),
		&order_service.OrderPrimaryKey{
			Id: orderId,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}
	h.handleResponse(c, http.OK, resp)
}

// GetOrderList godoc
// @ID get_order_list
// @Router /order [GET]
// @Summary Get Order List
// @Description Get Order List
// @Tags Order
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Success 200 {object} http.Response{data=order_service.GetListOrderResponse} "GetAllOrderResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetListOrder(c *gin.Context) {
	offset, err := h.getOffsetParam(c)
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	limit, err := h.getLimitParam(c)
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	resp, err := h.services.OrderService().GetList(
		context.Background(),
		&order_service.GetListOrderRequest{
			Limit:  int64(limit),
			Offset: int64(offset),
			Search: c.Query("search"),
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// @ID update_order
// @Router /order/{id} [PUT]
// @Summary Update Order
// @Description Update Order
// @Tags Order
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param profile body order_service.UpdateOrder true "UpdateOrderRequestBody"
// @Success 200 {object} http.Response{data=order_service.Order} "order data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateOrder(c *gin.Context) {

	var order order_service.UpdateOrder

	order.Id = c.Param("id")

	if !util.IsValidUUID(order.Id) {
		h.handleResponse(c, http.InvalidArgument, "order id is an invalid uuid")
		return
	}

	err := c.ShouldBindJSON(&order)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.OrderService().Update(
		c.Request.Context(),
		&order,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// PatchOrder godoc
// @ID patch_order
// @Router /order/{id} [PATCH]
// @Summary Patch Order
// @Description Patch Order
// @Tags Order
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param profile body models.UpdatePatch true "UpdatePatchRequestBody"
// @Success 200 {object} http.Response{data=order_service.Order} "Order data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdatePatchOrder(c *gin.Context) {

	var updatePatchOrder models.UpdatePatch

	err := c.ShouldBindJSON(&updatePatchOrder)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	updatePatchOrder.ID = c.Param("id")

	if !util.IsValidUUID(updatePatchOrder.ID) {
		h.handleResponse(c, http.InvalidArgument, "user id is an invalid uuid")
		return
	}

	structData, err := helper.ConvertMapToStruct(updatePatchOrder.Data)
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	resp, err := h.services.OrderService().UpdatePatch(
		c.Request.Context(),
		&order_service.UpdatePatchOrder{
			Id:     updatePatchOrder.ID,
			Fields: structData,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteOrder godoc
// @ID delete_order
// @Router /order/{id} [DELETE]
// @Summary Delete Order
// @Description Delete Order
// @Tags Order
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=object{}} "Order data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteOrder(c *gin.Context) {

	userId := c.Param("id")

	if !util.IsValidUUID(userId) {
		h.handleResponse(c, http.InvalidArgument, "user id is an invalid uuid")
		return
	}

	resp, err := h.services.OrderService().Delete(
		c.Request.Context(),
		&order_service.OrderPrimaryKey{Id: userId},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}
