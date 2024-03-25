package models

import "fmt"

type ApiError struct {
    Message string `json:"message"`
    Code int `json:"code"`
} 

type ErrorResponse struct {
    Error ApiError `json:"error"`
} 

func (e ApiError) Error() string {
    return fmt.Sprintf("code: %d, message: %s", e.Code, e.Message)
} 
