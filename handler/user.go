package handler

import (
	"errors"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/go-chi/render"
)

type LoginRequest struct {
	Email  string `json:"email"`
	Passwd string `json:"passwd"`
}

func (a *LoginRequest) Bind(r *http.Request) error {
	if a.Email == "" {
		return errors.New("missing required email address for login")
	}
	if a.Passwd == "" {
		return errors.New("missing required password for login")
	}
	return nil
}

/*
send login request:
curl -X POST -d '{"email":"flo","passwd":"1234"}' localhost:8081/login
*/

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {

	if h.Config.LogLevel == "debug" {
		reqDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%v\n", string(reqDump))
	}

	d := &LoginRequest{}
	if err := render.Bind(r, d); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	w.Write([]byte("login"))
}

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

// import (
// 	"net/http"
// 	"time"

// 	"github.com/floriwan/srcm/pkg/db/model"
// 	"github.com/golang-jwt/jwt/v5"
// 	"github.com/labstack/echo/v4"
// )

// /*

//  send request
//  curl -X POST -d 'email=test@test.de' -d 'password=shhh!' localhost:1323/login

//  response:
//  	{
//   		"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0NjE5NTcxMzZ9.RB3arc4-OyzASAaUhC2W3ReWaXAt_z2Fd3BN4aWTgEY"
// 	}

//  auth request:
//  curl localhost:1323/restricted -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0NjE5NTcxMzZ9.RB3arc4-OyzASAaUhC2W3ReWaXAt_z2Fd3BN4aWTgEY"

//  jwtCustomClaims are custom claims extending default ones.
//  See https://github.com/golang-jwt/jwt for more examples

// */

// type JwtCustomClaims struct {
// 	Email string `json:"email"`
// 	Admin bool   `json:"admin"`
// 	jwt.RegisteredClaims
// }

// func (h *Handler) Register(c echo.Context) error {
// 	return c.String(http.StatusNotImplemented, "")
// }

// func (h *Handler) Login(c echo.Context) error {
// 	email := c.FormValue("email")
// 	password := c.FormValue("password")

// 	// read user from database
// 	u := model.User{Email: email}
// 	res := h.DB.First(&u)

// 	// user not found error
// 	if res.Error != nil {
// 		return echo.ErrNotFound
// 	}

// 	// unauthorized error
// 	if err := u.CheckPassword(password); err != nil {
// 		return echo.ErrUnauthorized
// 	}

// 	// Set custom claims
// 	claims := &JwtCustomClaims{
// 		Email: "Florian",
// 		Admin: true,
// 		RegisteredClaims: jwt.RegisteredClaims{
// 			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
// 		},
// 	}

// 	// Create token with claims
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// 	// Generate encoded token and send it as response.
// 	t, err := token.SignedString([]byte("secret"))
// 	if err != nil {
// 		return err
// 	}

// 	return c.JSON(http.StatusOK, echo.Map{
// 		"token": t,
// 	})
// }

// func (h *Handler) CreateUser(c echo.Context) error {
// 	// Get name and email
// 	email := c.FormValue("email")
// 	passwd := c.FormValue("passwd")
// 	return c.String(http.StatusOK, "new user email:"+email+" passwd:"+passwd)
// }

// func (h *Handler) GetUser(c echo.Context) error {
// 	// User ID from path `users/:id`
// 	id := c.Param("id")
// 	return c.String(http.StatusOK, "get user: "+id)
// }

// func (h *Handler) GetAllUsers(c echo.Context) error {
// 	return c.String(http.StatusNotImplemented, "")
// }

// func (h *Handler) DeleteUser(c echo.Context) error {
// 	return c.String(http.StatusNotImplemented, "")
// }

// func (h *Handler) UpdateUser(c echo.Context) error {
// 	id := c.Param("id")
// 	return c.String(http.StatusNotImplemented, "update user: "+id)
// }
