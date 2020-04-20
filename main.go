package main

import (
	"log"
	"os"

	bheap "github.com/pennz/DataViz/viz/trees/binaryheap"

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
	heap := bheap.NewWithIntComparatorV()
	heap.EnableV()

	heap.Push(3)
	heap.Push(19)
	heap.Push(19)
	heap.Push(19)
	heap.Push(19)
	heap.Push(19)
	heap.Push(19)
	//var hv utils.Stepper
	//hv = heap
	gs, err := heap.SSteps()
	if err != nil {
		log.Println("graph genarion error")
	} else {
		for _, g := range gs {
			log.Println(g)
			log.Println()
		}
	}

	r := setupRouter()
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	r.Run(":" + port)
}
