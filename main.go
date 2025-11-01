package main

import (
	"check/config"
	"check/router"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	r := gin.Default()
	LoadEnv()
	// Serve the HTML page
	r.LoadHTMLGlob("templates/*")

	// Optional: serve static files if you want to access uploads
	r.Static("/static", "./static")

	// Main route for the upload page
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about.html", nil)
	})

	r.GET("/upload", func (c *gin.Context) {
		c.HTML(http.StatusOK, "upload_in_progress.html", nil)
	})

	// Upload endpoint
	r.POST("/upload", router.HandleUpload)
	r.POST("/save", router.HandleSave)

	// Run the server
	fmt.Println("🚀 Running server")
	r.Run(":8080")

}

func LoadEnv() {
	_ = godotenv.Load()

	config.EnvVars = config.EnvVariables{
		VeryfiAPIKEY:      os.Getenv("VeryfiAPI_KEY"),
		VeryfiAPIUsername: os.Getenv("VeryfiAPI_USERNAME"),
		VeryfiClientID:    os.Getenv("VeryfiClientID"),
		VeryfiURL:         os.Getenv("VeryfiURL"),
		DBHost:            os.Getenv("DBHost"),
		DBUsername:        os.Getenv("DBUsername"),
		DBPassword:        os.Getenv("DBPassword"),
		DBPort:            os.Getenv("DBPort"),
		DBName:            os.Getenv("DBName"),
	}
}
