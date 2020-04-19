package main

import (
	"log"
	"os"

	bheap "github.com/pennz/dataviz/trees/binaryheap"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.Static("/Viz", "GraphVizOnline")

	//r.LoadHTMLGlob("GraphVizOnline/*.html")
	//r.GET("/", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "index.html", nil)
	//})
	return r
}

func main() {
	heap := bheap.NewWithIntComparator()
	heap.Push(3)
	heap.Push(19)
	heap.Push(19)
	heap.Push(19)
	heap.Push(19)
	heap.Push(19)
	heap.Push(19)
	heap.Visualizer("heap.png")

	r := setupRouter()
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	r.Run(":" + port)
}
