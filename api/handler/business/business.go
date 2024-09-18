package handler

import (
	"api/api/token"
	"api/config"
	pb "api/genprotos/business"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create Business
// @Description Create Business
// @Tags Business
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Create body pb.CreateBusinessRequest true "Create"
// @Success 201 {object} string "Success"
// @Failure 400 {string} string "Error"
// @Router /business [post]
func (h *BusinessHandler) CreateBusiness(ctx *gin.Context) {
	req := &pb.CreateBusinessRequest{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c := config.Load()
	id, _ := token.GetIdFromToken(ctx.Request, &c)
	req.OwnerId = id

	_, err := h.Business.CreateBusiness(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Business Create Successfully"})
}

// @Summary Update Business
// @Description Update Business
// @Tags Business
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Business ID"
// @Param Update body pb.UpdateBusinessRequest true "Update"
// @Success 200 {object} string "Success"
// @Failure 400 {string} string "Error"
// @Router /business/{id} [put]
func (h *BusinessHandler) UpdateBusiness(ctx *gin.Context) {
	req := &pb.UpdateBusinessRequest{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	req.Id = ctx.Param("id")

	c := config.Load()
	id, _ := token.GetIdFromToken(ctx.Request, &c)
	req.OwnerId = id

	_, err := h.Business.UpdateBusiness(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Business Updated Successfully"})
}



// @Summary Delete Business
// @Description Delete Business
// @Tags Business
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Business ID"
// @Success 200 {object} string "Success"
// @Failure 400 {string} string "Error"
// @Router /business/{id} [delete]
func (h *BusinessHandler) DeleteBusiness(ctx *gin.Context) {
	req := &pb.DeleteBusinessRequest{}
	req.Id = ctx.Param("id")

	c := config.Load()
	id, _ := token.GetIdFromToken(ctx.Request, &c)
	req.OwnerId = id

	_, err := h.Business.DeleteBusiness(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Business Deleted Successfully"})
}



// @Summary Get Business by ID
// @Description Get Business by ID
// @Tags Business
// @Accept json
// @Produce json
// @Param id path string true "Business ID"
// @Success 200 {object} pb.GetByIdBusinessResponse
// @Failure 400 {string} string "Error"
// @Router /business/{id} [get]
func (h *BusinessHandler) GetBusinessById(ctx *gin.Context) {
	req := &pb.GetByIdBusinessRequest{}
	req.Id = ctx.Param("id")

	res, err := h.Business.GetByIdBusiness(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// @Summary Get All Businesses
// @Description Get All Businesses
// @Tags Business
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Success 200 {object} pb.GetAllBusinessesResponse
// @Failure 400 {string} string "Error"
// @Router /business/all [get]
func (h *BusinessHandler) GetAllBusinesses(ctx *gin.Context) {
	page := ctx.Query("page")
	if page == "" {
		page = "1"
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid page number")
		return
	}

	req := &pb.GetAllBusinessesRequest{Page: int32(pageInt)}
	res, err := h.Business.GetAllBusinesses(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, res)
}
