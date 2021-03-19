package app

import "github.com/mattrx/poc-package-requirements/internal/api"

type apiInterface interface {
	Do(req api.Request) (*api.Response, error)
}

type Handler struct {
	api apiInterface
}

func (h Handler) Do() {
	_, _ = h.api.Do(api.Request{})
}
