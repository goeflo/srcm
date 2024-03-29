package handler

// import (
// 	"errors"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"net/http/httputil"
// 	"time"

// 	"github.com/floriwan/srcm/pkg/db/model"
// 	"github.com/floriwan/srcm/pkg/jwtauth"
// 	"github.com/go-chi/chi"
// 	"github.com/lestrrat-go/jwx/v2/jwt"
// )

// type NewUserRequestData struct {
// 	Email     string `json:"email"`
// 	Passwd    string `json:"passwd"`
// 	Lastname  string `json:"lastname"`
// 	Firstname string `json:"firstname"`
// }

// func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
// 	log.Printf("create user handler\n")

// 	if h.Config.LogLevel == "debug" {
// 		reqDump, err := httputil.DumpRequest(r, true)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		log.Printf("%v\n", string(reqDump))
// 	}

// }

// /*
// send login request:
// curl -X POST -d '{"email":"flo","passwd":"1234"}' localhost:8081/login

// get response with token and get restricted access:
// curl localhost:8081/restricted/user/bla -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluIn0.A9dz8H4vRCdMb39m6nOlnl_HbF5zgof5LrLm2i0xEY0"

// */

// type LoginRequestData struct {
// 	Email  string `json:"email"`
// 	Passwd string `json:"passwd"`
// }

// func (a *LoginRequestData) Bind(r *http.Request) error {
// 	if a.Email == "" {
// 		return errors.New("missing required email address for login")
// 	}
// 	if a.Passwd == "" {
// 		return errors.New("missing required password for login")
// 	}
// 	return nil
// }

// func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
// 	log.Printf("login handler\n")

// 	if h.Config.LogLevel == "debug" {
// 		reqDump, err := httputil.DumpRequest(r, true)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		log.Printf("%v\n", string(reqDump))
// 	}

// 	d := &LoginRequestData{}

// 	d.Email = r.PostFormValue("email")
// 	d.Passwd = r.PostFormValue("passwd")
// 	log.Printf("login request data %+v\n", d)

// 	// body, err := io.ReadAll(r.Body)
// 	// if err != nil {
// 	// 	respondError(w, http.StatusBadRequest, err.Error())
// 	// }
// 	// json.Unmarshal(body, d)

// 	// get first element from db where email matches
// 	user := model.User{}
// 	res := h.DB.Where(&model.User{Email: d.Email, Active: true}).First(&user)
// 	if res.Error != nil {
// 		respondError(w, http.StatusInternalServerError, res.Error.Error())
// 		return
// 	}

// 	// verify password
// 	if err := user.CheckPassword(d.Passwd); err != nil {
// 		respondError(w, http.StatusUnauthorized, fmt.Sprintf("wrong password for email '%v'", d.Email))
// 		return
// 	}

// 	// generate jwt token for valid email address
// 	var tokenAuth *jwtauth.JWTAuth
// 	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil, jwt.WithAcceptableSkew(30*time.Second))
// 	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"email": d.Email})
// 	//	log.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)
// 	//	log.Printf("DEBUG: a sample jwt is %s\n\n", t)

// 	w.Write([]byte(tokenString))

// }

// func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
// 	log.Printf("getUser handler\n")
// 	w.Write([]byte(fmt.Sprintf("get user handler id %v", chi.URLParam(r, "id"))))
// }

// func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
// 	respondError(w, http.StatusNotImplemented, "delete user handler")
// 	//log.Printf("getUser handler\n")
// 	//w.Write([]byte(fmt.Sprintf("get user handler id %v", chi.URLParam(r, "id"))))
// }

// /*
// type ErrResponse struct {
// 	Err            error `json:"-"` // low-level runtime error
// 	HTTPStatusCode int   `json:"-"` // http response status code

// 	StatusText string `json:"status"`          // user-level status message
// 	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
// 	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
// }

// func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
// 	render.Status(r, e.HTTPStatusCode)
// 	return nil
// }

// func ErrStatusInternalServerError(err error) render.Renderer {
// 	return &ErrResponse{
// 		Err:            err,
// 		HTTPStatusCode: http.StatusInternalServerError,
// 		StatusText:     http.StatusText(http.StatusInternalServerError),
// 		ErrorText:      err.Error(),
// 	}
// }

// func ErrStatusUnauthorized(err error) render.Renderer {
// 	return &ErrResponse{
// 		Err:            err,
// 		HTTPStatusCode: http.StatusUnauthorized,
// 		StatusText:     http.StatusText(http.StatusUnauthorized),
// 		ErrorText:      err.Error(),
// 	}
// }

// func ErrInvalidRequest(err error) render.Renderer {
// 	return &ErrResponse{
// 		Err:            err,
// 		HTTPStatusCode: http.StatusBadRequest,
// 		StatusText:     http.StatusText(http.StatusBadRequest),
// 		ErrorText:      err.Error(),
// 	}
// }
// */
// // import (
// // 	"net/http"
// // 	"time"

// // 	"github.com/floriwan/srcm/pkg/db/model"
// // 	"github.com/golang-jwt/jwt/v5"
// // 	"github.com/labstack/echo/v4"
// // )

// // /*

// //  send request
// //  curl -X POST -d 'email=test@test.de' -d 'password=shhh!' localhost:1323/login

// //  response:
// //  	{
// //   		"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0NjE5NTcxMzZ9.RB3arc4-OyzASAaUhC2W3ReWaXAt_z2Fd3BN4aWTgEY"
// // 	}

// //  auth request:
// //  curl localhost:1323/restricted -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0NjE5NTcxMzZ9.RB3arc4-OyzASAaUhC2W3ReWaXAt_z2Fd3BN4aWTgEY"

// //  jwtCustomClaims are custom claims extending default ones.
// //  See https://github.com/golang-jwt/jwt for more examples

// // */

// // type JwtCustomClaims struct {
// // 	Email string `json:"email"`
// // 	Admin bool   `json:"admin"`
// // 	jwt.RegisteredClaims
// // }

// // func (h *Handler) Register(c echo.Context) error {
// // 	return c.String(http.StatusNotImplemented, "")
// // }

// // func (h *Handler) Login(c echo.Context) error {
// // 	email := c.FormValue("email")
// // 	password := c.FormValue("password")

// // 	// read user from database
// // 	u := model.User{Email: email}
// // 	res := h.DB.First(&u)

// // 	// user not found error
// // 	if res.Error != nil {
// // 		return echo.ErrNotFound
// // 	}

// // 	// unauthorized error
// // 	if err := u.CheckPassword(password); err != nil {
// // 		return echo.ErrUnauthorized
// // 	}

// // 	// Set custom claims
// // 	claims := &JwtCustomClaims{
// // 		Email: "Florian",
// // 		Admin: true,
// // 		RegisteredClaims: jwt.RegisteredClaims{
// // 			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
// // 		},
// // 	}

// // 	// Create token with claims
// // 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// // 	// Generate encoded token and send it as response.
// // 	t, err := token.SignedString([]byte("secret"))
// // 	if err != nil {
// // 		return err
// // 	}

// // 	return c.JSON(http.StatusOK, echo.Map{
// // 		"token": t,
// // 	})
// // }

// // func (h *Handler) CreateUser(c echo.Context) error {
// // 	// Get name and email
// // 	email := c.FormValue("email")
// // 	passwd := c.FormValue("passwd")
// // 	return c.String(http.StatusOK, "new user email:"+email+" passwd:"+passwd)
// // }

// // func (h *Handler) GetUser(c echo.Context) error {
// // 	// User ID from path `users/:id`
// // 	id := c.Param("id")
// // 	return c.String(http.StatusOK, "get user: "+id)
// // }

// // func (h *Handler) GetAllUsers(c echo.Context) error {
// // 	return c.String(http.StatusNotImplemented, "")
// // }

// // func (h *Handler) DeleteUser(c echo.Context) error {
// // 	return c.String(http.StatusNotImplemented, "")
// // }

// // func (h *Handler) UpdateUser(c echo.Context) error {
// // 	id := c.Param("id")
// // 	return c.String(http.StatusNotImplemented, "update user: "+id)
// // }
