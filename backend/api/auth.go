package api

import (
	"errors"
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/VirrageS/chirp/backend/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func (api *API) RegisterUser(context *gin.Context) {
	var newUserForm model.NewUserForm
	if err := context.BindJSON(&newUserForm); err != nil {
		context.AbortWithError(
			http.StatusBadRequest,
			errors.New("Fields: name, username, password and email are required."),
		)
		return
	}

	newUser, err := api.service.RegisterUser(&newUserForm)
	if err != nil {
		statusCode := getStatusCodeFromError(err)
		context.AbortWithError(statusCode, err)
		return
	}

	context.Header("Location", fmt.Sprintf("/user/%d", newUser.ID))
	context.IndentedJSON(http.StatusCreated, newUser)
}

func (api *API) LoginUser(context *gin.Context) {
	var loginForm model.LoginForm

	contentType := context.ContentType()
	switch contentType {
	case "application/json":
		if err := context.BindJSON(&loginForm); err != nil {
			context.AbortWithError(
				http.StatusBadRequest,
				errors.New("Fields: email and password are required."),
			)
			return
		}

	case "application/x-www-form-urlencoded":
		loginForm.Email = context.PostForm("email")
		loginForm.Password = context.PostForm("password")
		if loginForm.Email == "" {
			context.AbortWithError(
				http.StatusBadRequest,
				errors.New("Fields: email and password are erro."),
			)
			return
		}
	case "multipart/form-data":

	}
	log.WithFields(log.Fields{
		"followeeID": loginForm,
		"followerID": "followerID",
	}).Info("FollowUser query error.")
	loggedUser, err := api.service.LoginUser(&loginForm)
	if err != nil {
		statusCode := getStatusCodeFromError(err)
		context.AbortWithError(statusCode, err)
		return
	}

	authToken, refreshToken, err := api.createTokens(loggedUser.ID, context.Request)
	if err != nil {
		statusCode := getStatusCodeFromError(err)
		context.AbortWithError(statusCode, err)
		return
	}

	loginResponse := &model.LoginResponse{
		AuthToken:    authToken,
		RefreshToken: refreshToken,
		User:         loggedUser,
	}

	context.IndentedJSON(http.StatusOK, loginResponse)
}

func (api *API) RefreshAuthToken(context *gin.Context) {
	var requestData model.RefreshAuthTokenRequest
	if err := context.BindJSON(&requestData); err != nil {
		context.AbortWithError(
			http.StatusBadRequest,
			errors.New("Fields: `user_id` and `refresh_token` are required."),
		)
		return
	}

	response, err := api.refreshAuthToken(&requestData, context.Request)
	if err != nil {
		statusCode := getStatusCodeFromError(err)
		context.AbortWithError(statusCode, err)
		return
	}

	context.IndentedJSON(http.StatusOK, response)
}

func (api *API) GetGoogleAuthorizationURL(context *gin.Context) {
	token := "TODO" // TODO: this should be generated hash from IP Address and browser name / browser_id
	context.IndentedJSON(http.StatusOK, api.googleOAuth2.AuthCodeURL(token, oauth2.AccessTypeOffline))
}

func (api *API) CreateOrLoginUserWithGoogle(context *gin.Context) {
	var form model.GoogleLoginForm
	if err := context.BindJSON(&form); err != nil {
		context.AbortWithError(
			http.StatusBadRequest,
			errors.New("Fields: `code` and `state` are required."),
		)
		return
	}

	if form.State != "TODO" {
		context.AbortWithError(http.StatusUnauthorized, errors.New("Invalid Google login form."))
		return
	}

	user, err := api.getGoogleUser(form.Code)
	if err != nil {
		context.AbortWithError(http.StatusBadRequest, errors.New("Error fetching user from Google."))
	}

	loggedUser, err := api.service.CreateOrLoginUserWithGoogle(user)
	if err != nil {
		statusCode := getStatusCodeFromError(err)
		context.AbortWithError(statusCode, err)
		return
	}

	authToken, refreshToken, err := api.createTokens(loggedUser.ID, context.Request)
	if err != nil {
		statusCode := getStatusCodeFromError(err)
		context.AbortWithError(statusCode, err)
		return
	}

	loginResponse := &model.LoginResponse{
		AuthToken:    authToken,
		RefreshToken: refreshToken,
		User:         loggedUser,
	}

	context.IndentedJSON(http.StatusOK, loginResponse)
}
