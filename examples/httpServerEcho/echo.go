package main

import (
	"log"
	"net/http"
	"time"

	echo "github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/thing/{uuid}", GetThing)
	//e.PUT("/thing/{uuid}", UpdateThing)
	if err := e.Start("localhost:8080"); err != nil {
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

func GetThing(e echo.Context) error {
	uuid := e.Param("uuid")

	thingResponse := ThingResponse{
		UUID:  uuid,
		Name:  "example",
		Value: "example",
	}
	return e.JSON(http.StatusOK, thingResponse)
}
