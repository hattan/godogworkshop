package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	animals "github.com/hattan/godog"
	"github.com/hattan/godogapi/pkg/dogs"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/dog", func(c *gin.Context) {
		c.JSON(http.StatusOK, dogs.GetDogs())
	})

	r.GET("/dog/:name", func(c *gin.Context) {
		name := c.Params.ByName("name")
		dog := dogs.GetDogByName(name)
		c.JSON(http.StatusOK, dog)
	})

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":   "bar",
		"rover": "woof",
	}))

	authorized.POST("dog", func(c *gin.Context) {
		dog := animals.Dog{}
		bindErr := c.Bind(&dog)
		if bindErr != nil {
			slog.Error(bindErr.Error())
			return
		}

		// Add dog to in memory db

		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	authorized.DELETE("/dog/:name", func(c *gin.Context) {
		name := c.Params.ByName("name")
		err := dogs.RemoveDog(name)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "godo is up!")
	})
	return r
}

func main() {
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}
	r := setupRouter()
	r.Run(":" + port)
}
