package handler

import (
	"fmt"
	"strconv"
	"tg-bot/internal/backend"
	"tg-bot/internal/domain"
	"tg-bot/internal/schemas"
)

type ChatBotHandler interface {
	ChatBotProvider
	ChatBotSaver
	ChatBotFilesProvider
	ChatBotRemover
	ChatBotUpdater
	// StartChat()
	// EndChat()
}

type ChatBotSaver interface {
	SaveNew() (domain.UserAction, string)
	SaveNameEntered(name string) (domain.UserAction, string)
	SaveInstructionsEntered(token, name, instructions string) (domain.UserAction, string)
}

type ChatBotFilesProvider interface {
	ChatBotFilesNew() (domain.UserAction, string)
	ChatBotFilesIDEntered(token, id string) (domain.UserAction, string)
}

type ChatBotProvider interface {
	ChatBotNew() (domain.UserAction, string)
	ChatBotIDEntered(token, id string) (domain.UserAction, string)
}

type ChatBotRemover interface {
	RemoveNew() (domain.UserAction, string)
	RemoveIDEntered(token, id string) (domain.UserAction, string)
}

type ChatBotUpdater interface {
	UpdateNew() (domain.UserAction, string)
	UpdateIDEntered(id string) (domain.UserAction, string)
	UpdateNameEntered(id int64, name string) (domain.UserAction, string)
	UpdateInstructionsEntered(token string, id int64, name, instructions string) (domain.UserAction, string)
}

type chatBotHandler struct {
	BackendClient *backend.Client
}

func (h chatBotHandler) SaveNew() (domain.UserAction, string) {
	return domain.UserAction{
		Command: domain.SaveChatBot,
		Stage:   domain.SaveChatBotNew,
	}, "Please, enter chat-bot name..."
}

func (h chatBotHandler) SaveNameEntered(name string) (domain.UserAction, string) {
	return domain.UserAction{
		Command: domain.SaveChatBot,
		Stage:   domain.SaveChatBotNameEntered,
		State: domain.SaveChatBotState{
			Name: name,
		},
	}, "Please, enter chat-bot instructions. (e.g. \"You are a customer support in a pizzeria. Be funny and friendly.\")"
}

func (h chatBotHandler) SaveInstructionsEntered(token, name, instructions string) (domain.UserAction, string) {
	request := schemas.SaveChatBotRequest{
		Name:         name,
		Instructions: instructions,
	}

	response, err := h.BackendClient.Save(token, request)
	if err != nil {
		return domain.UserAction{
			Command: domain.SaveChatBot,
			Stage:   domain.SaveChatBotNew,
		}, "Unknown error occurred, please, try again by entering chat-bot name."
	}

	return domain.UserAction{
		Command: domain.SaveChatBot,
		Stage:   domain.SaveChatBotNameEntered,
		State: domain.SaveChatBotState{
			Name: name,
		},
	}, fmt.Sprintf("Your chat-bot has been successfully created with ID %v", response.ChatBotID)
}

func (h chatBotHandler) ChatBotFilesNew() (domain.UserAction, string) {
	return domain.UserAction{
		Command: domain.GetChatBotFiles,
		Stage:   domain.GetChatBotFilesNew,
	}, "Please, enter the chat-bot ID."
}

func (h chatBotHandler) ChatBotFilesIDEntered(token, id string) (domain.UserAction, string) {
	chatBotID, err := strconv.Atoi(id)
	if err != nil {
		return domain.UserAction{
			Command: domain.SaveChatBot,
			Stage:   domain.SaveChatBotNew,
		}, "Invalid ID format, please try again."
	}

	files, err := h.BackendClient.ChatBotFiles(token, int64(chatBotID))
	if err != nil {
		return domain.UserAction{
			Command: domain.SaveChatBot,
			Stage:   domain.SaveChatBotNew,
		}, "Invalid ID, please try again."
	}

	msg := "Here are the chat-bot files: "
	for i, file := range files.Files {
		msg += fmt.Sprintf("\n %v) ID: %v, Filename: %s, Chat-bot ID: %v", i+1, file.ID, file.Filename, file.ChatBotID)
	}

	return domain.UserAction{
		Command: domain.GetChatBotFiles,
		Stage:   domain.GetChatBotFilesIDEntered,
	}, msg
}

func (h chatBotHandler) ChatBotNew() (domain.UserAction, string) {
	return domain.UserAction{
		Command: domain.GetChatBot,
		Stage:   domain.GetChatBotNew,
	}, "Please, enter the chat-bot ID."
}

func (h chatBotHandler) ChatBotIDEntered(token, id string) (domain.UserAction, string) {
	chatBotID, err := strconv.Atoi(id)
	if err != nil {
		return domain.UserAction{
			Command: domain.SaveChatBot,
			Stage:   domain.SaveChatBotNew,
		}, "Invalid ID format, please try again."
	}

	chatBot, err := h.BackendClient.ChatBot(token, int64(chatBotID))
	if err != nil {
		return domain.UserAction{
			Command: domain.SaveChatBot,
			Stage:   domain.SaveChatBotNew,
		}, "Invalid ID, please try again."
	}

	return domain.UserAction{
		Command: domain.GetChatBotFiles,
		Stage:   domain.GetChatBotFilesIDEntered,
	}, fmt.Sprintf("Here is your chat-bot: \nName: \"%s\" \nInstructions: \"%s\"", chatBot.Name, chatBot.Instructions)
}

func (h chatBotHandler) RemoveNew() (domain.UserAction, string) {
	return domain.UserAction{
		Command: domain.RemoveChatBot,
		Stage:   domain.RemoveChatBotNew,
	}, "Please, enter the chat-bot ID."
}

func (h chatBotHandler) RemoveIDEntered(token, id string) (domain.UserAction, string) {
	chatBotID, err := strconv.Atoi(id)
	if err != nil {
		return domain.UserAction{
			Command: domain.RemoveChatBot,
			Stage:   domain.RemoveChatBotNew,
		}, "Invalid ID format, please try again."
	}

	_, err = h.BackendClient.RemoveChatBot(token, int64(chatBotID))
	if err != nil {
		return domain.UserAction{
			Command: domain.RemoveChatBot,
			Stage:   domain.RemoveChatBotNew,
		}, "Invalid ID, please try again."
	}

	return domain.UserAction{
		Command: domain.RemoveChatBot,
		Stage:   domain.RemoveChatBotIDEntered,
	}, fmt.Sprintf("Chat-bot removed successfully!")
}

func (h chatBotHandler) UpdateNew() (domain.UserAction, string) {
	return domain.UserAction{
		Command: domain.UpdateChatBot,
		Stage:   domain.UpdateChatBotNew,
	}, "Please, enter the chat-bot ID. (1/3)"
}

func (h chatBotHandler) UpdateIDEntered(id string) (domain.UserAction, string) {
	chatBotID, err := strconv.Atoi(id)
	if err != nil {
		return domain.UserAction{
			Command: domain.UpdateChatBot,
			Stage:   domain.UpdateChatBotNew,
		}, fmt.Sprintf("Invalid ID format, please try again.")
	}

	return domain.UserAction{
		Command: domain.UpdateChatBot,
		Stage:   domain.UpdateChatBotIDEntered,
		State: domain.UpdateChatBotState{
			ID: int64(chatBotID),
		},
	}, fmt.Sprintf("Now, please, provide a new chat-bot name. (2/3)")
}

func (h chatBotHandler) UpdateNameEntered(id int64, name string) (domain.UserAction, string) {
	return domain.UserAction{
		Command: domain.UpdateChatBot,
		Stage:   domain.UpdateChatBotNew,
		State: domain.UpdateChatBotState{
			ID:   id,
			Name: name,
		},
	}, fmt.Sprintf("Now, please, provide new chat-bot instructions. (3/3)")
}

func (h chatBotHandler) UpdateInstructionsEntered(token string, id int64, name, instructions string) (domain.UserAction, string) {
	request := schemas.UpdateChatBotRequest{
		Name: name,
		Instructions: instructions,
	}
	_, err := h.BackendClient.Update(token, id, request)
	if err != nil {
		return domain.UserAction{
			Command: domain.UpdateChatBot,
			Stage:   domain.UpdateChatBotNew,
		}, fmt.Sprintf("Invalid ID format, please try again.")
	}
	return domain.UserAction{
		Command: domain.UpdateChatBot,
		Stage:   domain.UpdateChatBotInstructionsEntered,
		State: domain.UpdateChatBotState{
			ID:   id,
			Name: name,
			Instructions: instructions,
		},
	}, "The chat-bot successfully updated!"
}
