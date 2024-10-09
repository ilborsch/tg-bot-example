package backend

import (
	"encoding/json"
	"fmt"
	"tg-bot/internal/schemas"
)

type fileClient interface {
	ChatBotFiles(token string, chatBotID int64) (schemas.ChatBotFiles, error)
	FileMetadata(token string, fileID int64) (schemas.File, error)
	SaveFile(token string, chatbotID int64, filename string, fileData []byte) (schemas.SaveFileResponse, error)
	RemoveFile(token string, fileID int64) (schemas.RemoveFileResponse, error)
}

type FileClient struct {
	baseURL string
}

func (c FileClient) ChatBotFiles(token string, chatBotID int64) (schemas.ChatBotFiles, error) {
	url := fmt.Sprintf("%s/chat-bot/%v/files", c.baseURL, chatBotID)

	response, err := sendGetRequest(url, token)
	if err != nil {
		return schemas.ChatBotFiles{}, err
	}

	var chatBotFiles schemas.ChatBotFiles
	if err := json.Unmarshal(response, &chatBotFiles); err != nil {
		return schemas.ChatBotFiles{}, err
	}
	return chatBotFiles, err
}

func (c FileClient) FileMetadata(token string, fileID int64) (schemas.File, error) {
	url := fmt.Sprintf("%s/file/%v", c.baseURL, fileID)

	response, err := sendGetRequest(url, token)
	if err != nil {
		return schemas.File{}, err
	}

	var file schemas.File
	if err := json.Unmarshal(response, &file); err != nil {
		return schemas.File{}, err
	}
	return file, err
}

func (c FileClient) SaveFile(token string, chatbotID int64, filename string, fileData []byte) (schemas.SaveFileResponse, error) {
	url := fmt.Sprintf("%s/file/", c.baseURL)

	body, err := json.Marshal(map[string]interface{}{"chat_bot_id": chatbotID})
	if err != nil {
		return schemas.SaveFileResponse{}, err
	}

	response, err := sendPostWithFile(url, body, token, filename, fileData)
	if err != nil {
		return schemas.SaveFileResponse{}, err
	}

	var saveResponse schemas.SaveFileResponse
	if err := json.Unmarshal(response, &saveResponse); err != nil {
		return schemas.SaveFileResponse{}, err
	}
	return saveResponse, err
}

func (c FileClient) RemoveFile(token string, fileID int64) (schemas.RemoveFileResponse, error) {
	url := fmt.Sprintf("%s/file/%v", c.baseURL, fileID)

	response, err := sendDeleteRequest(url, token)
	if err != nil {
		return schemas.RemoveFileResponse{}, err
	}

	var removeResponse schemas.RemoveFileResponse
	if err := json.Unmarshal(response, &removeResponse); err != nil {
		return schemas.RemoveFileResponse{}, err
	}
	return removeResponse, err
}
