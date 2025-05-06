package presenter

import "net/http"

type IErrorPresenter interface {
	BadRequest(w http.ResponseWriter)
	MethodNotAllowed(w http.ResponseWriter)
}

type ErrorPresenter struct {
}

func NewErrorPresenter() *ErrorPresenter {
	return &ErrorPresenter{}
}

func (p *ErrorPresenter) BadRequest(w http.ResponseWriter){
	w.WriteHeader(http.StatusNotFound)
}

func (p *ErrorPresenter) MethodNotAllowed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusMethodNotAllowed)
}