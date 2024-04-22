package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"testing"

	"github.com/Brandon-G-Tripp/ai-language-teacher/internal/testutil"

	handler_models "github.com/Brandon-G-Tripp/ai-language-teacher/src/app/models"
	db_models "github.com/Brandon-G-Tripp/ai-language-teacher/src/database/models"
)

type SignUpTestCase struct {
    Input handler_models.SignUpRequest
    StatusCode int
    Response interface{} `json:"response"`
} 

func TestSignUpHandler(t *testing.T) {
    // Setup 

    hashedPassword, err := authService.HashPassword("password")
    if err != nil {
        t.Fatal("Error returned from hashing password")
    }
    // Define test cases
    cases := map[string]SignUpTestCase{
        "valid": {
            Input: handler_models.SignUpRequest{
                Name:  "John",
                Email: "john@doe.com",   
                Password: hashedPassword,
            },
            StatusCode: http.StatusOK, 
            Response: handler_models.SignUpResponse{
                User: &db_models.User{
                    ID:    0,
                    Name:  "John Doe",
                    Email: "john@doe.com",
                },
                Token: "stub-token",
            },
        },
        "invalid email": {
            Input: handler_models.SignUpRequest{
                Email: "invalid",
                Password: "password",
            },
            StatusCode: http.StatusBadRequest,
            Response: handler_models.ErrorResponse{
                Error: handler_models.ApiError{
                    Message: "Invalid email address",
                    Code: 400,
                },
            },
        },
        // Other test cases for different scenarios
    }


    // Setup 


    // Run through test cases
    for _, tc := range cases {
        // Arrange
        ctx := testutil.GetTestGinContext()

        // Create http request
        data, _ := json.Marshal(tc.Input)
        req, err := http.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewReader(data))
        if err != nil {
            log.Printf("There was an error while creating new Request in test: %v", err)
            ctx.AbortWithStatus(400)
        } 

        ctx.Request = req
        
        // Act
        user, token, err := signUpHandler.SignUp(tc.Input)
        if err != nil {
            apiErr, ok := err.(handler_models.ApiError)
            if !ok {
                ctx.JSON(http.StatusInternalServerError, "SignUp Handler - Internal Server error")
            } else {
                ctx.JSON(apiErr.Code, "Error from apiErr in Signup")
            }
        } else {
            ctx.JSON(http.StatusOK, handler_models.SignUpResponse{
                User: user,
                Token: token, 
            })
        } 

        // Assert
        // Validate status code
        if got, want := ctx.Writer.Status(), tc.StatusCode; got != want {
            t.Errorf("Unexpected status code: got %v want %v", got, want)
        } 
        
        // Decode response
        var response interface{}
        err = json.NewDecoder(ctx.Request.Body).Decode(&response)
        log.Printf("Response type: %T", response)
        if err != nil {
            if got, want := ctx.Writer.Status(), tc.StatusCode; got != want {
                t.Errorf("Invalid status code")
            } 
        } 


        
        // Validate fields
        if errorResp, ok := response.(handler_models.ErrorResponse); ok {
            if errorResp.Error.Message != "Invalid email address" {
                t.Errorf("Validation failed: %v", errorResp)
            } 
        }
        
    }
}
