package main

import (
	"Demo/auth"
	"Demo/controller"
	"Demo/repository"
	"Demo/router"
	posts2 "Demo/service/posts"
	users2 "Demo/service/users"
	"fmt"
	"net/http"
)

var (
	userRepo       repository.Repository     = repository.NewRepository()
	postRepo       repository.PostRepository = repository.NewPostRepository()
	postService    posts2.PostService        = posts2.NewPostService(postRepo)
	userService    users2.Service            = users2.NewService(userRepo)
	postController controller.PostController = controller.NewPostService(postService)
	userController controller.UserController = controller.NewUserController(userService)
	httpRouter     router.Router             = router.NewMuxRouter()
)

func main() {
	const port string = ":8080"

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Server is running")
	})
	httpRouter.POST("/users/register", userController.Register)
	httpRouter.POST("/users/login", userController.Login)
	httpRouter.POST("/users/posts", auth.IsLoggedIn(postController.AddPost))
	httpRouter.GET("/users/posts", auth.IsAdmin(postController.GetPosts))

	httpRouter.SERVER(port)
}
