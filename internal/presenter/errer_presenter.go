package presenter

import "net/http"

type IErrerPresenter interface {
	BadRequest(w http.ResponseWriter, err error) error
	MethodNotAllowed(w http.ResponseWriter) error
}

type ErrerPresenter struct {
}

func NewErrerPresenter() *ErrerPresenter {
	return &ErrerPresenter{}
}

func (p *ErrerPresenter) BadRequest(w http.ResponseWriter){
	w.WriteHeader(http.StatusBadRequest)
}

func (p *ErrerPresenter) MethodNotAllowed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusMethodNotAllowed)
}