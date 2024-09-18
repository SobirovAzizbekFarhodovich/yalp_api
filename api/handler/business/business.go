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
// @Description Create a new business entry
// @Tags Business
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Create body pb.CreateBusiness true "Create Business"
// @Success 201 {object} pb.CreateBusinessResponse "Success"
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

	ctx.JSON(http.StatusCreated, gin.H{"message": "Business Created Successfully"})
}


// @Summary Update Business
// @Description Update an existing business entry
// @Tags Business
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Business ID"
// @Param Update body pb.UpdateBusiness true "Update Business"
// @Success 200 {object} pb.UpdateBusinessResponse "Success"
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
// @Security BearerAuth
// @Param id path string true "Business ID"
// @Success 200 {object} pb.GetByIdBusinessResponse
// @Failure 400 {string} string "Error"
// @Router /business/{id} [get]
func (h *BusinessHandler) GetBusinessById(ctx *gin.Context) {
	req := &pb.GetByIdBusinessRequest{
		Id: ctx.Param("id"),
	}

	res, err := h.Business.GetByIdBusiness(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// @Summary Get All Businesses
// @Description Get All Businesses
// @Tags Business
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number" default(1)
// @Param name query string false "Business name"
// @Param average_ratings query float32 false "Average ratings" // Use float32 instead of float
// @Param category query string false "Category"
// @Param contact_info query string false "Contact info"
// @Param hours_of_operation query string false "Hours of operation"
// @Param owner_id query string false "Owner ID"
// @Success 200 {object} pb.GetAllBusinessesResponse
// @Failure 400 {string} string "Error"
// @Router /business/all [get]
func (h *BusinessHandler) GetAllBusinesses(ctx *gin.Context) {
	page := ctx.Query("page")
	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}
	if pageInt == 0 {
		pageInt = 1
	}

	req := &pb.GetAllBusinessesRequest{
		Page:               int32(pageInt),
		Name:               ctx.Query("name"),
		AverageRatings:     parseFloat32(ctx.Query("average_ratings")),
		Category:           ctx.Query("category"),
		ContactInfo:        ctx.Query("contact_info"),
		HoursOfOperation:   ctx.Query("hours_of_operation"),
		OwnerId:            ctx.Query("owner_id"),
	}

	res, err := h.Business.GetAllBusinesses(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func parseFloat32(value string) float32 {
	if value == "" {
		return 0
	}

	f, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return 0
	}
	return float32(f)
}
