package api

import (
	"encoding/json"
	"net/http"

	"github.com/MiteshSharma/project/wrapper"
)

func (a *BiAPI) InitBi() {
	a.Router.BI.Handle("", wrapper.RequestHandler(a.handleBiEvent)).Methods("POST")
	a.Router.BI.Handle("", wrapper.RequestHandler(a.handleBiEventBatch)).Methods("PATCH")
}

func (a *BiAPI) handleBiEvent(rc *wrapper.RequestContext, w http.ResponseWriter, r *http.Request) {
	var event map[string]interface{}
	json.NewDecoder(r.Body).Decode(&event)
	a.App.HandleBIEvent(event)
}

func (a *BiAPI) handleBiEventBatch(rc *wrapper.RequestContext, w http.ResponseWriter, r *http.Request) {
	var events []map[string]interface{}
	json.NewDecoder(r.Body).Decode(&events)
	a.App.HandleBIBatch(events)
}
