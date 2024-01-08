package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// e.GET("/users/:id", GetUser)
func GetUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, "get user: "+id)
}

// e.POST("/new", NewUser)
func NewUser(c echo.Context) error {
	// Get name and email
	email := c.FormValue("email")
	passwd := c.FormValue("passwd")
	return c.String(http.StatusOK, "new user email:"+email+" passwd:"+passwd)
}

// e.POST("/users/list", GetUserList)
func GetUserList(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "")
}

func UpdateUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusNotImplemented, "update user: "+id)
}
