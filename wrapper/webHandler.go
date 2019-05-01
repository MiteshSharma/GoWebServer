package wrapper

import (
	"net/http"

	"github.com/MiteshSharma/project/model"
	uuid "github.com/satori/go.uuid"
)

type WebHandler struct {
	HandlerFunc func(*RequestContext, http.ResponseWriter, *http.Request)
	IsLoggedIn  bool
}

func (wh *WebHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//now := time.Now()

	rc := &RequestContext{}
	rc.RequestID = uuid.NewV4().String()
	rc.Path = r.URL.Path

	w.Header().Set(model.HEADER_REQUEST_ID, rc.RequestID)

	if wh.IsLoggedIn {
		rc.Claims, rc.Err = rc.GetClaim(r)
	}

	if rc.Err == nil {
		wh.HandlerFunc(rc, w, r)
	}

	statusCode := http.StatusOK
	if rc.Err != nil {
		statusCode = rc.Err.Status
		rc.Err.RequestId = rc.RequestID
		w.Write([]byte(rc.Err.ToJson()))
	}
	if rc.AppResponse != nil {
		statusCode = rc.AppResponse.Status
	}
	w.WriteHeader(statusCode)
	if rc.AppResponse != nil {
		w.Write([]byte(rc.AppResponse.Response))
	}

	// if rc.App.Metrics != nil {
	// 	elapsedDuration := float64(time.Since(now).Nanoseconds()) / float64(time.Millisecond)
	// 	rc.App.Metrics.RequestReceivedDetail(rc.Path, r.Method, statusCode, elapsedDuration)
	// }
}
