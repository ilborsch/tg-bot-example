package handler

import (
	"fmt"
	"tg-bot/internal/backend"
	"tg-bot/internal/domain"
)

type UserHandler interface {
	User(token string) (domain.UserAction, string)
	GetUserChatBots(token string) (domain.UserAction, string)
	//UpdatePlanNew() (domain.UserAction, string)
	//UpdatePlanEntered(token string, newPlan string) (domain.UserAction, string)
}

type userHandler struct {
	BackendClient *backend.Client
}

func (h userHandler) User(token string) (domain.UserAction, string) {
	user, err := h.BackendClient.User(token)
	if err != nil {
		return domain.UserAction{
			Command: domain.GetUserInfo,
		}, "You are not logged in, please, login and try again."
	}

	plan := ""
	if user.Plan == "free_plan" {
		plan = "Free"
	}
	if user.Plan == "business_plan" {
		plan = "Business"
	}
	if user.Plan == "enterprise_plan" {
		plan = "Enterprise"
	}
	dataLeftMB := float32(user.BytesDataLeft / 1024 / 1024)

	return domain.UserAction{
			Command: domain.GetUserInfo,
		}, fmt.Sprintf(`
		Your email is: %s
		Plan: %s
		Chat-bots left: %v
		Messages in this month left: %v
		Data left: %v
`, user.Email, plan, user.BotsLeft, user.MessagesLeft, dataLeftMB)
}

func (h userHandler) GetUserChatBots(token string) (domain.UserAction, string) {
	chatBots, err := h.BackendClient.ChatBots(token)
	if err != nil {
		return domain.UserAction{
			Command: domain.GetUserInfo,
		}, "You are not logged in, please, login and try again."
	}

	msg := "Here are your chat-bots: "
	for i, cb := range chatBots.ChatBots {
		msg += fmt.Sprintf("\n%v) ID: %v, Name: %s, Instructions: %s", i+1, cb.ID, cb.Name, cb.Instructions)
	}

	return domain.UserAction{
		Command: domain.GetAllUserChatBots,
	}, msg
}
