package routes

import "github.com/gorilla/mux"

type BlobService interface {

}

type BlobRouter struct {
	r   	*mux.Router
	service BlobService
}

func NewBlobRouter(r *mux.Router, db any) *BlobRouter {
	return &BlobRouter{
		r: r,
	}
}

func (h *BlobRouter) Routes() {
	// define all the blob related routes
	// maybe check auth the request and session
}