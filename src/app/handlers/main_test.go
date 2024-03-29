package handlers

import (
    "os"
    "testing"

    "github.com/Brandon-G-Tripp/ai-language-teacher/internal/testutil"
    "github.com/Brandon-G-Tripp/ai-language-teacher/src/database"
    "github.com/Brandon-G-Tripp/ai-language-teacher/src/app/services/auth"
    gorm_repos "github.com/Brandon-G-Tripp/ai-language-teacher/src/database/repositories"
)

var (
    signUpHandler *SignUpHandler
    loginHandler *LoginHandler
    logoutHandler *LogoutHandler
    conversationHandler *ConversationHandler
    authService *auth.AuthService
    userRepo *gorm_repos.UserRepository
    conversationRepo *gorm_repos.ConversationRepository
)

func TestMain(m *testing.M) {
    // Init DB
    test_db := testutil.InitTestDB()
    database.DB = test_db
    // Setup handler
    userRepo = gorm_repos.NewUserRepository(database.DB)
    conversationRepo = gorm_repos.NewConversationRepository(database.DB)
    authService = auth.NewAuthService()
    signUpHandler = NewSignUpHandler(userRepo, authService)
    loginHandler = NewLoginHandler(userRepo, authService)
    logoutHandler = NewLogoutHandler(userRepo, authService)

    // run tests
    exitCode := m.Run()

    // Close connection 
    testutil.CloseTestDB(database.DB)

    os.Exit(exitCode)
} 
