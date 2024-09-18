package handler

import (
	business "api/genprotos/business"

	"google.golang.org/grpc"
)

type BusinessHandler struct {
	BookmarkedBusiness business.Bookmarked_BusinessesClient
	Business           business.BusinessClient
	BusinessPhotos     business.Business_PhotosClient
	Location           business.LocationClient
	Review             business.ReviewsClient
}

func NewBusinessHandler(businessConn *grpc.ClientConn) *BusinessHandler {
	return &BusinessHandler{
		Business:           business.NewBusinessClient(businessConn),
		BookmarkedBusiness: business.NewBookmarked_BusinessesClient(businessConn),
		BusinessPhotos:     business.NewBusiness_PhotosClient(businessConn),
		Location:           business.NewLocationClient(businessConn),
		Review:             business.NewReviewsClient(businessConn),
	}
}
