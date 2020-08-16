package router

import (
	routing "github.com/buaazp/fasthttprouter"
	handler "gitlab.com/pplayground/pet_tracking/user-service/handler"
)

func New() *routing.Router {
	return routing.New()
}

func Mount(r *routing.Router) {
	r.GET("/",handler.GetUsers)
	r.GET("/:userId/", handler.GetUser)
	r.POST("/", handler.CreateUser)
	r.PUT("/:userId", handler.UpdateUser)
	r.DELETE("/:userId", handler.DeleteUser)
}
