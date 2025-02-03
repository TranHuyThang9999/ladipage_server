package routers

import (
	"ladipage_server/apis/controllers"
	"ladipage_server/apis/middlewares"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

type ApiRouter struct {
	Engine *gin.Engine
}

func NewApiRouter(
	//cors *middlewares.MiddlewareCors,
	user *controllers.UserController,
	auth *middlewares.MiddlewareJwt,
	vehicleCategories *controllers.VehicleCategoriesController,
	fileDescController *controllers.FileDescController,
	vehicle *controllers.VehicleController,
) *ApiRouter {
	engine := gin.New()
	//engine.Use(cors.CorsAPI())
	engine.Use(cors.AllowAll())
	gin.DisableConsoleColor()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

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
			vehicleCategoriesGroup.GET("/public/list", vehicleCategories.ListVehicle)
			authorizedVehicleCategories := vehicleCategoriesGroup.Use(auth.Authorization())
			{
				authorizedVehicleCategories.POST("/add", vehicleCategories.AddVehicle)
				authorizedVehicleCategories.GET("/list", vehicleCategories.ListVehicle)
				authorizedVehicleCategories.PATCH("/update", vehicleCategories.UpdateVehicleCategoryByID)
				authorizedVehicleCategories.DELETE("/delete/:id", vehicleCategories.DeleteVehicleCategoryByID)
				authorizedVehicleCategories.POST("/file_desc/add_list", vehicleCategories.AddListFileByObjectID)
				authorizedVehicleCategories.DELETE("/file_desc/delete", vehicleCategories.DeleteListFileByID)
				authorizedVehicleCategories.GET("/file_desc/:objectID", vehicleCategories.ListFileByObjectID)
			}
		}
		vehicleGroup := r.Group("/vehicle")
		{
			vehicleGroup.GET("/public/list", vehicle.GetVehicles)
			vehicleGroup.GET("/public/file_desc/:objectID", vehicle.GetListFileVehicleById)
			vehicleGroup.GET("/public/list/vehicle/:vehicleCategoryID", vehicle.GetVehiclesByVehicleCategoryIDForPublic)
			authorizedVehicle := vehicleGroup.Use(auth.Authorization())
			{
				authorizedVehicle.POST("/add", vehicle.AddVehicle)
				authorizedVehicle.GET("/list", vehicle.GetVehicles)
				authorizedVehicle.GET("/file_desc/:objectID", vehicle.GetListFileVehicleById)
				authorizedVehicle.PATCH("/update", vehicle.UpdateVehicleById)
				authorizedVehicle.DELETE("/delete/:id", vehicle.DeleteVehicleById)
				authorizedVehicle.POST("/file_desc/add_list", vehicle.AddListFileByObjectID)
				authorizedVehicle.DELETE("/file_desc/delete", vehicle.DeleteListFileByID)

			}
		}

	}

	return &ApiRouter{
		Engine: engine,
	}
}
