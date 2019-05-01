package api

import (
	"net/http"

	"github.com/MiteshSharma/project/model"
	"github.com/MiteshSharma/project/wrapper"
)

func (a *UserAPI) InitUserDetail() {
	a.Router.User.Handle("/{userId:[0-9]+}/userDetail", wrapper.RequestWithAuthHandler(a.updateUserDetail)).Methods("PUT")
}

func (a *UserAPI) updateUserDetail(rc *wrapper.RequestContext, w http.ResponseWriter, r *http.Request) {
	userDetail := model.UserDetailFromJson(r.Body)
	if userDetail == nil {
		rc.SetError("Body received for user detail is invalid.", http.StatusBadRequest)
		return
	}
	if userDetail.UserID == 0 {
		rc.SetError("UserId received to update user detail is 0.", http.StatusBadRequest)
		return
	}
	var err *model.AppError
	if userDetail, err = a.App.UpdateUserDetail(userDetail); err != nil {
		rc.SetError("User object update failed.", http.StatusInternalServerError)
		return
	}
	rc.SetAppResponse(userDetail.ToJson(), http.StatusOK)
}
