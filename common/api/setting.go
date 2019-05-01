package api

import (
	"net/http"

	"github.com/MiteshSharma/project/wrapper"
)

func (a *CommonAPI) InitSetting() {
	a.Router.Setting.Handle("", wrapper.RequestHandler(a.getSetting)).Methods("GET")
}

func (a *CommonAPI) getSetting(rc *wrapper.RequestContext, w http.ResponseWriter, r *http.Request) {
	// rc.SetAppResponse(a.App.Config.ToJson(), http.StatusOK)
}
