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
    messageHandler *MessageHandler
    authService *auth.AuthService
    userRepo *gorm_repos.UserRepository
    conversationRepo *gorm_repos.ConversationRepository
    messageRepo * gorm_repos.MessageRepository
)

func TestMain(m *testing.M) {
    // Init DB
    test_db := testutil.InitTestDB()
    database.DB = test_db
    
    // Setup Repos 
    userRepo = gorm_repos.NewUserRepository(database.DB)
    conversationRepo = gorm_repos.NewConversationRepository(database.DB)
    messageRepo = gorm_repos.NewMessageRepository(database.DB)

    // Setup handler
    authService = auth.NewAuthService()
    signUpHandler = NewSignUpHandler(userRepo, authService)
    loginHandler = NewLoginHandler(userRepo, authService)
    logoutHandler = NewLogoutHandler(userRepo, authService)
    conversationHandler =  NewConversationHandler(conversationRepo, authService)
    messageHandler = NewMessageHandler(messageRepo, authService)

    // run tests
    exitCode := m.Run()

    // Close connection 
    testutil.CloseTestDB(database.DB)

    os.Exit(exitCode)
} 
