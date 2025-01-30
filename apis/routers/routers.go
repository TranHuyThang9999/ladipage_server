package routers

import (
	"ladipage_server/apis/controllers"
	"ladipage_server/apis/middlewares"

	"github.com/gin-gonic/gin"
)

type ApiRouter struct {
	Engine *gin.Engine
}

func NewApiRouter(
	cors *middlewares.MiddlewareCors,
	user *controllers.UserController,
	auth *middlewares.MiddlewareJwt,
	vehicleCategories *controllers.VehicleCategoriesController,
	fileDescController *controllers.FileDescController,
) *ApiRouter {
	engine := gin.New()

	gin.DisableConsoleColor()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	engine.Use(cors.CorsAPI())

	r := engine.RouterGroup.Group("/manager")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	//admin
	userGroup := r.Group("/user")
	{
		userGroup.POST("/register", user.Register)
		userGroup.POST("/login", user.Login)
		userGroup.POST("/register/auth2", user.LoginWithGG)
		authorized := userGroup.Group("/")
		authorized.Use(auth.Authorization())
		{
			authorized.GET("/profile", user.Profile)
		}

		vehicleCategoriesGroup := r.Group("/vehicle_categories")
		{
			vehicleCategoriesGroup.Use(auth.Authorization())
			{
				vehicleCategoriesGroup.POST("/add", vehicleCategories.AddVehicle)
				vehicleCategoriesGroup.GET("/list", vehicleCategories.ListVehicle)
				vehicleCategoriesGroup.PATCH("/update", vehicleCategories.UpdateVehicleCategoryByID)
				vehicleCategoriesGroup.DELETE("/delete/:id", vehicleCategories.DeleteVehicleCategoryByID)
				vehicleCategoriesGroup.POST("/file_desc/add_list", vehicleCategories.AddListFileByObjectID)
				vehicleCategoriesGroup.DELETE("/file_desc/delete", vehicleCategories.DeleteListFileByID)
				vehicleCategoriesGroup.GET("/file_desc/:objectID", vehicleCategories.ListFileByObjectID)
			}
		}
	}

	return &ApiRouter{
		Engine: engine,
	}
}
