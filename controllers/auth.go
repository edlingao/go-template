package controllers

import (
	"exercise-app/models"
	"exercise-app/utils"
	view_auth "exercise-app/view/auth"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct{}
type RegisterBody struct {
	Username string
	Email    string
	Password string
}

type SigninStatus struct {
  Success bool
  Message string
}

func (a AuthController) RegisterView(c echo.Context) error {
	return utils.Render(
		c,
		view_auth.Register(),
		http.StatusOK,
	)
}

func (a AuthController) LoginView(c echo.Context) error {
	return utils.Render(
		c,
		view_auth.SignIn(),
		http.StatusOK,
	)
}

func (a AuthController) RegisterUser(c echo.Context) error {
	email := c.FormValue("email")
	username := c.FormValue("username")
	password := c.FormValue("password")

	user, err := models.USER.Register(email, username, password)

	if err != nil {
		log.Print("ERROR", err)
		return c.HTML(
			http.StatusBadGateway,
			fmt.Sprintf("<h1>AQUi: %v </h1>", err.Error()),
		)
	}

	signed_api, err := models.GenerateAPIKey(user.Username)

	if err != nil {
		log.Fatal("ERROR AQUI", err)
		return c.HTML(
			http.StatusBadGateway,
			fmt.Sprintf("<h1>ERROR: %v </h1>", err.Error()),
		)
	}

	return setCookieAndRedirect(c, signed_api)
}

func (a AuthController) SignIn(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	user, err := models.USER.Login(email, password)

	if err != nil {
		return c.JSON(
      http.StatusBadGateway,
      SigninStatus{Success: false, Message: "Error, user not found"},
    )
	}

	signed_api, err := models.GenerateAPIKey(user.Username)

	if err != nil {
		return c.JSON(
      http.StatusBadGateway,
      SigninStatus{Success: false, Message: "Error generating API key"},
    )
  }

	return setCookieAndRedirect(c, signed_api)
}

func setCookieAndRedirect(c echo.Context, signed_api string) error {
	cookie := models.GetCookie(signed_api)
	c.SetCookie(cookie)
	c.Response().Header().Set("HX-Location", "/")
  return c.JSON(http.StatusOK, SigninStatus{Success: true, Message: "Success"})
}
