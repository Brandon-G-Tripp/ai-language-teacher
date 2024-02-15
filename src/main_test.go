package main

import (
	"net/http"
	"testing"
	"time"
)

func TestMain(t *testing.T) {
    // Start web server

    go func() {
        main()
    }() 

    // Wait to start 
    time.Sleep(1 * time.Second)

    res, err := http.Get("http://localhost:8080")

    if err != nil {
        t.Fatal(err)
    } 

    if res.StatusCode != 200 {
        t.Fatalf("Expected status code 200, got %d", res.StatusCode)
    } 
} 
