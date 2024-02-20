package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Brandon-G-Tripp/ai-language-teacher/src/database"

	handler_models "github.com/Brandon-G-Tripp/ai-language-teacher/src/app/models"
	db_models "github.com/Brandon-G-Tripp/ai-language-teacher/src/database/models"
	user_repo "github.com/Brandon-G-Tripp/ai-language-teacher/src/database/repositories"
)

func SignUp(c *gin.Context) {
    // parse request
    var req handler_models.SignUpRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(400, err)
        return
    }

    log.Printf("Request: %+v", req)

    // Validate Request
    if !IsEmailValid(req.Email) {
        err := handler_models.ApiError{
            Message: "Invalid email address",
            Code: 400,
        }

        c.AbortWithError(err.Code, err)

        return 
    } 

    // cReate user repo 
    repo := user_repo.NewUserRepository(database.DB)


    // Check for existing email
    user, err := repo.GetByEmail(req.Email)
    if err != nil {
    }

    if user.ID == 0 {
        // user not found 
        // Create user 
        // Hash Password 
        hashed, err := HashPassword(req.Password)
        if err != nil {
            c.JSON(500, err)
            return 
        } 

        user := db_models.User{
            Name: req.Name, 
            Email: req.Email,
            Password: hashed,
        }
        if err := repo.Create(&user); err != nil {
            c.JSON(500, err)
            return 
        } 

        // Generate token 
        token, err := GenerateToken(&user)
        if err != nil {
            c.JSON(500, err)
            return 
        } 

        c.JSON(http.StatusOK, handler_models.SignUpResponse{
            User: user, 
            Token: token,
        })
        return 
    }  else {
        c.JSON(http.StatusBadRequest, "Email already exists")
        return 
    } 
} 
