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
	httpRouter.POST("/users", userController.PostUsers)

	httpRouter.SERVE(port)
}
