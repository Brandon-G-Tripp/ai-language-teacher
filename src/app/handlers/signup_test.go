package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/Brandon-G-Tripp/ai-language-teacher/internal/testutil"
	"github.com/Brandon-G-Tripp/ai-language-teacher/src/database"

	handler_models "github.com/Brandon-G-Tripp/ai-language-teacher/src/app/models"
	db_models "github.com/Brandon-G-Tripp/ai-language-teacher/src/database/models"
)

type SignUpTestCase struct {
    Input handler_models.SignUpRequest
    StatusCode int
    Response interface{} `json:"response"`
} 

func TestMain(m *testing.M) {
    // Init DB
    test_db := testutil.InitTestDB()
    database.DB = test_db
    // run tests
    exitCode := m.Run()

    // Close connection 
    testutil.CloseTestDB(database.DB)

    os.Exit(exitCode)
} 


func TestSignUpHandler(t *testing.T) {


    // Define test cases
    cases := map[string]SignUpTestCase{
        "valid": {
            Input: handler_models.SignUpRequest{
                Name:  "John",
                Email: "john@doe.com",   
                Password: "password",
            },
            StatusCode: http.StatusOK, 
            Response: handler_models.SignUpResponse{
                User: db_models.User{
                    ID:    1,
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
        SignUp(ctx)
        if ctx.Errors != nil {
            ctx.JSON(400, ctx.Errors)
            ctx.Writer.WriteHeader(400)
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
