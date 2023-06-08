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

// CreateCar godoc
// @ID create_car
// @Router /car [POST]
// @Summary Create Car
// @Description  Create Car
// @Tags Car
// @Accept json
// @Produce json
// @Param profile body order_service.CreateCar true "CreateCar"
// @Success 200 {object} http.Response{data=order_service.Car} "GetCarBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateCar(c *gin.Context) {
	var car order_service.CreateCar

	err := c.ShouldBindJSON(&car)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.CarService().Create(
		c.Request.Context(),
		&car,
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetCarByID godoc
// @ID get_car_by_id
// @Router /car/{id} [GET]
// @Summary Get Car By ID
// @Description Get Car By ID
// @Tags Car
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=order_service.Car} "Car"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetCarByID(c *gin.Context) {
	carId := c.Param("id")

	if !util.IsValidUUID(carId) {
		h.handleResponse(c, http.InvalidArgument, "car id is an invalid uuid")
		return
	}

	resp, err := h.services.CarService().GetByID(
		context.Background(),
		&order_service.CarPrimaryKey{
			Id: carId,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// GetCarList godoc
// @ID get_car_list
// @Router /car [GET]
// @Summary Get Car List
// @Description Get Car List
// @Tags Car
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Success 200 {object} http.Response{data=order_service.GetListCarResponse} "GetAllCarResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetCarList(c *gin.Context) {

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

	resp, err := h.services.CarService().GetList(
		context.Background(),
		&order_service.GetListCarRequest{
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

// @ID update_car
// @Router /car/{id} [PUT]
// @Summary Update Car
// @Description Update Car
// @Tags Car
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param profile body order_service.UpdateCar true "UpdateCarRequestBody"
// @Success 200 {object} http.Response{data=order_service.Car} "Car data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateCar(c *gin.Context) {

	var car order_service.UpdateCar

	car.Id = c.Param("id")

	if !util.IsValidUUID(car.Id) {
		h.handleResponse(c, http.InvalidArgument, "car id is an invalid uuid")
		return
	}

	err := c.ShouldBindJSON(&car)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.CarService().Update(
		c.Request.Context(),
		&car,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// PatchCar godoc
// @ID patch_car
// @Router /car/{id} [PATCH]
// @Summary Patch Car
// @Description Patch Car
// @Tags Car
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param profile body models.UpdatePatch true "UpdatePatchRequestBody"
// @Success 200 {object} http.Response{data=order_service.Car} "Car data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdatePatchCar(c *gin.Context) {

	var updatePatchCar models.UpdatePatch

	err := c.ShouldBindJSON(&updatePatchCar)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	updatePatchCar.ID = c.Param("id")

	if !util.IsValidUUID(updatePatchCar.ID) {
		h.handleResponse(c, http.InvalidArgument, "car id is an invalid uuid")
		return
	}

	structData, err := helper.ConvertMapToStruct(updatePatchCar.Data)
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	resp, err := h.services.CarService().UpdatePatch(
		c.Request.Context(),
		&order_service.UpdatePathCar{
			Id:     updatePatchCar.ID,
			Fields: structData,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteCar godoc
// @ID delete_car
// @Router /car/{id} [DELETE]
// @Summary Delete Car
// @Description Delete Car
// @Tags Car
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=object{}} "Car data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteCar(c *gin.Context) {

	carId := c.Param("id")

	if !util.IsValidUUID(carId) {
		h.handleResponse(c, http.InvalidArgument, "car id is an invalid uuid")
		return
	}

	resp, err := h.services.CarService().Delete(
		c.Request.Context(),
		&order_service.CarPrimaryKey{Id: carId},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}
