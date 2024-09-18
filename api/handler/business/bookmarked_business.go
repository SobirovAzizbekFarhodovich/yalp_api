package handler

import (
	"api/api/token"
	"api/config"
	pb "api/genprotos/business"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Create Bookmarked Business
// @Description Create a bookmarked business for a user
// @Tags Bookmarked Businesses
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Create body pb.CreateBookmarked true "Create"
// @Success 201 {object} string "Success"
// @Failure 400 {string} string "Error"
// @Router /bookmarked-business [post]
func (h *BusinessHandler) CreateBookmarkedBus(ctx *gin.Context) {
	req := &pb.CreateBookmarkedBusinessRequest{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c := config.Load()
	id, _ := token.GetIdFromToken(ctx.Request, &c)
	req.UserId = id

	_, err := h.BookmarkedBusiness.CreateBookmarkedBusiness(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Business Bookmarked Successfully"})
}

// @Summary Delete Bookmarked Business
// @Description Delete a bookmarked business by ID
// @Tags Bookmarked Businesses
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Bookmark ID"
// @Success 200 {object} string "Success"
// @Failure 400 {string} string "Error"
// @Router /bookmarked-business/{id} [delete]
func (h *BusinessHandler) DeleteBookmarkedBus(ctx *gin.Context) {
	req := &pb.DeleteBookmarkedBusinessRequest{}
	req.Id = ctx.Param("id")

	c := config.Load()
	id, _ := token.GetIdFromToken(ctx.Request, &c)
	req.UserId = id

	_, err := h.BookmarkedBusiness.DeleteBookmarkedBusiness(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Bookmarked Business Deleted Successfully"})
}

// @Summary Get Bookmarked Business by ID
// @Description Get a specific bookmarked business by ID
// @Tags Bookmarked Businesses
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Bookmark ID"
// @Success 200 {object} pb.GetBookmarkedBusinessByIdResponse
// @Failure 400 {string} string "Error"
// @Router /bookmarked-business/{id} [get]
func (h *BusinessHandler) GetBookmarkedBusById(ctx *gin.Context) {
	req := &pb.GetBookmarkedBusinessByIdRequest{}
	req.Id = ctx.Param("id")

	c := config.Load()
	id, _ := token.GetIdFromToken(ctx.Request, &c)
	req.UserId = id

	resp, err := h.BookmarkedBusiness.GetBookmarkedBusinessById(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Get All Bookmarked Businesses
// @Description Get all bookmarked businesses for a user
// @Tags Bookmarked Businesses
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} pb.GetAllBookmarkedBusinessResponse
// @Failure 400 {string} string "Error"
// @Router /bookmarked-business [get]
func (h *BusinessHandler) GetAllBookmarkedBus(ctx *gin.Context) {
	req := &pb.GetAllBookmarkedBusinessRequest{}

	c := config.Load()
	id, _ := token.GetIdFromToken(ctx.Request, &c)
	req.UserId = id

	resp, err := h.BookmarkedBusiness.GetAllBookmarkedBusiness(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
