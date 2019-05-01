package wrapper

import (
	"fmt"
	"net/http"
	"time"

	"github.com/MiteshSharma/project/model"
	jwt "github.com/dgrijalva/jwt-go"
)

type RequestContext struct {
	RequestID   string
	Path        string
	AppResponse *model.AppResponse
	Err         *model.AppError
	Claims      map[string]interface{}
}

func (rc *RequestContext) SetError(message string, statusCode int) {
	rc.Err = model.NewAppError(message, statusCode)
}

func (rc *RequestContext) SetAppResponse(response string, statusCode int) {
	rc.AppResponse = model.NewAppResponse(response, statusCode)
}

func (rc *RequestContext) GetClaim(r *http.Request) (map[string]interface{}, *model.AppError) {
	token, err := rc.GetToken(r)
	if err == nil {
		return rc.verifyAndParseToken(token)
	}
	return nil, err
}

func (rc *RequestContext) GetToken(r *http.Request) (string, *model.AppError) {
	return r.Header.Get(model.AUTHENTICATION), nil
}

func (rc *RequestContext) verifyAndParseToken(tokenString string) (map[string]interface{}, *model.AppError) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})
	if err != nil {
		return nil, model.NewAppError(err.Error(), http.StatusUnauthorized)
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//userID := int(claims["userID"].(float64))
		exp := claims["exp"].(int64)
		currentTime := time.Now().Unix()
		if currentTime > exp {
			return claims, nil
		}
		return nil, model.NewAppError("user token expired", http.StatusUnauthorized)
	}
	return nil, model.NewAppError("incorrect token", http.StatusBadRequest)
}

func (rc *RequestContext) GetUserID() int {
	return rc.Claims["userId"].(int)
}
