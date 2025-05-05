package presenter

import (
	"encoding/json"
	"net/http"

	"github.com/mikuto0831/cybozu-http-server/internal/api"
	"github.com/mikuto0831/cybozu-http-server/internal/usecase/dto/output"
)

type IObjectsPresenter interface {
	GetObjectByID(w http.ResponseWriter, output *output.GetObjectByID) error
	PutObject(w http.ResponseWriter) error
}

type ObjectsPresenter struct {
}

func NewObjectsPresenter() *ObjectsPresenter {
	return &ObjectsPresenter{}
}

func (p *ObjectsPresenter) GetObjectByID(w http.ResponseWriter, output *output.GetObjectByID) error {
	response := api.GetObjectResponse{
		ObjectID: output.Object.ID,
		Data:     string(output.Object.Data),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// JSON形式でレスポンスを返す
	if err := json.NewEncoder(w).Encode(response); err != nil {
		return err
	}
	return nil
}

func (p *ObjectsPresenter) PutObject(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
