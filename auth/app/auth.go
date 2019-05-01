package app

import (
	"net/http"
	"time"

	"github.com/MiteshSharma/project/core/logger"
	"github.com/MiteshSharma/project/model"
	"github.com/MiteshSharma/project/util"
	jwt "github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

func (a *App) SignToken(userID int, roles []model.Role) (string, *model.AppError) {
	currentTime := time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":      userID,
		"currentTime": currentTime,
		"sub":         userID,
		"exp":         (currentTime + 86400),
		"roles":       roles,
	})
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(a.Config.AuthConfig.HmacSecret))

	if err != nil {
		a.Log.Debug("Token signing got error", logger.Error(err))
		return "", model.NewAppError("Token signing got error", http.StatusInternalServerError)
	}

	return tokenString, nil
}

// UserRefreshToken function
func (a *App) UserRefreshToken(userAuth *model.UserAuth) (*model.UserAuth, *model.AppError) {
	result := a.Repository.Auth().GetSession(userAuth.UserID)
	var existingUserSession *model.UserSession
	if result.Err == nil {
		existingUserSession = result.Data.(*model.UserSession)
	}
	if existingUserSession.Token == userAuth.RefreshToken {
		// dbUser, err := a.External.GetUser(userAuth.UserID)
		// if err != nil {
		// 	return nil, model.NewAppError("User not found", http.StatusNotFound)
		// }
		token, errToken := a.SignToken(userAuth.UserID, nil)
		if errToken != nil {
			return nil, errToken
		}
		a.Repository.Auth().UpdateSession(existingUserSession)
		userAuth := &model.UserAuth{
			UserID:       userAuth.UserID,
			AuthToken:    token,
			RefreshToken: userAuth.RefreshToken,
		}
		return userAuth, nil
	}
	return nil, model.NewAppError("User token expired", http.StatusNotFound)
}

// UserLogin function
func (a *App) UserLogin(user *model.User) (*model.UserAuth, *model.AppError) {
	dbUser, err := a.External.GetUser(user.UserID)
	if err != nil {
		return nil, model.NewAppError("User not found", http.StatusNotFound)
	}
	if util.CheckPasswordHash(user.Password, dbUser.Salt, dbUser.Password) {
		token, err := a.SignToken(dbUser.UserID, nil)
		if err != nil {
			return nil, err
		}
		// Create refresh token and save it
		refreshToken := uuid.NewV4().String()
		result := a.Repository.Auth().GetSession(dbUser.UserID)
		var existingUserSession *model.UserSession
		if result.Err == nil {
			existingUserSession = result.Data.(*model.UserSession)
		}
		session := &model.UserSession{
			UserID: dbUser.UserID,
			Token:  refreshToken,
		}
		if existingUserSession == nil {
			a.Repository.Auth().CreateSession(session)
		} else {
			a.Repository.Auth().UpdateSession(session)
		}

		userAuth := &model.UserAuth{
			UserID:       dbUser.UserID,
			AuthToken:    token,
			RefreshToken: refreshToken,
		}
		return userAuth, nil
	}
	return nil, model.NewAppError("User credentials incorrect", http.StatusNotFound)
}

// UserLogout function
func (a *App) UserLogout(userID int) {
	a.Repository.Auth().DeleteSession(userID)
}
