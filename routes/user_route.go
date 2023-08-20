package routes

import "github.com/gorilla/mux"

type UserService interface {

}

type UserRouter struct {
	r   	*mux.Router
	service  UserService
}

func NewUserRouter(r *mux.Router, db any) *UserRouter {
	return &UserRouter{
		r: r,
		// service: ?
	}
}

func (h *UserRouter) Routes() {
	// define all the user related routes
	// maybe login logout too
}