package handler

import (
	"api/api/token"
	"api/config"
	pb "api/genprotos/business"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Create Business Photo
// @Description Add a new photo for a business
// @Tags Business Photos
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Create body pb.CreateBusinessPhotosRequest true "Create"
// @Success 201 {object} string "Success"
// @Failure 400 {string} string "Error"
// @Router /business-photos [post]
func (h *BusinessHandler) CreateBusinessPhotos(ctx *gin.Context) {
	req := &pb.CreateBusinessPhotosRequest{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c := config.Load()
	id, _ := token.GetIdFromToken(ctx.Request, &c)
	req.OwnerId = id

	_, err := h.BusinessPhotos.CreateBusinessPhotos(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Business Photo Created Successfully"})
}

// @Summary Update Business Photo
// @Description Update an existing business photo
// @Tags Business Photos
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Update body pb.UpdateBusinessPhotosRequest true "Update"
// @Success 200 {object} string "Success"
// @Failure 400 {string} string "Error"
// @Router /business-photos [put]
func (h *BusinessHandler) UpdateBusinessPhotos(ctx *gin.Context) {
	req := &pb.UpdateBusinessPhotosRequest{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c := config.Load()
	id, _ := token.GetIdFromToken(ctx.Request, &c)
	req.OwnerId = id

	_, err := h.BusinessPhotos.UpdateBusinessPhotos(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Business Photo Updated Successfully"})
}

// @Summary Delete Business Photo
// @Description Delete a specific business photo by ID
// @Tags Business Photos
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Photo ID"
// @Success 200 {object} string "Success"
// @Failure 400 {string} string "Error"
// @Router /business-photos/{id} [delete]
func (h *BusinessHandler) DeleteBusinessPhotos(ctx *gin.Context) {
	req := &pb.DeleteBusinessPhotosRequest{}
	req.Id = ctx.Param("id")

	c := config.Load()
	id, _ := token.GetIdFromToken(ctx.Request, &c)
	req.OwnerId = id

	_, err := h.BusinessPhotos.DeleteBusinessPhotos(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Business Photo Deleted Successfully"})
}

// @Summary Get Business Photos by Business ID
// @Description Get all photos for a specific business
// @Tags Business Photos
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param businessId path string true "Business ID"
// @Success 200 {object} pb.GetBusinessIdResponse
// @Failure 400 {string} string "Error"
// @Router /business-photos/{businessId} [get]
func (h *BusinessHandler) GetByBusinessId(ctx *gin.Context) {
	req := &pb.GetBusinessIdRequest{}
	req.BusinessId = ctx.Param("businessId")

	resp, err := h.BusinessPhotos.GetByBusinessId(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Get Business Photos by Owner
// @Description Get all photos uploaded by the current user (owner)
// @Tags Business Photos
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} pb.GetBusinessPhotosByOwnerResponse
// @Failure 400 {string} string "Error"
// @Router /business-photos/owner [get]
func (h *BusinessHandler) GetBusinessPhotosByOwner(ctx *gin.Context) {
	req := &pb.GetBusinessPhotosByOwnerRequest{}

	c := config.Load()
	id, _ := token.GetIdFromToken(ctx.Request, &c)
	req.OwnerId = id

	resp, err := h.BusinessPhotos.GetBusinessPhotosByOwner(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
