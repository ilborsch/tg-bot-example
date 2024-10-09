package schemas

import "time"

type ChatBot struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Instructions string `json:"instructions"`
}

type SaveChatBotRequest struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	Instructions string `json:"instructions"`
}

type SaveChatBotResponse struct {
	ChatBotID int64 `json:"chat_bot_id"`
}

type UpdateChatBotRequest struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	Instructions string `json:"instructions"`
}

type UpdateChatBotResponse struct {
	Success bool `json:"success"`
}

type RemoveChatBotResponse struct {
	Success bool `json:"success"`
}

type SendMessageResponse struct {
	Response string `json:"response"`
}

type SendMessageRequest struct {
	Message string `json:"message"`
}

type SendMessageError struct {
	Error string `json:"error"`
}

type UserChatBots struct {
	ChatBots []ChatBot `json:"chat_bots"`
}

type ChatBotFilesRequest struct {
	ChatBotID int64 `json:"chat_bot_id"`
}

type File struct {
	ID        int64  `json:"id"`
	Filename  string `json:"filename"`
	ChatBotID int64  `json:"chat_bot_id"`
}

type ChatBotFiles struct {
	Files []File `json:"files"`
}

type SaveFileResponse struct {
	FileID int64 `json:"file_id"`
}

type RemoveFileRequest struct {
	FileID int64 `json:"file_id"`
}

type RemoveFileResponse struct {
	Success bool `json:"success"`
}
type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	ID int64 `json:"id"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}
type User struct {
	ID             int64     `json:"id"`
	Email          string    `json:"email"`
	Plan           string    `json:"plan"`
	PlanBoughtDate time.Time `json:"plan_bought_date"`
	MessagesLeft   int       `json:"messages_left"`
	BytesDataLeft  int       `json:"bytes_data_left"`
	BotsLeft       int       `json:"bots_left"`
}

type SaveUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Plan     string `json:"plan"`
}

type SaveUserResponse struct {
	UserID int64 `json:"id"`
}

type UpdatePlanRequest struct {
	Plan string `json:"plan"`
}

type UpdatePlanResponse struct {
	Success bool `json:"success"`
}
