package wrapper

import (
	"net/http"
)

func RequestHandler(handler func(c *RequestContext, w http.ResponseWriter, r *http.Request)) http.Handler {
	return &WebHandler{
		HandlerFunc: handler,
		IsLoggedIn:  false,
	}
}

func RequestWithAuthHandler(handler func(c *RequestContext, w http.ResponseWriter, r *http.Request)) http.Handler {
	return &WebHandler{
		HandlerFunc: handler,
		IsLoggedIn:  true,
	}
}
