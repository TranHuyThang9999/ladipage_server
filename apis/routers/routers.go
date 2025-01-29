package routers

import (
	"github.com/gin-gonic/gin"
	"ladipage_server/apis/controllers"
	"ladipage_server/apis/middlewares"
)

type ApiRouter struct {
	Engine *gin.Engine
}

func NewApiRouter(
	cors *middlewares.MiddlewareCors,
	user *controllers.UserController,
	auth *middlewares.MiddlewareJwt,
	vehicleCategories *controllers.VehicleCategoriesController,
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
			}
		}
	}

	return &ApiRouter{
		Engine: engine,
	}
}
