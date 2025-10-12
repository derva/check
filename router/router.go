package router

import (
	"check/config"
	"check/parse"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func HandleSave(c *gin.Context) {
	var outer config.OuterData
	if err := c.BindJSON(&outer); err != nil {
		fmt.Println("failed to decode outer data")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not decode outer data"})
		return
	}

	for _, element := range outer.Text.LineItems {
		parse.SaveItemToDatabase(element)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data processed successfully!"})
}

func HandleUpload(c *gin.Context) {
	file, err := c.FormFile("photo")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tmpPath := filepath.Join("static/uploads", filepath.Base(file.Filename))
	if err := c.SaveUploadedFile(file, tmpPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ocrText, err := parse.SendToVerify(tmpPath) // your OCR logic
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Receipt processed successfully!",
		"text":    ocrText,
	})
}
