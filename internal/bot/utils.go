package bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"io"
	"net/http"
	"strings"
)

// extractMessageCommand extracts a command from user message
func extractMessageCommand(message *tgbotapi.Message) string {
	text := message.Text
	if strings.HasPrefix(text, "/") {
		parts := strings.Fields(text)
		if len(parts) > 0 {
			command := parts[0]
			return command
		}
	}
	return ""
}

// getDocumentBytes downloads the document and returns its binary data.
func getDocumentBytes(bot *tgbotapi.BotAPI, document *tgbotapi.Document) []byte {
	// Get the file URL
	file, err := bot.GetFile(tgbotapi.FileConfig{FileID: document.FileID})
	if err != nil {
		return nil
	}

	// Download the file
	fileURL := fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", bot.Token, file.FilePath)
	resp, err := http.Get(fileURL)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	// Read the binary data
	fileData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	return fileData
}

// isValidExtension checks if the file extension is in the list of valid extensions.
func isValidExtension(ext string) bool {
	for _, validExt := range []string{".pdf", ".json", ".csv", ".docx", ".html", ".txt"} {
		if strings.EqualFold(ext, validExt) {
			return true
		}
	}
	return false
}
