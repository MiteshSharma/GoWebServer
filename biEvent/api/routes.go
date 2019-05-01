package api

import (
	"github.com/gorilla/mux"
)

type Router struct {
	Root    *mux.Router // ''
	APIRoot *mux.Router // 'api/v1'
	BI      *mux.Router // 'api/v1/bi'
}

func (a *BiAPI) setupRoutes() {
	a.Router.Root = a.MainRouter
	a.Router.APIRoot = a.MainRouter.PathPrefix("/api/v1").Subrouter()
	a.Router.BI = a.Router.APIRoot.PathPrefix("/bi").Subrouter()

	a.InitBi()
}
