package handlers

import (
	"Projects/Car24/car24_api_gateway/api/http"
	"Projects/Car24/car24_api_gateway/genproto/order_service"
	"Projects/Car24/car24_api_gateway/pkg/util"
	"context"

	"github.com/gin-gonic/gin"
)

// CreateModel godoc
// @ID create_tarif
// @Router /tarif [POST]
// @Summary Create Tarif
// @Description  Create Tarif
// @Tags Tarif
// @Accept json
// @Produce json
// @Param profile body order_service.CreateTarif true "CreateTarif"
// @Success 200 {object} http.Response{data=order_service.Tarif} "GetTarifBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateTarif(c *gin.Context) {
	var tarif order_service.CreateTarif

	err := c.ShouldBindJSON(&tarif)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.TarifService().Create(
		c.Request.Context(),
		&tarif,
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetTarifByID godoc
// @ID get_tarif_by_id
// @Router /tarif/{id} [GET]
// @Summary Get Model By ID
// @Description Get Model By ID
// @Tags Tarif
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=order_service.Tarif} "Tarif"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetTarifByID(c *gin.Context) {
	dId := c.Param("id")

	if !util.IsValidUUID(dId) {
		h.handleResponse(c, http.InvalidArgument, "car id is an invalid uuid")
		return
	}

	resp, err := h.services.TarifService().GetByID(
		context.Background(),
		&order_service.TarifPK{
			Id: dId,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteTarif godoc
// @ID delete_tarif
// @Router /tarif/{id} [DELETE]
// @Summary Delete Tarif
// @Description Delete Tarif
// @Tags Tarif
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=object{}} "Tarif data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteTarif(c *gin.Context) {

	dId := c.Param("id")

	if !util.IsValidUUID(dId) {
		h.handleResponse(c, http.InvalidArgument, "discount id is an invalid uuid")
		return
	}

	resp, err := h.services.TarifService().Delete(
		c.Request.Context(),
		&order_service.TarifPK{Id: dId},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}
