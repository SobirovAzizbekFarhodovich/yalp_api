package api

import (
	handlerC "api/api/handler/business"

	_ "api/docs"
	_ "api/genprotos/auth"
	_ "api/genprotos/business"

	"api/api/middleware"
	"log"

	"github.com/casbin/casbin/v2"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
)

// @title Budgeting SYSTEM API
// @version 1.0
// @description Developing a platform that helps users track their spending, set a budget and manage their financial goals
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewGin( /*AuthConn, */ businessConn *grpc.ClientConn) *gin.Engine {
	business := handlerC.NewBusinessHandler(businessConn)

	router := gin.Default()

	enforcer, err := casbin.NewEnforcer("/home/sobirov/go/src/gitlab.com/yelp/api/api/model.conf", "/home/sobirov/go/src/gitlab.com/yelp/api/api/policy.csv")
	if err != nil {
		log.Fatal(err)
	}

	// sw := router.Group("/")
	router.Use(middleware.NewAuth(enforcer))

	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	biznes := router.Group("/business")
	{
		biznes.POST("", business.CreateBusiness)
		biznes.PUT("/:id", business.UpdateBusiness)
		biznes.GET("/:id", business.GetBusinessById)
		biznes.DELETE("/:id",business.DeleteBusiness)
		biznes.GET("/all",business.GetAllBusinesses)
	}
	bookmarked := router.Group("/bookmarked-business")
	{
		bookmarked.POST("", business.CreateBookmarkedBus)
		bookmarked.DELETE("/:id",business.DeleteBookmarkedBus)
		bookmarked.GET("/:id",business.GetBookmarkedBusById)
		bookmarked.GET("",business.GetAllBookmarkedBus)
	}
	BusinessPhotos := router.Group("/business-photo")
	{
		BusinessPhotos.POST("", business.CreateBusinessPhotos)
		BusinessPhotos.DELETE("/:id",business.DeleteBusinessPhotos)
		BusinessPhotos.PUT("/:id",business.UpdateBusinessPhotos)
		BusinessPhotos.GET("/:businessId",business.GetByBusinessId)
	}
	location := router.Group("/location")
	{
		location.POST("", business.CreateLocation)
		location.DELETE("/:id",business.DeleteLocation)
		location.GET("/:id",business.GetLocationById)
		location.GET("",business.GetAllLocations)
	}
	review := router.Group("/review")
	{
		review.POST("", business.CreateReview)
		review.PUT("/:id", business.UpdateReview)
		review.GET("/user/:user_id", business.GetOwnReviews)
		review.DELETE("/:id",business.DeleteReview)
		review.GET("/:business_id",business.GetAllBusinesses)
	}

	return router
}
