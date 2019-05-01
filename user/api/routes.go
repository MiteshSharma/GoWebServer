package api

import (
	"github.com/gorilla/mux"
)

type Router struct {
	Root    *mux.Router // ''
	APIRoot *mux.Router // 'api/v1'
	User    *mux.Router // 'api/v1/user'
	Setting *mux.Router // 'api/v1/setting'
	Metrics *mux.Router // 'api/v1/metrics'
}

func (a *UserAPI) setupRoutes() {
	a.Router.Root = a.MainRouter
	a.Router.APIRoot = a.MainRouter.PathPrefix("/api/v1").Subrouter()
	a.Router.User = a.Router.APIRoot.PathPrefix("/user").Subrouter()

	a.InitUser()
	a.InitUserDetail()
}
