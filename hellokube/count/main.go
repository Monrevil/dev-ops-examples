package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync/atomic"

	"github.com/gin-gonic/gin"
)

var (
	port  = flag.Int("port", 8081, "The server port")
	count uint64
)

func main() {
	flag.Parse()
	router := gin.Default()

	router.GET("/count", func(c *gin.Context) {
		atomic.AddUint64(&count, 1)
		c.String(http.StatusOK, "Said hello %d times", count)
	})
	router.GET("/health/ready", func(c *gin.Context) {
		c.String(http.StatusOK, "Service is ready")
	})

	router.Run(fmt.Sprintf(":%d", *port))
}
