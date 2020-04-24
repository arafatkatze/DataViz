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

func read2buf(rc io.ReadCloser) *bytes.Buffer {
	buf := new(bytes.Buffer)
	buf.ReadFrom(rc)
	return buf
}

func compileHandler_debug(c *gin.Context) {
	log.Printf("in deubug%v\n", readCloser2String(c.Request.Body))
}

func compileHandler(c *gin.Context) {
	//log.Printf("%v\n", readCloser2String(c.Request.Body))
	// https://github.com/gin-gonic/gin#try-to-bind-body-into-different-structs
	// The normal methods for binding request body consumes c.Request.Body and
	// they cannot be called multiple times.
	buf := read2buf(c.Request.Body)
	var relay io.Reader = bytes.NewReader(buf.Bytes())
	response, _ := http.Post("https://go-algorithm-dev-mqqjokeeppul8.herokuapp.com/compile_debug", "application/x-www-form-urlencoded; charset=UTF-8", relay)
	log.Println(readCloser2String(response.Body))
}
