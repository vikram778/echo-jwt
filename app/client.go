package app

import (
	"echo-jwt/app/errs"
	"echo-jwt/app/resource/api/client"
	"echo-jwt/model"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func (a *App) Login(c echo.Context) (err error) {
	var (
		user, _ = model.NewClient(a.DB)
		req     = client.LoginRequest{}
		//err     error
	)

	defer func() {
		a.Defer(c.Response().Writer)
	}()

	err = a.GetParams(&req, c.Response().Writer, c.Request())
	if err != nil {
		a.FormatException(c.Request(), err)
		return
	}

	user.Email.SetValid(req.Email)
	user.Password.SetValid(req.Password)

	resp, err := findOne(req.Password, user)
	if err != nil {
		a.FormatException(c.Request(), err)
		return
	}

	a.Status = http.StatusOK
	a.RawBody = resp

	return
}

func findOne(password string, user *model.Client) (map[string]interface{}, error) {

	var (
		err error
	)

	err = user.GetClientByEmail()
	if err != nil {
		return nil, err
	}

	if !user.ID.Valid {
		err = errors.New(errs.ErrUserNotExist)
		return nil, err
	}

	expiresAt := time.Now().Add(time.Minute * 100000).Unix()

	err = bcrypt.CompareHashAndPassword([]byte(user.Password.String), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		err = errors.New(errs.ErrInvalidCreds)
		return nil, err
	}

	tk := &model.Token{
		UserID:   uint(user.ID.Int64),
		Email:    user.Email.String,
		UserName: user.UserName.String,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	var resp = map[string]interface{}{"status": false, "message": "logged in"}
	resp["token"] = tokenString //Store the token in the response
	resp["user"] = user
	return resp, nil
}

//CreateUser function -- create a new user
func (a *App) CreateUser(c echo.Context) (err error) {

	var (
		user, _ = model.NewClient(a.DB)
		req     = client.RegisterRequest{}
		res     = client.RegisterClientResponse{}
	)

	defer func() {
		a.Defer(c.Response().Writer)
	}()

	err = a.GetParams(&req, c.Response().Writer, c.Request())
	if err != nil {
		a.FormatException(c.Request(), err)
		return
	}

	user.UserName.SetValid(req.UserName)
	user.Email.SetValid(req.Email)
	err = user.GetClient()
	if err != nil {
		a.FormatException(c.Request(), err)
		return
	}

	if user.ID.Valid {
		a.FormatException(c.Request(), errors.New(errs.ErrUserExist))
		return
	}
	pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		a.FormatException(c.Request(), err)
		return
	}

	user.Password.SetValid(string(pass))

	err = user.InsertOrUpdate(true)
	if err != nil {
		a.FormatException(c.Request(), err)
		return
	}

	res.Email = req.Email
	res.UserName = req.UserName
	res.Message = "User created successfully"

	a.Status = http.StatusOK
	a.RawBody = res

	return err
}

//FetchUser function
func (a *App) FetchUser(c echo.Context) (err error) {
	var (
		user, _ = model.NewClient(a.DB)
	)

	usr := c.Param("username")
	user.UserName.SetValid(usr)
	err = user.GetClient()
	if err != nil {
		a.FormatException(c.Request(), err)
		return
	}

	a.Status = http.StatusOK
	a.RawBody = user

	return
}

func (a *App) UpdateUser(c echo.Context) (err error) {
	var (
		user, _ = model.NewClient(a.DB)
		req     = client.RegisterRequest{}
	)

	usr := c.Param("username")

	err = a.GetParams(&req, c.Response().Writer, c.Request())
	if err != nil {
		a.FormatException(c.Request(), err)
		return
	}

	user.UserName.SetValid(usr)
	err = user.GetClient()
	if err != nil {
		a.FormatException(c.Request(), err)
		return
	}

	if !user.ID.Valid {
		a.FormatException(c.Request(), errors.New(errs.ErrUserNotExist))
		return
	}

	user.UserName.SetValid(req.UserName)
	user.Email.SetValid(req.Email)
	err = user.InsertOrUpdate(false)
	if err != nil {
		a.FormatException(c.Request(), err)
		return
	}

	a.Status = http.StatusOK
	a.RawBody = user

	return
}

func (a *App) DeleteUser(c echo.Context) (err error) {
	var (
		user, _ = model.NewClient(a.DB)
		req     = client.RegisterRequest{}
	)

	usr := c.Param("username")

	err = a.GetParams(&req, c.Response().Writer, c.Request())
	if err != nil {
		a.FormatException(c.Request(), err)
		return
	}

	user.UserName.SetValid(usr)
	err = user.GetClient()
	if err != nil {
		a.FormatException(c.Request(), err)
		return
	}

	if !user.ID.Valid {
		a.FormatException(c.Request(), errors.New(errs.ErrUserNotExist))
		return
	}

	err = user.DeleteClient()
	if err != nil {
		a.FormatException(c.Request(), err)
		return
	}

	a.Status = http.StatusOK
	a.RawBody = user

	return
}
