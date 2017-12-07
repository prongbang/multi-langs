package controller

import (
	"database/sql"
	"fmt"
	"multi-langs/model"
	"net/http"
	"time"

	"golang.org/x/net/websocket"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type Controller struct {
	DB *sql.DB
}

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

func (ctrl Controller) Home(c echo.Context) error {

	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"name": "Dolly!",
	})
}

func (ctrl Controller) GenToken(c echo.Context) error {

	fmt.Println(ctrl.DB.Ping)

	return c.JSON(http.StatusOK, echo.Map{"message": "Hello", "time": time.Now()})
}

func (ctrl Controller) RefreshToken(c echo.Context) error {
	// Set custom claims
	claims := &jwtCustomClaims{
		"prongbang",
		true,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("noodang-secret-l1ackme-pls"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{"token": t})
}

func (ctrl Controller) Process(c echo.Context) error {
	text := c.QueryParam("q")
	if text == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Query Not found!"})
	}

	msg := model.Message{}
	msg.Value = text

	return c.JSON(http.StatusOK, msg)
}

func (ctrl Controller) Ws(c echo.Context) error {
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
