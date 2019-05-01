package app

import (
	"fmt"
	"strconv"

	"github.com/MiteshSharma/project/model"
	"github.com/MiteshSharma/project/util"
)

func (a *App) CreateUser(user *model.User) (*model.User, *model.AppError) {
	user.Salt = util.RandStringBytes(6)
	hashedPassword, err := util.HashPassword(user.Password, user.Salt)
	if err != nil {
		a.Log.Error(fmt.Sprintf("Hashing password returned error for userId %d", user.UserID))
	}
	user.Password = hashedPassword

	storageResult := a.Repository.User().CreateUser(user)
	if storageResult.Err != nil {
		return nil, storageResult.Err
	}

	user = storageResult.Data.(*model.User)

	userDetail := &model.UserDetail{
		UserID: user.UserID,
	}
	storageResult = a.Repository.User().CreateUserDetail(userDetail)

	if storageResult.Err != nil {
		return nil, storageResult.Err
	}

	// roles := []model.Role{model.ADMIN}
	userRole := &model.UserRole{
		UserID: user.UserID,
		Role:   model.ADMIN,
	}

	a.Repository.User().AttachRole(userRole)

	notificationMessage := model.NewNotificationMessage("Welcome to backend", "This is your welcome email", "plain")
	notificationData := model.NewNotificationData(user.Email, notificationMessage, model.SLACK)
	a.Bus.Publish(model.SEND_NOTIFICATION, notificationData)

	eventData := map[string]interface{}{
		"userId": strconv.Itoa(user.UserID),
		"email":  user.Email,
	}
	a.BiEventHandler.Send("user_create", eventData)

	return user, nil
}

func (a *App) UpdateUser(user *model.User) (*model.User, *model.AppError) {
	storageResult := a.Repository.User().GetUser(user.UserID)
	if storageResult.Err != nil {
		return nil, storageResult.Err
	}
	existingUser := storageResult.Data.(*model.User)
	if user.FirstName != "" {
		existingUser.FirstName = user.FirstName
	}
	if user.LastName != "" {
		existingUser.LastName = user.LastName
	}

	storageResult = a.Repository.User().UpdateUser(existingUser)
	if storageResult.Err != nil {
		return nil, storageResult.Err
	}

	user = storageResult.Data.(*model.User)
	return user, nil
}

func (a *App) GetUser(userID int) (*model.User, *model.AppError) {
	storageResult := a.Repository.User().GetUser(userID)
	if storageResult.Err != nil {
		return nil, storageResult.Err
	}

	user := storageResult.Data.(*model.User)
	return user, nil
}

func (a *App) GetAllUser() ([]*model.User, *model.AppError) {
	storageResult := a.Repository.User().GetAllUsers()
	if storageResult.Err != nil {
		return nil, storageResult.Err
	}

	users := storageResult.Data.([]*model.User)
	return users, nil
}

func (a *App) DeleteUser(userID int) (*model.User, *model.AppError) {
	storageResult := a.Repository.User().DeleteUser(userID)
	if storageResult.Err != nil {
		return nil, storageResult.Err
	}

	return nil, nil
}

func (a *App) UpdateUserDetail(userDetail *model.UserDetail) (*model.UserDetail, *model.AppError) {
	storageResult := a.Repository.User().GetUserDetail(userDetail.UserID)
	if storageResult.Err != nil {
		return nil, storageResult.Err
	}
	dbUserDetail := storageResult.Data.(*model.UserDetail)
	userDetail.UserDetailID = dbUserDetail.UserDetailID
	storageResult = a.Repository.User().UpdateUserDetail(userDetail)
	if storageResult.Err != nil {
		return nil, storageResult.Err
	}

	userDetail = storageResult.Data.(*model.UserDetail)
	return userDetail, nil
}
