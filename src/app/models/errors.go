package models

type ApiError struct {
    Message string `json:"message"`
    Code int `json:"code"`
} 

type ErrorResponse struct {
    Error ApiError `json:"error"`
} 
