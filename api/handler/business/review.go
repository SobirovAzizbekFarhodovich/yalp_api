package handler

import (
	"api/api/token"
	"api/config"
	pb "api/genprotos/business"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Create Review
// @Description Create a new review for a business by a user
// @Tags Reviews
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Create body pb.CreateReviewRequest true "Create Review"
// @Success 201 {object} pb.CreateReviewResponse "Success"
// @Failure 400 {string} string "Error"
// @Router /review [post]
func (h *BusinessHandler) CreateReview(ctx *gin.Context) {
	req := &pb.CreateReviewRequest{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.Review.CreateReview(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, resp)
}

// @Summary Update Review
// @Description Update an existing review by ID
// @Tags Reviews
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Update body pb.UpdateReviewRequest true "Update Review"
// @Success 200 {object} pb.UpdateReviewResponse "Success"
// @Failure 400 {string} string "Error"
// @Router /review/{id} [put]
func (h *BusinessHandler) UpdateReview(ctx *gin.Context) {
	req := &pb.UpdateReviewRequest{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.Review.UpdateReview(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Delete Review
// @Description Delete a review by ID
// @Tags Reviews
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Review ID"
// @Success 200 {object} string "Success"
// @Failure 400 {string} string "Error"
// @Router /review/{id} [delete]
func (h *BusinessHandler) DeleteReview(ctx *gin.Context) {
	req := &pb.DeleteReviewRequest{}
	req.Id = ctx.Param("id")

	_, err := h.Review.DeleteReview(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Review Deleted Successfully"})
}

// @Summary Get Own Reviews
// @Description Retrieve all reviews written by the authenticated user
// @Tags Reviews
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} pb.GetOwnReviewsResponse "Success"
// @Failure 400 {string} string "Error"
// @Router /review/user/{user_id} [get]
func (h *BusinessHandler) GetOwnReviews(ctx *gin.Context) {
	req := &pb.GetOwnReviewsRequest{}

	c := config.Load()
	id, _ := token.GetIdFromToken(ctx.Request, &c)
	req.UserId = id
	resp, err := h.Review.GetOwnReviews(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// @Summary Get Reviews by Business ID
// @Description Retrieve all reviews for a specific business
// @Tags Reviews
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param business_id path string true "Business ID"
// @Success 200 {object} pb.GetReviewByBusinessIdResponse "Success"
// @Failure 400 {string} string "Error"
// @Router /review/{business_id} [get]
func (h *BusinessHandler) GetReviewByBusinessId(ctx *gin.Context) {
	req := &pb.GetReviewByBusinessIdRequest{}
	req.BusinessId = ctx.Param("business_id")

	resp, err := h.Review.GetReviewByBusinessId(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
