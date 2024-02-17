package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
    // parse request
    var input SignUpRequest

    if err := c.ShouldBindJSON(&input); err != nil {
        // return error 
    }

    // Todo: Validate input

    // Todo: create user 
    // call to repo create

    // Generate token 

    // return user

    // stub of response
    response := SignUpResponse{
        User: User{
            ID: 1,
            Name: "John Doe",
            Email: input.Email,
        },
        Token: "stub-token",
    }

    c.JSON(http.StatusOK, response)
} 
