package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type MessageRequest struct {
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}

func SendMessage(text string) error {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Ошибка при загрузке .env файла")
	}

	chatID := os.Getenv("TELEGRAM_CHAT_ID")
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")

	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)

	message := MessageRequest{
		ChatID: chatID,
		Text:   text,
	}

	// Преобразуем сообщение в JSON
	jsonData, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("ошибка при преобразовании в JSON: %v", err)
	}

	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("ошибка при отправке запроса: %v", err)
	}
	defer resp.Body.Close()
	fmt.Println("Сообщение успешно отправлено!")
	return nil
}