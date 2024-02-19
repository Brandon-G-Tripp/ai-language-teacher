package handlers

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func GetTestGinContext() *gin.Context {
    gin.SetMode(gin.TestMode)

    w := httptest.NewRecorder()
    ctx, _ := gin.CreateTestContext(w)
    ctx.Request = &http.Request{
        Header: make(http.Header),
    } 

    return ctx
} 
