package handler

import (
	"tg-bot/internal/backend"
	"tg-bot/internal/domain"
)

type GeneralHandler interface {
	Menu() (domain.UserAction, string)
	Start() (domain.UserAction, string)
	Cancel() (domain.UserAction, string)
	Plans() (domain.UserAction, string)
	UnknownCommand() (domain.UserAction, string)
}

type generalHandler struct {
	BackendClient *backend.Client
}

func (h generalHandler) Menu() (domain.UserAction, string) {
	msg := `
		Hi! here is the list of all available commands:
		
		/menu - :)
		/start - Start interaction with me
		/cancel - Cancel your latest command
		
		/login - Login into your Bot Factory account
		/register - Register a new account
		
		/user - Get your user account info
		/plans - Get a list of all plans
		/plan - Get your plan info
		/update-plan - Update your current plan
		
		/bots - Get a list of all your chat-bots
		/create-bot - Create a new chat-bot
		/bot - Get chat-bot info
		/files - Get chat-bot files
		/remove-bot - Remove your chat-bot
		/update-bot - Update your chat-bot info
		
		/chat - Start chat with chat-bot
		/end - End chat with your chat-bot (same as /cancel)
		
		/upload-file - Upload a file to a chat-bot data storage
		/file-metadata - Get file meta-data
		/remove-file - Remove a file from a chat-bot data storage
	`
	return domain.UserAction{
		Command: domain.Menu,
	}, msg
}

func (h generalHandler) Start() (domain.UserAction, string) {
	msg := `
		Hi! I am the Bot Factory telegram bot, and here you can create a fully functional chat-bot for your business just in one click!
		Continue by clicking /menu to get acknowledged with commands that I provide!
	`
	return domain.UserAction{
		Command: domain.Start,
	}, msg
}

func (h generalHandler) Cancel() (domain.UserAction, string) {
	msg := `
		Latest action cancelled.
	`
	return domain.UserAction{
		Command: domain.CancelLatestAction,
	}, msg
}

func (h generalHandler) Plans() (domain.UserAction, string) {
	msg := `
		Bot Factory provides three pricing plans for our users:

		1. Free Plan: 1 chat-bot, 50 messages, 1MB data
		2. Business Plan: 3 chat-bots, 2500 messages, 100MB data
		3. Enterprise Plan: 10 chat-bots, 10000 messages, 500MB data
		
		For pricing, please, contact @borschhh
	`
	return domain.UserAction{
		Command: domain.CancelLatestAction,
	}, msg
}

func (h generalHandler) UnknownCommand() (domain.UserAction, string) {
	return domain.UserAction{
		Command: domain.NoCommand,
	}, "Unknown command. You should try one from /menu"
}
