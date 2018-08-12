package main

import (
	"os"

	ctr "github.com/heroku/go-getting-started/static/controller"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	//var err error
	port := os.Getenv("PORT")

	if port == "" {
		//log.Fatal("$PORT must be set")
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	//router.GET("/", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "index.tmpl.html", nil)
	//})

	router.GET("/", ctr.GETIndex)
	router.GET("/Collection/:name", ctr.GETAllCollection)
	router.POST("/Collection/POST", ctr.POSTSaveCollection)
	router.GET("/index/JSON", ctr.JSONIndex)
	router.POST("/index/POST", ctr.POSTIndex)

	router.GET("/Teste/GET", ctr.GetTesteCollection)
	router.POST("/Teste/POST", ctr.POSTSaveTesteCollection)

	router.Run(":" + port)
}
