package handler

import (
	"fmt"
	"strconv"
	"tg-bot/internal/backend"
	"tg-bot/internal/domain"
)

type FileHandler interface {
	FileSaver
	FileProvider
	FileRemover
}

type FileSaver interface {
	SaveNew() (domain.UserAction, string)
	SaveIDEntered(id string) (domain.UserAction, string)
	SaveFileSent(token string, id int64, filename string, fileData []byte) (domain.UserAction, string)
}

type FileProvider interface {
	FileNew() (domain.UserAction, string)
	FileIDEntered(token, id string) (domain.UserAction, string)
}

type FileRemover interface {
	RemoveNew() (domain.UserAction, string)
	RemoveIDEntered(token, id string) (domain.UserAction, string)
}

type fileHandler struct {
	BackendClient *backend.Client
}

func (h fileHandler) SaveNew() (domain.UserAction, string) {
	return domain.UserAction{
		Command: domain.SaveFile,
		Stage:   domain.SaveFileNew,
	}, "For which chat-bot would you like to attach a file? Please, send an ID. (1/2)"
}

func (h fileHandler) SaveIDEntered(id string) (domain.UserAction, string) {
	chatBotID, err := strconv.Atoi(id)
	if err != nil {
		return domain.UserAction{
			Command: domain.SaveFile,
			Stage:   domain.SaveFileNew,
		}, "Unknown error occurred, please, try again by entering chat-bot ID. (1/2)"
	}
	return domain.UserAction{
		Command: domain.SaveFile,
		Stage:   domain.SaveFileChatBotIDEntered,
		State: domain.SaveFileState{
			ChatBotID: int64(chatBotID),
		},
	}, "Now, send me the file, please."
}

func (h fileHandler) SaveFileSent(token string, id int64, filename string, fileData []byte) (domain.UserAction, string) {

	_, err := h.BackendClient.SaveFile(token, id, filename, fileData)
	if err != nil {
		return domain.UserAction{
			Command: domain.SaveFile,
			Stage:   domain.SaveFileChatBotIDEntered,
			State: domain.SaveFileState{
				ChatBotID: id,
			},
		}, "Unknown error occurred, please, try again by sending another file."
	}

	return domain.UserAction{
		Command: domain.SaveFile,
		Stage:   domain.SaveFileSent,
		State: domain.SaveFileState{
			ChatBotID: id,
			Filename:  filename,
		},
	}, fmt.Sprintf("The file %s has been successfully uploaded. Now you can test the chat-bot with /chat", filename)
}

func (h fileHandler) FileNew() (domain.UserAction, string) {
	return domain.UserAction{
		Command: domain.GetFileMetaData,
		Stage:   domain.GetFileNew,
	}, "Please, enter the file ID. (1/1)"
}

func (h fileHandler) FileIDEntered(token, id string) (domain.UserAction, string) {
	chatBotID, err := strconv.Atoi(id)
	if err != nil {
		return domain.UserAction{
			Command: domain.GetFileMetaData,
			Stage:   domain.GetFileNew,
		}, "Invalid ID format, please try again."
	}

	file, err := h.BackendClient.FileMetadata(token, int64(chatBotID))
	if err != nil {
		return domain.UserAction{
			Command: domain.GetFileMetaData,
			Stage:   domain.GetFileNew,
		}, "Invalid ID, please try again."
	}

	return domain.UserAction{
		Command: domain.GetFileMetaData,
		Stage:   domain.GetFileIDEntered,
	}, fmt.Sprintf("ID: %v, filename: \"%s\", chat-bot ID: %v", file.ID, file.Filename, file.ChatBotID)
}

func (h fileHandler) RemoveNew() (domain.UserAction, string) {
	return domain.UserAction{
		Command: domain.RemoveFile,
		Stage:   domain.RemoveFileNew,
	}, "Please, enter the file ID."
}

func (h fileHandler) RemoveIDEntered(token, id string) (domain.UserAction, string) {
	chatBotID, err := strconv.Atoi(id)
	if err != nil {
		return domain.UserAction{
			Command: domain.RemoveFile,
			Stage:   domain.RemoveFileNew,
		}, "Invalid ID format, please try again."
	}

	_, err = h.BackendClient.RemoveFile(token, int64(chatBotID))
	if err != nil {
		return domain.UserAction{
			Command: domain.RemoveFile,
			Stage:   domain.RemoveFileNew,
		}, "Invalid ID, please try again."
	}

	return domain.UserAction{
		Command: domain.RemoveFile,
		Stage:   domain.RemoveFileIDEntered,
	}, fmt.Sprintf("File removed successfully!")
}
