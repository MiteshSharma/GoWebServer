package api

import (
	"net/http"

	"github.com/MiteshSharma/project/model"
	"github.com/MiteshSharma/project/wrapper"
)

func (a *UserAPI) InitUser() {
	a.Router.User.Handle("", wrapper.RequestHandler(a.createUser)).Methods("POST")
	a.Router.User.Handle("/{userId:[0-9]+}", wrapper.RequestWithAuthHandler(a.updateUser)).Methods("PUT")
	a.Router.User.Handle("/{userId:[0-9]+}", wrapper.RequestWithAuthHandler(a.deleteUser)).Methods("DELETE")
	a.Router.User.Handle("/{userId:[0-9]+}", wrapper.RequestWithAuthHandler(a.getUser)).Methods("GET")
	a.Router.User.Handle("", wrapper.RequestWithAuthHandler(a.getAllUser)).Methods("GET")
}

// CreateHandler func is used to create user
func (a *UserAPI) createUser(rc *wrapper.RequestContext, w http.ResponseWriter, r *http.Request) {
	user := model.UserFromJson(r.Body)
	if user == nil {
		rc.SetError("Body received for user creation is invalid.", http.StatusBadRequest)
		return
	}
	if err := user.Valid(); err != nil {
		rc.SetError("User object received is not valid.", http.StatusBadRequest)
		return
	}

	var err *model.AppError
	if user, err = a.App.CreateUser(user); err != nil {
		rc.SetError("User object creation failed.", http.StatusInternalServerError)
		return
	}
	rc.SetAppResponse(user.ToJson(), http.StatusCreated)
}

// UpdateHandler func is used to create user
func (a *UserAPI) updateUser(rc *wrapper.RequestContext, w http.ResponseWriter, r *http.Request) {
	user := model.UserFromJson(r.Body)
	if user == nil {
		rc.SetError("Body received for user creation is invalid.", http.StatusBadRequest)
		return
	}
	if user.UserID == 0 {
		rc.SetError("UserId received to update userID is 0.", http.StatusBadRequest)
		return
	}
	var err *model.AppError
	if user, err = a.App.UpdateUser(user); err != nil {
		rc.SetError("User object update failed.", http.StatusInternalServerError)
		return
	}
	rc.SetAppResponse(user.ToJson(), http.StatusOK)
}

// GetHandler func is used to get user or users
func (a *UserAPI) getUser(rc *wrapper.RequestContext, w http.ResponseWriter, r *http.Request) {
	userID := rc.GetUserID()
	var user *model.User
	var appErr *model.AppError
	if user, appErr = a.App.GetUser(userID); appErr != nil {
		rc.SetError("User object get failed.", http.StatusInternalServerError)
		return
	}

	rc.SetAppResponse(user.ToJson(), http.StatusOK)
}

// DeleteHandler func is to delete user
func (a *UserAPI) deleteUser(rc *wrapper.RequestContext, w http.ResponseWriter, r *http.Request) {
	userID := rc.GetUserID()
	if _, err := a.App.DeleteUser(userID); err != nil {
		rc.SetError("User object get failed.", http.StatusInternalServerError)
		return
	}
	rc.SetAppResponse("{'response': 'OK'}", http.StatusOK)
}

// GetHandler func is used to get user or users
func (a *UserAPI) getAllUser(rc *wrapper.RequestContext, w http.ResponseWriter, r *http.Request) {
	var users []*model.User
	var err *model.AppError
	if users, err = a.App.GetAllUser(); err != nil {
		rc.SetError("All users object get failed.", http.StatusInternalServerError)
		return
	}

	rc.SetAppResponse(model.UsersToJson(users), http.StatusOK)
}
