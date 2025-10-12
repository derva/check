package parse

import (
	"bytes"
	"check/config"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/jackc/pgx/v5"
)

type LineItem2 struct {
	Description string  `json:"description"`
	Total       float64 `json:"total"`
}

type Vendor struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type VeryfiResponse struct {
	ID           int         `json:"id"`
	Vendor       Vendor      `json:"vendor"`
	Date         string      `json:"date"`
	Total        float64     `json:"total"`
	CurrencyCode string      `json:"currency_code"`
	LineItems    []LineItem2 `json:"line_items"`
}

func SendToVerify(filePath string) (string, error) {
	apiKey := config.EnvVars.VeryfiAPIKEY
	VeryfiURL := config.EnvVars.VeryfiURL
	clientID := config.EnvVars.VeryfiClientID
	username := config.EnvVars.VeryfiAPIUsername

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return "", err
	}
	if _, err = io.Copy(part, file); err != nil {
		return "", err
	}

	writer.Close()

	req, err := http.NewRequestWithContext(context.Background(), "POST", VeryfiURL, body)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Client-Id", clientID)
	req.Header.Set("Authorization", fmt.Sprintf("apikey %s:%s", username, apiKey))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal(data, &result)

	formatted, _ := json.MarshalIndent(result, "", "  ")
	return string(formatted), nil
}

func SaveItemToDatabase(item config.LineItem) {
	conString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", config.EnvVars.DBName, config.EnvVars.DBPassword, config.EnvVars.DBHost, config.EnvVars.DBPort, config.EnvVars.DBName)
	conn, err := pgx.Connect(context.Background(), conString)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())
	_, err = conn.Exec(context.Background(), `INSERT INTO checks (description, price_of_one, total_price) VALUES ($1, $2, $3)`, item.Description, item.Price, item.Total)
	if err != nil {
		log.Fatal("Error while trying to insert new value", err)
	}

}
