package main

import (
	"net/http"
	"strconv"

	"github.com/aleskeinovikov/playground/internal"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		sum := internal.Sum[int64](1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		c.String(http.StatusOK, strconv.Itoa(int(sum)))
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
