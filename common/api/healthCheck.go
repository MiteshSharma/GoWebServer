package api

import (
	"io"
	"net/http"

	"github.com/MiteshSharma/project/wrapper"
)

func (a *CommonAPI) InitHealthCheck() {
	a.Router.APIRoot.Handle("/healthCheck", wrapper.RequestHandler(a.HealthCheck))
}

// HealthCheck func is used to check health check status
func (a *CommonAPI) HealthCheck(c *wrapper.RequestContext, w http.ResponseWriter, r *http.Request) {
	// Check for status of DB or cache (Redis) in future
	io.WriteString(w, `{"alive": true}`)
}
