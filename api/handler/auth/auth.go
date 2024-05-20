package auth

import (
	"context"
	"golang-websocket/api/database"
	"golang-websocket/api/helper"
	"golang-websocket/api/helper/authentication"
	"golang-websocket/api/models"
	"golang-websocket/api/repository/user"
	"golang-websocket/api/usecase"
	ucase "golang-websocket/api/usecase/auth"
	ucaseUser "golang-websocket/api/usecase/user"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type AuthHandler struct {
	AuthUsecase usecase.AuthUsecase
	UserUsecase usecase.UserUsecase
}

func NewAuthHandler() AuthHandler {
	timeout := time.Duration(viper.GetInt(`context.timeout`)) * time.Second
	db := database.Load()
	repoUser := user.NewUserRepository(db)
	ucaseAutch := ucase.NewAuthUsecase(repoUser, timeout)
	ucaseUsers := ucaseUser.NewUserUsecase(repoUser, timeout)

	return AuthHandler{
		AuthUsecase: ucaseAutch,
		UserUsecase: ucaseUsers,
	}
}

func (u *AuthHandler) Login(c *gin.Context) {
	var res = c.Writer
	var data = make(map[string]interface{})
	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")

	user, msg, err := u.AuthUsecase.Login(ctx, username, password)
	if msg != "" && err != nil {
		helper.ErrorCustomStatus(res, http.StatusBadRequest, msg)
		return
	}
	if err != nil {
		helper.ErrorCustomStatus(res, http.StatusBadGateway, err.Error())
		return
	}

	token, err := authentication.GenerateToken(user)
	if err != nil {
		helper.ErrorCustomStatus(res, http.StatusInternalServerError, err.Error())
		return
	}

	data["user"] = user
	data["token"] = token
	helper.Responses(res, http.StatusOK, msg, data)
}

func (u *AuthHandler) Register(c *gin.Context) {
	var res = c.Writer
	var user models.User
	var ctx = c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	nama := c.Request.FormValue("nama")
	username := c.Request.FormValue("username")
	email := c.Request.FormValue("email")
	password := c.Request.FormValue("password")
	passwordConfirm := c.Request.FormValue("password_confirm")

	if password != passwordConfirm {
		helper.ErrorCustomStatus(res, http.StatusBadRequest, "Password doesn't match")
		return
	}

	if helper.IsEmail(email) == false {
		helper.ErrorCustomStatus(res, http.StatusBadRequest, "Invalid email")
		return
	}

	user.Nama = nama
	user.Username = username
	user.Email = email
	user.Password = password

	result, err := u.UserUsecase.Insert(ctx, user)
	if err != nil {
		helper.ErrorCustomStatus(res, http.StatusInternalServerError, err.Error())
		return
	}

	helper.Responses(res, http.StatusOK, "Success", result)
}
