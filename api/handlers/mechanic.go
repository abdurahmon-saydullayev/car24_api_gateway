package handlers

import (
	"Projects/Car24/car24_api_gateway/api/http"
	"Projects/Car24/car24_api_gateway/genproto/order_service"
	"Projects/Car24/car24_api_gateway/pkg/util"
	"context"

	"github.com/gin-gonic/gin"
)

// CreateMechanic godoc
// @ID create_mechanic
// @Router /mechanic [POST]
// @Summary Create Mechanic
// @Description  Create Mechanic
// @Tags Mechanic
// @Accept json
// @Produce json
// @Param profile body order_service.CreateMechanic true "CreateMechanic"
// @Success 200 {object} http.Response{data=order_service.Mechanic} "GetMechanicBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateMechanic(c *gin.Context) {
	var mechanic order_service.CreateMechanic

	err := c.ShouldBindJSON(&mechanic)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.MechanicService().Create(
		c.Request.Context(),
		&mechanic,
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetMechanicByID godoc
// @ID get_mechanic_by_id
// @Router /mechanic/{id} [GET]
// @Summary Get Mechanic By ID
// @Description Get Mechanic By ID
// @Tags Mechanic
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=order_service.Mechanic} "Mechanic"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetMechanicByID(c *gin.Context) {
	dId := c.Param("id")

	if !util.IsValidUUID(dId) {
		h.handleResponse(c, http.InvalidArgument, "car id is an invalid uuid")
		return
	}

	resp, err := h.services.MechanicService().GetByID(
		context.Background(),
		&order_service.MechanicPK{
			Id: dId,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteMechanic godoc
// @ID delete_discount
// @Router /discount/{id} [DELETE]
// @Summary Delete Mechanic
// @Description Delete Mechanic
// @Tags Mechanic
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=object{}} "Mechanic data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteMechanic(c *gin.Context) {

	dId := c.Param("id")

	if !util.IsValidUUID(dId) {
		h.handleResponse(c, http.InvalidArgument, "discount id is an invalid uuid")
		return
	}

	resp, err := h.services.MechanicService().Delete(
		c.Request.Context(),
		&order_service.MechanicPK{Id: dId},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}
