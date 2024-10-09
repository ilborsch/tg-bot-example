package domain

import "fmt"

type SaveFileState struct {
	ChatBotID int64
	Filename  string
	Data      []byte
}

func (c SaveFileState) State() string {
	return fmt.Sprintf("ID: %v, Filename: %s", c.ChatBotID, c.Filename)
}

type RemoveFileState struct {
	FileID int64
}

func (c RemoveFileState) State() string {
	return fmt.Sprintf("ID: %v", c.FileID)
}
