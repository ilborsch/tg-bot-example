package handler

import (
	"errors"
	"tg-bot/internal/backend"
)

var CannotPerformThisAction = errors.New("cannot perform this action")

type Handler struct {
	AuthHandler
	ChatBotHandler
	FileHandler
	GeneralHandler
	UserHandler
}

func New(baseURL string) *Handler {
	return &Handler{
		AuthHandler: authHandler{
			BackendClient: backend.NewClient(baseURL),
		},
		ChatBotHandler: chatBotHandler{
			BackendClient: backend.NewClient(baseURL),
		},
		FileHandler: fileHandler{
			BackendClient: backend.NewClient(baseURL),
		},
		GeneralHandler: generalHandler{
			BackendClient: backend.NewClient(baseURL),
		},
		UserHandler: userHandler{
			BackendClient: backend.NewClient(baseURL),
		},
	}
}
