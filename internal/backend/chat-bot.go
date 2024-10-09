package backend

import (
	"encoding/json"
	"fmt"
	"tg-bot/internal/schemas"
)

type chatBotClient interface {
	ChatBot(token string, chatBotID int64) (schemas.ChatBot, error)
	ChatBots(token string) (schemas.UserChatBots, error)
	Save(token string, request schemas.SaveChatBotRequest) (schemas.SaveChatBotResponse, error)
	RemoveChatBot(token string, chatBotID int64) (schemas.RemoveChatBotResponse, error)
	Update(token string, chatBotID int64, request schemas.UpdateChatBotRequest) (schemas.UpdateChatBotResponse, error)
}

type ChatBotClient struct {
	baseURL string
}

func (c ChatBotClient) ChatBot(token string, chatBotID int64) (schemas.ChatBot, error) {
	url := fmt.Sprintf("%s/chat-bot/%v", c.baseURL, chatBotID)

	response, err := sendGetRequest(url, token)
	if err != nil {
		return schemas.ChatBot{}, err
	}

	var chatBot schemas.ChatBot
	if err := json.Unmarshal(response, &chatBot); err != nil {
		return schemas.ChatBot{}, err
	}
	return chatBot, err
}

func (c ChatBotClient) ChatBots(token string) (schemas.UserChatBots, error) {
	url := fmt.Sprintf("%s/user/chat-bots/", c.baseURL)

	response, err := sendGetRequest(url, token)
	if err != nil {
		return schemas.UserChatBots{}, err
	}

	var chatBots schemas.UserChatBots
	if err := json.Unmarshal(response, &chatBots); err != nil {
		return schemas.UserChatBots{}, err
	}
	return chatBots, err
}

func (c ChatBotClient) Save(token string, request schemas.SaveChatBotRequest) (schemas.SaveChatBotResponse, error) {
	url := fmt.Sprintf("%s/user/chat-bots/", c.baseURL)

	body, err := json.Marshal(request)
	if err != nil {
		return schemas.SaveChatBotResponse{}, err
	}

	response, err := sendPostRequest(url, body, token)
	if err != nil {
		return schemas.SaveChatBotResponse{}, err
	}

	var saveResponse schemas.SaveChatBotResponse
	if err := json.Unmarshal(response, &saveResponse); err != nil {
		return schemas.SaveChatBotResponse{}, err
	}
	return saveResponse, err
}

func (c ChatBotClient) RemoveChatBot(token string, chatBotID int64) (schemas.RemoveChatBotResponse, error) {
	url := fmt.Sprintf("%s/chat-bots/%v", c.baseURL, chatBotID)

	response, err := sendDeleteRequest(url, token)
	if err != nil {
		return schemas.RemoveChatBotResponse{}, err
	}

	var removeResponse schemas.RemoveChatBotResponse
	if err := json.Unmarshal(response, &removeResponse); err != nil {
		return schemas.RemoveChatBotResponse{}, err
	}
	return removeResponse, err
}

func (c ChatBotClient) Update(token string, chatBotID int64, request schemas.UpdateChatBotRequest) (schemas.UpdateChatBotResponse, error) {
	url := fmt.Sprintf("%s/chat-bots/%v", c.baseURL, chatBotID)

	body, err := json.Marshal(request)
	if err != nil {
		return schemas.UpdateChatBotResponse{}, err
	}

	response, err := sendPutRequest(url, body, token)
	if err != nil {
		return schemas.UpdateChatBotResponse{}, err
	}

	var updateResponse schemas.UpdateChatBotResponse
	if err := json.Unmarshal(response, &updateResponse); err != nil {
		return schemas.UpdateChatBotResponse{}, err
	}
	return updateResponse, err
}
