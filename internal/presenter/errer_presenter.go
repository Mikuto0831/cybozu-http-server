package presenter

import "net/http"

type IErrerPresenter interface {
	BadRequest(w http.ResponseWriter, err error)
	MethodNotAllowed(w http.ResponseWriter)
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