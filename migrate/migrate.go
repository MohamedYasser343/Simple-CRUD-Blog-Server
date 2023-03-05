package main

import (
	"github.com/MohamedYasser343/initializers"
	"github.com/MohamedYasser343/models"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}