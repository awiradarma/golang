package main

import (
	"os"

	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

func main() {
	m := os.Getenv("MESSAGE")
	h, _ := os.Hostname()
	if m == "" {
		m = "hello"
	}
	r := gin.Default()
	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)
	r.GET("/greetings", func(c *gin.Context) {
		c.String(200, m+" ("+h+")")
	})
	r.Run()
}
