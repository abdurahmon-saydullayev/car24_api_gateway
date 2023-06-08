package handlers

import (
	"Projects/Car24/car24_api_gateway/api/http"
	"Projects/Car24/car24_api_gateway/genproto/order_service"
	"Projects/Car24/car24_api_gateway/pkg/util"
	"context"

	"github.com/gin-gonic/gin"
)

// CreateModel godoc
// @ID create_model
// @Router /model [POST]
// @Summary Create Model
// @Description  Create Model
// @Tags Model
// @Accept json
// @Produce json
// @Param profile body order_service.CreateModel true "CreateModel"
// @Success 200 {object} http.Response{data=order_service.Model} "GetModelBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateModel(c *gin.Context) {
	var model order_service.CreateModel

	err := c.ShouldBindJSON(&model)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.ModelService().Create(
		c.Request.Context(),
		&model,
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetModelByID godoc
// @ID get_model_by_id
// @Router /model/{id} [GET]
// @Summary Get Model By ID
// @Description Get Model By ID
// @Tags Model
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=order_service.Model} "Model"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetModelByID(c *gin.Context) {
	dId := c.Param("id")

	if !util.IsValidUUID(dId) {
		h.handleResponse(c, http.InvalidArgument, "car id is an invalid uuid")
		return
	}

	resp, err := h.services.ModelService().GetByID(
		context.Background(),
		&order_service.ModelPK{
			Id: dId,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteModel godoc
// @ID delete_model
// @Router /model/{id} [DELETE]
// @Summary Delete Model
// @Description Delete Model
// @Tags Model
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=object{}} "Model data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteModel(c *gin.Context) {

	dId := c.Param("id")

	if !util.IsValidUUID(dId) {
		h.handleResponse(c, http.InvalidArgument, "discount id is an invalid uuid")
		return
	}

	resp, err := h.services.ModelService().Delete(
		c.Request.Context(),
		&order_service.ModelPK{Id: dId},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}
