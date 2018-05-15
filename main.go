package main

import (
	"html/template"
	"multi-langs/db"
	"multi-langs/handler"
	"multi-langs/utils"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/middleware"
)

func main() {

	port := "9000"
	production := true
	handle := handler.Handler{db.OpenRethink(production)}

	// middleware
	e := echo.New()
	if production {
		file := utils.LoggerConfig("logs", utils.TimeToString(time.Now())+".log")
		e.Debug = !production
		e.Logger.SetOutput(file)
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Output: file,
		}))
	} else {
		e.Use(middleware.Logger()) // /dev/stdout
	}
	e.Use(middleware.Recover())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(utils.COOKIES_SECRET_KEY))))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:" + port},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS},
	}))

	// template
	renderer := &utils.TemplateRenderer{
		Templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e.Renderer = renderer
	e.Static("/", "public")

	// e.GET("/", handle.Home)

	e.GET("/api/v1/refresh-token", handle.RefreshToken)
	e.POST("/api/v1/login", handle.PostLogin)

	g := e.Group("/api/v1", middleware.JWT([]byte(utils.JWT_SECRET_KEY)))

	// attribute
	g.GET("/attributes", handle.GetAppAttributes)
	g.GET("/attributes/:id", handle.GetAppAttributesById)
	g.POST("/attributes", handle.PostAppAttributes)
	g.PUT("/attributes/:id", handle.PutAppAttributes)
	g.DELETE("/attributes/:id", handle.DeleteAppAttributes)

	// language
	g.GET("/language", handle.GetLanguage)
	g.GET("/language/:id", handle.GetLanguageById)
	g.POST("/language", handle.PostLanguage)
	g.PUT("/language/:id", handle.GetLanguage)
	g.DELETE("/language/:id", handle.DeleteLanguage)

	// application
	g.GET("/application", handle.GetApplication)
	g.GET("/application/:id", handle.GetApplication)
	g.POST("/application", handle.GetApplication)
	g.PUT("/application/:id", handle.GetApplication)
	g.DELETE("/application/:id", handle.GetApplication)

	// authorities
	g.GET("/authorities", handle.GetAuthorities)
	g.GET("/authorities/:id", handle.GetAuthorities)
	g.POST("/authorities", handle.GetAuthorities)
	g.PUT("/authorities/:id", handle.GetAuthorities)
	g.DELETE("/authorities/:id", handle.GetAuthorities)

	// access-token
	g.GET("/access-token", handle.GetAccessToken)
	g.GET("/access-token/:id", handle.GetAccessToken)
	g.POST("/access-token", handle.GetAccessToken)
	g.PUT("/access-token/:id", handle.GetAccessToken)
	g.DELETE("/access-token/:id", handle.GetAccessToken)

	// user
	g.GET("/user/:id", handle.GetMember)
	g.POST("/user", handle.GetAddMember)
	g.PUT("/user/:id", handle.GetAddMember)
	g.DELETE("/user/:id", handle.GetAddMember)

	// user profile
	g.GET("/user/profile/:id", handle.GetMemberProfile)
	g.POST("/user/profile", handle.GetMemberProfile)
	g.PUT("/user/profile/:id", handle.GetMemberProfile)
	g.DELETE("/user/profile/:id", handle.GetMemberProfile)

	// Add handler for websocket server
	e.GET("/api/v1/ws/attributes", func(c echo.Context) error {
		return handle.NewChangesHandler(handle.AttributesChanges, c)
	})
	g.GET("/token", handle.GenToken)
	g.GET("/ws", handle.Ws)
	g.GET("/process", handle.Process)
	g.GET("/app-language", handle.GetAppLanguage)

	e.Logger.Fatal(e.Start(":" + port))
}
