package main

import (
	"html/template"
	"io"
	"multi-langs/controller"
	"multi-langs/db"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	ctrl := controller.Controller{db.Open()}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e.Renderer = renderer
	e.Static("/", "../public")
	e.GET("/", ctrl.Home)
	e.GET("/v1/token", ctrl.GenToken)
	e.GET("/v1/refresh-token", ctrl.RefreshToken)

	g := e.Group("/v1", middleware.JWT([]byte("noodang-secret-l1ackme-pls")))
	g.GET("/ws", ctrl.Ws)
	g.GET("/process", ctrl.Process)

	e.Logger.Fatal(e.Start(":" + port))
}

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}
