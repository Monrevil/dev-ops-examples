package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	port = flag.Int("port", 8080, "The server port")
	adr  = flag.String("adr", "http://localhost:8081/count", "Address of the count service")
)

func main() {
	flag.Parse()
	router := gin.Default()

	// This handler will match /user/john but will not match /user/ or /user
	router.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		resp := get(*adr)
		log.Printf("[INFO] Got response from counter: %v", resp)
		c.String(http.StatusOK, "Hello %s", name)
	})
	router.GET("/health/ready", func(c *gin.Context) {
		c.String(http.StatusOK, "Service is ready")
	})

	router.Run(fmt.Sprintf(":%d", *port))
}

// Call http endpoint, return response body as string
func get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(body)
}
