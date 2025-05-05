package presenter

import (
	"encoding/json"
	"net/http"

	"github.com/mikuto0831/cybozu-http-server/internal/api"
	"github.com/mikuto0831/cybozu-http-server/internal/usecase/dto/output"
)

type IObjectsPresenter interface {
	GetObjectByID(w http.ResponseWriter, output *output.GetObjectByID)
	PutObject(w http.ResponseWriter)
}

type ObjectsPresenter struct {
}

func NewObjectsPresenter() *ObjectsPresenter {
	return &ObjectsPresenter{}
}

func (p *ObjectsPresenter) GetObjectByID(w http.ResponseWriter, output *output.GetObjectByID) {
	response := api.GetObjectResponse{
		Data:     string(output.Object.Data),
	}

	outputJSON, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.Write(outputJSON)
}

func (p *ObjectsPresenter) PutObject(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
