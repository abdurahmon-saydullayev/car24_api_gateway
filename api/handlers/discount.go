package handlers

import (
	"Projects/Car24/car24_api_gateway/api/http"
	"Projects/Car24/car24_api_gateway/genproto/order_service"
	"Projects/Car24/car24_api_gateway/pkg/util"
	"context"

	"github.com/gin-gonic/gin"
)

// CreateDiscount godoc
// @ID create_car
// @Router /discount [POST]
// @Summary Create Discount
// @Description  Create Discount
// @Tags Discount
// @Accept json
// @Produce json
// @Param profile body order_service.CreateDiscount true "CreateDiscount"
// @Success 200 {object} http.Response{data=order_service.Discount} "GetDiscountBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateDiscount(c *gin.Context) {
	var discount order_service.CreateDiscount

	err := c.ShouldBindJSON(&discount)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.DiscountService().Create(
		c.Request.Context(),
		&discount,
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetDiscountByID godoc
// @ID get_discount_by_id
// @Router /discount/{id} [GET]
// @Summary Get Discount By ID
// @Description Get Discount By ID
// @Tags Discount
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=order_service.Discount} "Discount"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetDiscountByID(c *gin.Context) {
	dId := c.Param("id")

	if !util.IsValidUUID(dId) {
		h.handleResponse(c, http.InvalidArgument, "car id is an invalid uuid")
		return
	}

	resp, err := h.services.DiscountService().GetByID(
		context.Background(),
		&order_service.DiscountPK{
			Id: dId,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteDiscount godoc
// @ID delete_discount
// @Router /discount/{id} [DELETE]
// @Summary Delete Discount
// @Description Delete Discount
// @Tags Discount
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=object{}} "Discount data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteDiscount(c *gin.Context) {

	dId := c.Param("id")

	if !util.IsValidUUID(dId) {
		h.handleResponse(c, http.InvalidArgument, "discount id is an invalid uuid")
		return
	}

	resp, err := h.services.DiscountService().Delete(
		c.Request.Context(),
		&order_service.DiscountPK{Id: dId},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}
