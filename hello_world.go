package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Person struct {
	Id   string `form:"id" uri:"id" binding:"required"`
	Name string `form:"name"`
}

func main() {
	people := [...]Person{{Id: "ccomper", Name: "Bruno"}}

	server := gin.Default()

	server.GET("/hello", func(context *gin.Context) {
		var person Person

		queryPersonErr := context.Bind(&person)

		if queryPersonErr != nil || person.Name == "" {
			context.JSON(http.StatusOK, gin.H{
				"message": "Hello my friend",
			})

			return
		}

		context.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Hello %s", person.Name),
		})
	})

	server.GET("/hello/:id", func(context *gin.Context) {
		var person Person

		if err := context.ShouldBindUri(&person); err != nil {
			context.JSON(http.StatusNotFound, gin.H{
				"message": fmt.Sprintf("Error with path param binding. %s", err),
			})

			return
		}

		for i := 0; i < len(people); i++ {
			if people[i].Id == person.Id {
				person = people[i]
				break
			}
		}

		if person.Name != "" {
			context.JSON(http.StatusOK, person)
		} else {
			context.JSON(http.StatusNotFound, gin.H{
				"message": fmt.Sprintf("People with Id=%s not found", person.Id),
			})
		}
	})

	serverErr := server.Run()

	if serverErr != nil {
		fmt.Println(serverErr)
		return
	}
}
