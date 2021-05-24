package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/thing/{uuid}", GetThing)
	//r.PUT("/thing/{uuid}", UpdateThing)

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}

type ThingResponse struct {
	UUID  string `json:"uuid"`
	Name  string `json:"name"`
	Value string `json:"value"`

	Updated time.Time `json:"updated"`
	Created time.Time `json:"created"`
}

func GetThing(c *gin.Context) {
	uuid := c.Param("uuid")

	thingResponse := ThingResponse{
		UUID:  uuid,
		Name:  "example",
		Value: "example",
	}
	c.JSON(http.StatusOK, thingResponse)
}
