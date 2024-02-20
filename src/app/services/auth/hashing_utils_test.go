package auth

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestVerifyPassword(t *testing.T) {
    // Generate a hashed password
    hashedPassword, err := hashPassword("password")
    if err != nil {
        t.Fatal(err)
    } 

    // Verify it 
    err = verifyPassword(hashedPassword, "password")
    if err != nil {
        t.Errorf("Expected no error verifying password, but got: %v", err)
    }

    // verify incorrect password fails
    err = verifyPassword(hashedPassword, "wrongpassword")
    if err == nil {
        t.Error("Expected error verifying incorrect password, but got nil")
    } 
}

func TestHashPasswor(t *testing.T) {
    hashed, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
    if err != nil {
        t.Error("Error using bcrypt")
    } 

    testCases := []struct {
        name     string
        password string
        hashed   string
        wantErr  bool
    }{
        {
            name:     "valid password",
            password: "password",
            hashed: string(hashed), // pre-hashed
            wantErr: false, 
        },
        {
            name:     "invalid password", 
            password: "wrongpassword",
            hashed: string(hashed), // pre-hashed
            wantErr: true,
        },
        // additional test cases...
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            err := verifyPassword(tc.hashed, tc.password)
            if tc.wantErr {
                if err == nil {
                    t.Errorf("expected error, but got nil")
                }
            } else {
                if err != nil {
                    t.Errorf("expected no error, but got: %v", err)
                } 
            }
        })
    }
}
