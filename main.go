package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pennz/DataViz/viz"
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

func readCloser2SVG(rc io.ReadCloser) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(rc)
	newStr := viz.Dot2SVG(buf.String())
	return newStr
}

func read2buf(rc io.ReadCloser) *bytes.Buffer {
	buf := new(bytes.Buffer)
	buf.ReadFrom(rc)
	return buf
}

func compileHandler_debug(c *gin.Context) {
	// version := c.PostForm("version")
	body := c.PostForm("body")
	withVet := c.PostForm("withVet")
	log.Println(version, body, withVet)
}

// compileHandler will relay the request to play.golang.org
func compileHandler(c *gin.Context) {
	//log.Printf("%v\n", readCloser2String(c.Request.Body))
	// https://github.com/gin-gonic/gin#try-to-bind-body-into-different-structs
	// The normal methods for binding request body consumes c.Request.Body and
	// they cannot be called multiple times.
	//buf := read2buf(c.Request.Body)
	// we can change the body in the go
	body := c.PostForm("body")
	s := fmt.Sprintf("version=%d&body=%s&withVet=%s", 2, url.QueryEscape(body), "true")
	log.Println(s)
	buf := bytes.NewBufferString(s)
	var relay io.Reader = bytes.NewReader(buf.Bytes())
	response, err := http.Post("https://play.golang.org/compile", "application/x-www-form-urlencoded; charset=UTF-8", relay)
	if err == nil {
		if response.StatusCode == 200 && response.Body != nil {
			c.String(response.StatusCode, readCloser2SVG(response.Body))
		} else {
			c.String(response.StatusCode, "Error or cannot get response from play.golang.org.")
		}
	} else {
		c.String(404, "Cannot access play.golang.org.")
	}
}
