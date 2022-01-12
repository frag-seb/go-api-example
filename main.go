package main

import (
	"demo/connection"
	"demo/controller"
	router "demo/http"
)

var (
	httpRouter     router.Router             = router.NewMuxRouter() // router.NewChiRouter()
	userController controller.UserController = controller.NewUserController()
)

func main() {
	const port string = ":8000"
	connection.ConnectDatabase() // new

	httpRouter.GET("/users", userController.GetUsers)
	httpRouter.GET("/users/{id}", userController.GetUser)
	httpRouter.POST("/users", userController.PostUsers)
	httpRouter.DELETE("/users/{id}", userController.DeleteUser)

	httpRouter.SERVE(port)
}
