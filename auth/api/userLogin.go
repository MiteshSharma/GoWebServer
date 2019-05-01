package api

import (
	"net/http"

	"github.com/MiteshSharma/project/model"
	"github.com/MiteshSharma/project/wrapper"
)

func (a *AuthAPI) InitUserLogin() {
	a.Router.User.Handle("/auth", wrapper.RequestHandler(a.userLogin)).Methods("POST")
	a.Router.User.Handle("/auth", wrapper.RequestHandler(a.userRefresh)).Methods("PUT")
	a.Router.User.Handle("/{userId:[0-9]+}/auth", wrapper.RequestWithAuthHandler(a.userLogout)).Methods("DELETE")
}

func (a *AuthAPI) userLogin(rc *wrapper.RequestContext, w http.ResponseWriter, r *http.Request) {
	user := model.UserFromJson(r.Body)
	if user == nil {
		rc.SetError("Body received for user creation is invalid.", http.StatusBadRequest)
		return
	}
	if err := user.Valid(); err != nil {
		rc.SetError("User object received is not valid.", http.StatusBadRequest)
		return
	}

	userAuth, err := a.App.UserLogin(user)
	if err != nil {
		rc.SetError(err.Message, err.Status)
		return
	}

	rc.SetAppResponse(userAuth.ToJson(), http.StatusCreated)
}

func (a *AuthAPI) userRefresh(rc *wrapper.RequestContext, w http.ResponseWriter, r *http.Request) {
	userAuth := model.UserAuthFromJson(r.Body)
	if userAuth == nil {
		rc.SetError("Body received for user refresh is invalid.", http.StatusBadRequest)
		return
	}

	userAuth, err := a.App.UserRefreshToken(userAuth)
	if err != nil {
		rc.SetError(err.Message, err.Status)
		return
	}

	rc.SetAppResponse(userAuth.ToJson(), http.StatusOK)
}

func (a *AuthAPI) userLogout(rc *wrapper.RequestContext, w http.ResponseWriter, r *http.Request) {
	userID := a.App.UserSession.UserID
	if userID == 0 {
		rc.SetError("UserId received is invalid.", http.StatusBadRequest)
		return
	}
	a.App.UserLogout(userID)

	rc.SetAppResponse("{'response': 'OK'}", http.StatusOK)
}
