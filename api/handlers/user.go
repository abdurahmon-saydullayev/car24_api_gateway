package handlers

import (
	"Projects/Car24/car24_api_gateway/api/http"
	"Projects/Car24/car24_api_gateway/genproto/client_service"
	"Projects/Car24/car24_api_gateway/models"
	"Projects/Car24/car24_api_gateway/pkg/helper"
	"Projects/Car24/car24_api_gateway/pkg/util"
	"context"

	"github.com/gin-gonic/gin"
)

// CreateClient godoc
// @ID create_user
// @Router /user [POST]
// @Summary Create User
// @Description  Create User
// @Tags User
// @Accept json
// @Produce json
// @Param profile body client_service.CreateClient true "CreateClient"
// @Success 200 {object} http.Response{data=client_service.Client} "GetUserBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateClient(c *gin.Context) {
	var user client_service.CreateClient

	err := c.ShouldBindJSON(&user)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.UserService().Create(
		c.Request.Context(),
		&user,
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetClientByID godoc
// @ID get_client_by_id
// @Router /client/{id} [GET]
// @Summary Get Client By ID
// @Description Get Client By ID
// @Tags Client
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=client_service.Client} "Client"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetClientByID(c *gin.Context) {
	userId := c.Param("id")

	if !util.IsValidUUID(userId) {
		h.handleResponse(c, http.InvalidArgument, "user id is an invalid uuid")
		return
	}

	resp, err := h.services.UserService().GetByID(
		context.Background(),
		&client_service.CLientPrimaryKey{
			Id: userId,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// GetClientList godoc
// @ID get_client_list
// @Router /client [GET]
// @Summary Get Client List
// @Description Get Client List
// @Tags Client
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Success 200 {object} http.Response{data=client_service.GetListClientResponse} "GetAllClientResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetClientList(c *gin.Context) {

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

	resp, err := h.services.UserService().GetList(
		context.Background(),
		&client_service.GetListClientRequest{
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

// @ID update_client
// @Router /client/{id} [PUT]
// @Summary Update Client
// @Description Update Client
// @Tags Client
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param profile body client_service.UpdateClient true "UpdateClientRequestBody"
// @Success 200 {object} http.Response{data=client_service.Client} "Client data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateClient(c *gin.Context) {

	var user client_service.UpdateClient

	user.Id = c.Param("id")

	if !util.IsValidUUID(user.Id) {
		h.handleResponse(c, http.InvalidArgument, "user id is an invalid uuid")
		return
	}

	err := c.ShouldBindJSON(&user)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.UserService().Update(
		c.Request.Context(),
		&user,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// PatchUser godoc
// @ID patch_client
// @Router /client/{id} [PATCH]
// @Summary Patch Client
// @Description Patch Client
// @Tags Client
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param profile body models.UpdatePatch true "UpdatePatchRequestBody"
// @Success 200 {object} http.Response{data=client_service.Client} "Client data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdatePatchClient(c *gin.Context) {

	var updatePatchUser models.UpdatePatch

	err := c.ShouldBindJSON(&updatePatchUser)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	updatePatchUser.ID = c.Param("id")

	if !util.IsValidUUID(updatePatchUser.ID) {
		h.handleResponse(c, http.InvalidArgument, "user id is an invalid uuid")
		return
	}

	structData, err := helper.ConvertMapToStruct(updatePatchUser.Data)
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	resp, err := h.services.UserService().UpdatePatch(
		c.Request.Context(),
		&client_service.UpdatePatchClient{
			Id:     updatePatchUser.ID,
			Fields: structData,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteClient godoc
// @ID delete_client
// @Router /client/{id} [DELETE]
// @Summary Delete Client
// @Description Delete Client
// @Tags Client
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=object{}} "Client data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteClient(c *gin.Context) {

	userId := c.Param("id")

	if !util.IsValidUUID(userId) {
		h.handleResponse(c, http.InvalidArgument, "user id is an invalid uuid")
		return
	}

	resp, err := h.services.UserService().Delete(
		c.Request.Context(),
		&client_service.CLientPrimaryKey{Id: userId},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}
