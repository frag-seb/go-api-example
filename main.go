package main

import (
	"demo/controller"
	router "demo/http"
	"fmt"
	"net/http"
)

var (
	httpRouter     router.Router             = router.NewMuxRouter() // router.NewChiRouter()
	userController controller.UserController = controller.NewUserController()
)

func main() {
	const port string = ":8000"

	httpRouter.REQUEST("/", func(response http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprintln(response, "Up and running ...")
	})

	httpRouter.GET("/users", userController.GetUsers)
	httpRouter.POST("/users", userController.PostUsers)

	httpRouter.SERVE(port)

}
