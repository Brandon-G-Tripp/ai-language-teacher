package models

import "github.com/gin-gonic/gin"

type ApiError struct {
    Message string
    Error gin.Error
} 
