package database

func AutoMigrate() {
    db := connectDB()
    db.AutoMigrate(&User{})
}
