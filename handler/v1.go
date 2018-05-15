package handler

import (
	"fmt"
	"multi-langs/controller"
	"multi-langs/model"
	"multi-langs/utils"
	"net/http"
	"time"

	"golang.org/x/net/websocket"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

func (handle Handler) Home(c echo.Context) error {

	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"name": "MULTI LANGS!",
	})
}

func (handle Handler) PostLogin(c echo.Context) error {
	// var login model.Login
	params := map[string]interface{}{}

	// if err := c.Bind(&login); err == nil {

	// 	if login.Username != "" && login.Password != "" {

	// 		// Find by username and password
	// 		var users model.Users
	// 		handle.DB.Where(&model.Users{Username: login.Username}).First(&users)

	// 		if utils.ComparePassword(users.Password, login.Password) {

	// 			var auth model.Authorities
	// 			handle.DB.Where("user_id = ?", users.ID).Find(&auth)

	// 			admin := false
	// 			if auth.Authority == utils.ROLE_ADMIN {
	// 				admin = true
	// 			}
	// 			t, err := CreateJwt("prongbang", admin)

	// 			// Create or Update Access Token
	// 			var accessToken model.AccessToken
	// 			handle.DB.Where("user_id = ?", users.ID).First(&accessToken)
	// 			if accessToken.ID == 0 {
	// 				var atLast model.AccessToken
	// 				handle.DB.Last(&atLast)
	// 				var id int64 = 1
	// 				if atLast.ID > 0 {
	// 					id = atLast.ID + 1
	// 				}
	// 				handle.DB.Create(&model.AccessToken{
	// 					ID:     id,
	// 					UserID: users.ID,
	// 					Token:  t,
	// 				})
	// 			} else {
	// 				handle.DB.Model(&accessToken).Update(model.AccessToken{
	// 					Token: t,
	// 				})
	// 			}

	// 			if err == nil {
	// 				return c.JSON(http.StatusOK, echo.Map{
	// 					"message": "success",
	// 					"token":   t,
	// 				})
	// 			}
	// 			params["error"] = "Login fail!"
	// 		} else {
	// 			params["error"] = "Email or Password incorrect"
	// 		}
	// 	} else if login.Username == "" && login.Password == "" {
	// 		params["error"] = "Please Enter Email and Password"
	// 	} else if login.Username == "" {
	// 		params["error"] = "Please Enter Email"
	// 	} else if login.Password == "" {
	// 		params["error"] = "Please Enter Password"
	// 	}
	// }
	return c.JSON(http.StatusBadRequest, params)
}

func (handle Handler) GetApplication(c echo.Context) error {

	var apps []model.Application
	// handle.DB.Find(&apps)

	return c.JSON(http.StatusOK, apps)
}

// START LANGUAGES ----------------------------------------------------------------------------------------------------------
func (handle Handler) PostLanguage(c echo.Context) error {

	var maps echo.Map
	if c.Bind(&maps) == nil {
		fmt.Println(maps)
		response, err := controller.Controller{c, handle.RTDb}.PostLanguage(maps)
		if err == nil {
			return c.JSON(http.StatusOK, response)
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	return c.JSON(http.StatusBadRequest, echo.Map{
		"message": "Bad Request",
	})
}

func (handle Handler) DeleteLanguage(c echo.Context) error {
	id := c.Param("id")
	if id != "" {
		rows, err := controller.Controller{c, handle.RTDb}.DeleteLanguage(id)

		// check response
		if err != nil {
			return c.JSON(http.StatusNotFound, err)
		}
		return c.JSON(http.StatusOK, rows)
	}
	return c.JSON(http.StatusBadRequest, echo.Map{
		"message": "Bad Request",
	})
}

func (handle Handler) GetLanguageById(c echo.Context) error {

	key := c.QueryParam("app")
	lang := c.QueryParam("lang")
	id := c.Param("id")

	fmt.Println(key, lang, id)

	rows := controller.Controller{c, handle.RTDb}.GetLanguage(id, key, lang)

	// check response
	if len(rows) == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{})
	} else if len(rows) == 1 {
		if rows[0] == nil {
			return c.JSON(http.StatusNotFound, echo.Map{})
		}
		return c.JSON(http.StatusOK, rows[0])
	}

	return c.JSON(http.StatusOK, rows)
}

func (handle Handler) GetLanguage(c echo.Context) error {

	key := c.QueryParam("app")
	lang := c.QueryParam("lang")
	id := c.QueryParam("id")

	fmt.Println(key, lang, id)

	rows := controller.Controller{c, handle.RTDb}.GetLanguage(id, key, lang)

	// check response
	if len(rows) == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{})
	} else if len(rows) == 1 {
		if rows[0] == nil {
			return c.JSON(http.StatusNotFound, echo.Map{})
		}
		return c.JSON(http.StatusOK, rows[0])
	}

	return c.JSON(http.StatusOK, rows)
}

func (handle Handler) GetAppLanguage(c echo.Context) error {

	key := c.QueryParam("app")
	lang := c.QueryParam("lang")

	fmt.Println(key, lang)

	var appLangs []model.AppLangs

	return c.JSON(http.StatusOK, appLangs)
}

// END LANGUAGES ----------------------------------------------------------------------------------------------------------

// START ATTRIBUTES -------------------------------------------------------------------------------------------------------
func (handle Handler) DeleteAppAttributes(c echo.Context) error {
	id := c.Param("id")
	if id != "" {
		rows, err := controller.Controller{c, handle.RTDb}.DeleteAppAttributes(id)

		// check response
		if err != nil {
			return c.JSON(http.StatusNotFound, err)
		}
		return c.JSON(http.StatusOK, rows)
	}
	return c.JSON(http.StatusBadRequest, echo.Map{
		"message": "Bad Request",
	})
}

func (handle Handler) PutAppAttributes(c echo.Context) error {
	var maps echo.Map
	id := c.Param("id")
	if c.Bind(&maps) == nil && id != "" {

		response, err := controller.Controller{c, handle.RTDb}.PutAppAttribute(id, maps)
		if err == nil {
			return c.JSON(http.StatusOK, response)
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	return c.JSON(http.StatusBadRequest, echo.Map{
		"message": "Bad Request",
	})
}

func (handle Handler) PostAppAttributes(c echo.Context) error {

	var maps echo.Map
	if c.Bind(&maps) == nil {

		response, err := controller.Controller{c, handle.RTDb}.PostAppAttribute(maps)
		if err == nil {
			return c.JSON(http.StatusOK, response)
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	return c.JSON(http.StatusBadRequest, echo.Map{
		"message": "Bad Request",
	})
}

func (handle Handler) GetAppAttributesById(c echo.Context) error {
	key := c.QueryParam("app")
	lang := c.QueryParam("lang")
	id := c.Param("id")

	fmt.Println(key, lang, id)

	rows := controller.Controller{c, handle.RTDb}.GetAttribute(id, key, lang)

	// check response
	if len(rows) == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{})
	} else if len(rows) == 1 {
		if rows[0] == nil {
			return c.JSON(http.StatusNotFound, echo.Map{})
		}
		return c.JSON(http.StatusOK, rows[0])
	}

	return c.JSON(http.StatusOK, rows)
}

func (handle Handler) GetAppAttributes(c echo.Context) error {
	key := c.QueryParam("app")
	lang := c.QueryParam("lang")
	id := c.QueryParam("id")

	fmt.Println(key, lang, id)

	rows := controller.Controller{c, handle.RTDb}.GetAttribute(id, key, lang)

	// check response
	if len(rows) == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{})
	} else if len(rows) == 1 {
		if rows[0] == nil {
			return c.JSON(http.StatusNotFound, echo.Map{})
		}
		return c.JSON(http.StatusOK, rows[0])
	}
	return c.JSON(http.StatusOK, rows)
}

// END ATTRIBUTES -------------------------------------------------------------------------------------------------------

func (handle Handler) GetAuthorities(c echo.Context) error {

	var auth []model.Authorities
	// handle.DB.Find(&auth)

	return c.JSON(http.StatusOK, auth)
}

func (handle Handler) GetAccessToken(c echo.Context) error {
	var token []model.AccessToken
	// handle.DB.Find(&token)

	return c.JSON(http.StatusOK, token)
}

func (handle Handler) GetMember(c echo.Context) error {

	var users []model.Users
	// handle.DB.Find(&users)

	return c.JSON(http.StatusOK, users)
}

func (handle Handler) GetMemberProfile(c echo.Context) error {

	var users []model.UserProfile
	// handle.DB.Find(&users)

	return c.JSON(http.StatusOK, users)
}

func (handle Handler) GetUserProfile(c echo.Context) error {

	var users []model.UserProfile
	// handle.DB.Find(&users)

	return c.JSON(http.StatusOK, users)
}

func (handle Handler) GetAddMember(c echo.Context) error {

	return c.JSON(http.StatusOK, echo.Map{"message": "Hello", "time": time.Now()})
}

func (handle Handler) GenToken(c echo.Context) error {

	return c.JSON(http.StatusOK, echo.Map{"message": "Hello", "time": time.Now()})
}

func (handle Handler) RefreshToken(c echo.Context) error {
	t, _ := CreateJwt("prongbang", true)
	return c.JSON(http.StatusOK, echo.Map{"token": t})
}

func (handle Handler) Process(c echo.Context) error {
	text := c.QueryParam("q")
	if text == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Query Not found!"})
	}

	msg := model.Message{}
	msg.Value = text

	return c.JSON(http.StatusOK, msg)
}

func (handle Handler) Ws(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		for {
			// Write
			err := websocket.Message.Send(ws, "Hello, Client!")
			if err != nil {
				c.Logger().Error(err)
			}

			// Read
			msg := ""
			err = websocket.Message.Receive(ws, &msg)
			if err != nil {
				c.Logger().Error(err)
			}
			fmt.Printf("%s\n", msg)
		}
	}).ServeHTTP(c.Response(), c.Request())

	return nil
}

func CreateJwt(name string, admin bool) (string, error) {
	// Set custom claims
	claims := &jwtCustomClaims{
		name,
		admin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30 * 12).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(utils.JWT_SECRET_KEY))
	if err != nil {
		return "", err
	}
	return t, nil
}
