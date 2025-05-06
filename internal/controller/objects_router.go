package controller

import (
	"net/http"

	"github.com/mikuto0831/cybozu-http-server/internal/pkg/logger"
)

type IObjectsRouter interface {
	HandleObjectsRequest(w http.ResponseWriter, r *http.Request)
}

type objectsRouter struct {
	objectsController IObjectsController
	logger            logger.ILogger
}

func NewObjectsRouter(objectsController IObjectsController) IObjectsRouter {
	return &objectsRouter{
		objectsController: objectsController,
		logger:          logger.NewLogger("objectsRouter"),
	}
}

func (r *objectsRouter) HandleObjectsRequest(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
		case http.MethodGet:
			r.logger.Info("GET request received")
			r.objectsController.GetObjectByID(w, req)
		case http.MethodPut:
			r.logger.Info("PUT request received")
			r.objectsController.PutObject(w, req)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
