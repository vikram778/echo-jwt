package auth

import (
	"context"
	"echo-jwt/app/errs"
	"echo-jwt/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

//Exception struct
type Exception errs.ErrorResponse

// JwtVerify Middleware function
func JwtVerify(next echo.HandlerFunc) echo.HandlerFunc {
	return (func(c echo.Context) (err error) {

		var header = c.Request().Header.Get("x-access-token") //Grab the token from the header

		header = strings.TrimSpace(header)

		if header == "" {
			//Token is missing, returns with error code 403 Unauthorized
			c.Response().WriteHeader(http.StatusForbidden)
			return c.JSON(http.StatusForbidden, Exception{Error: "Missing auth token"})
		}
		tk := &model.Token{}

		_, err = jwt.ParseWithClaims(header, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if err != nil {
			c.Response().WriteHeader(http.StatusForbidden)
			return c.JSON(http.StatusForbidden, Exception{Error: err.Error()})
		}

		ctx := context.WithValue(c.Request().Context(), "user", tk)
		c.Request().WithContext(ctx)
		next(c)
		return
	})
}
