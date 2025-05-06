package controller

import (
	"net/http"
)

type IObjectsRouter interface {
	HandleObjectsRequest(w http.ResponseWriter, r *http.Request)
}

type objectsRouter struct {
	objectsController IObjectsController
}

func NewObjectsRouter(objectsController IObjectsController) IObjectsRouter {
	return &objectsRouter{
		objectsController: objectsController,
	}
}

func (r *objectsRouter) HandleObjectsRequest(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
		case http.MethodGet:
			r.objectsController.GetObjectByID(w, req)
		case http.MethodPut:
			r.objectsController.PutObject(w, req)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
