package routes

import (
	"github.com/gin-gonic/gin"
	"james.practice/ecommerce/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/admin/addproduct", controllers.ProductViewerAdmin())
	incomingRoutes.GET("/users/productview", controllers.SearchProdcut())
	incomingRoutes.GET("/users/search", controllers.ProductByQuery())
}