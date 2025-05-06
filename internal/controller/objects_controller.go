package controller

import (
	"net/http"
	"path"

	"github.com/mikuto0831/cybozu-http-server/internal/presenter"
	"github.com/mikuto0831/cybozu-http-server/internal/pkg/unit"
	"github.com/mikuto0831/cybozu-http-server/internal/pkg/logger"
	"github.com/mikuto0831/cybozu-http-server/internal/usecase"
	"github.com/mikuto0831/cybozu-http-server/internal/usecase/dto/input"
	"github.com/mikuto0831/cybozu-http-server/internal/usecase/dto/output"
)

type IObjectsController interface {
	GetObjectByID(w http.ResponseWriter, r *http.Request)
	PutObject(w http.ResponseWriter, r *http.Request)
}

type ObjectsController struct {
	objectsUsecase usecase.IObjectsUsecase
	objectsPresenter presenter.IObjectsPresenter
	errorPresenter presenter.IErrorPresenter
	logger logger.ILogger
}

func NewObjectsController() *ObjectsController {
	return &ObjectsController{
		objectsUsecase:  usecase.NewObjectsUsecase(),
		objectsPresenter: presenter.NewObjectsPresenter(),
		errorPresenter:   presenter.NewErrorPresenter(),
		logger:          logger.NewLogger("objectsController"),
	}
}

func (c *ObjectsController) GetObjectByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		// GETメソッド以外は405 Method Not Allowedを返す
		c.errorPresenter.MethodNotAllowed(w)
		return
	}

	objectId, err := unit.IsValidID(path.Base(r.URL.Path))

	if err != nil {
		c.errorPresenter.NotFound(w)
		return
	}

	input := input.GetObjectByID{
		ID: objectId,
	}

	object, err := c.objectsUsecase.GetObjectByID(input.ID)
	if err != nil {
		c.errorPresenter.NotFound(w)
		return
	}

	outputObject := &output.GetObjectByID{
		Object: object,
	}

	c.objectsPresenter.GetObjectByID(w, outputObject)
}

func (c *ObjectsController) PutObject(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		// PUTメソッド以外は405 Method Not Allowedを返す
		c.errorPresenter.MethodNotAllowed(w)
		return
	}
	
	objectId, err := unit.IsValidID(path.Base(r.URL.Path))

	if err != nil {
		c.errorPresenter.NotFound(w)
		return
	}

	data := make([]byte, r.ContentLength)
	r.Body.Read(data)

	input := input.PutObject{
		ID:   objectId,
		Data: data,
	}

	err = c.objectsUsecase.PutObject(input.ID, input.Data)
	if err != nil {
		c.errorPresenter.NotFound(w)
		return
	}

	c.objectsPresenter.PutObject(w)
}
