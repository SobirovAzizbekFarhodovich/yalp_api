package handler

import (
	pb "api/genprotos/business"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create Location
// @Description Create a new location with latitude and longitude. If no address is provided, it will be fetched using the coordinates.
// @Tags Locations
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Create body pb.CreateLocationRequest true "Create"
// @Success 201 {object} pb.CreateLocationResponse "Success"
// @Failure 400 {string} string "Error"
// @Router /location [post]
func (h *BusinessHandler) CreateLocation(ctx *gin.Context) {
	req := &pb.CreateLocationRequest{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.Location.CreateLocation(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, resp)
}

// @Summary Delete Location
// @Description Delete a location by its ID
// @Tags Locations
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Location ID"
// @Success 200 {object} string "Success"
// @Failure 400 {string} string "Error"
// @Router /location/{id} [delete]
func (h *BusinessHandler) DeleteLocation(ctx *gin.Context) {
	req := &pb.DeleteLocationRequest{}
	req.Id = ctx.Param("id")

	_, err := h.Location.DeleteLocation(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Location Deleted Successfully"})
}

// @Summary Get Location by ID
// @Description Retrieve a location by its ID
// @Tags Locations
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Location ID"
// @Success 200 {object} pb.GetLocationByIdResponse
// @Failure 400 {string} string "Error"
// @Router /location/{id} [get]
func (h *BusinessHandler) GetLocationById(ctx *gin.Context) {
	req := &pb.GetLocationByIdRequest{}
	req.Id = ctx.Param("id")

	resp, err := h.Location.GetLocationById(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Get All Locations
// @Description Get a paginated list of locations
// @Tags Locations
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param page query int true "Page number"
// @Success 200 {object} pb.GetAllLocationResponse
// @Failure 400 {string} string "Error"
// @Router /location [get]
func (h *BusinessHandler) GetAllLocations(ctx *gin.Context) {
	req := &pb.GetAllLocationRequest{}

	page, _ := strconv.Atoi(ctx.Query("page"))
	req.Page = int32(page)

	resp, err := h.Location.GetAllLocations(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
