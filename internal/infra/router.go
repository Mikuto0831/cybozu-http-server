package infra

import (
	"net/http"

	"github.com/mikuto0831/cybozu-http-server/internal/controller"
)

type Router struct {
	objectsController controller.IObjectsRouter
}

func NewRouter() *Router {
	return &Router{
		objectsController: controller.NewObjectsRouter(controller.NewObjectsController()),
	}
}

func (r *Router) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/objects/", r.objectsController.HandleObjectsRequest)
}