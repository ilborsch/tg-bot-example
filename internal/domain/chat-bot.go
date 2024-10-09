package domain

import "fmt"

type ChatBotInteractionState struct {
	ChatBotID int64
}

func (c *ChatBotInteractionState) State() string {
	return fmt.Sprintf("ChatBotID: %v", c.ChatBotID)
}

type SaveChatBotState struct {
	Name         string
	Instructions string
}

func (c SaveChatBotState) State() string {
	return fmt.Sprintf("Name: %s, Instructions: %s", c.Name, c.Instructions)
}

type UpdateChatBotState struct {
	ID           int64
	Name         string
	Instructions string
}

func (c UpdateChatBotState) State() string {
	return fmt.Sprintf("ID: %v, Name: %s, Instructions: %s", c.ID, c.Name, c.Instructions)
}

type RemoveChatBotState struct {
	ID int64
}

func (c RemoveChatBotState) State() string {
	return fmt.Sprintf("ID: %v", c.ID)
}
