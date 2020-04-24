package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.Static("/Viz", "GraphVizOnline")
	r.POST("/compile", compileHandler)
	r.POST("/compile_debug", compileHandler_debug)

	//r.LoadHTMLGlob("GraphVizOnline/*.html")
	//r.GET("/", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "index.html", nil)
	//})
	return r
}

func main() {
	r := setupRouter()
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	r.Run(":" + port)
}

func readCloser2String(rc io.ReadCloser) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(rc)
	newStr := buf.String()
	return newStr
}

func compileHandler_debug(c *gin.Context) {
	log.Printf("in deubug%v\n", readCloser2String(c.Request.Body))
}

func compileHandler(c *gin.Context) {
	//log.Printf("%v\n", readCloser2String(c.Request.Body))
	response, _ := http.Post("https://go-algorithm-dev-mqqjokeeppul8.herokuapp.com/compile_debug", "application/x-www-form-urlencoded; charset=UTF-8", c.Request.Body)
	log.Println(readCloser2String(response.Body))
}
