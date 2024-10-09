package domain

// Commands that user can execute when interacting with tg bot
const (
	NoCommand          = ""
	Start              = "/start"
	Menu               = "/menu"
	CancelLatestAction = "/cancel"

	Login    = "/login"
	Register = "/register"

	GetUserInfo        = "/user"
	GetPlanInfo        = "/plan"
	UpdateUserPlan     = "/update-plan"
	GetAllUserChatBots = "/bots"

	SaveChatBot          = "/create-bot"
	GetChatBotFiles      = "/files"
	GetChatBot           = "/bot"
	RemoveChatBot        = "/remove-bot"
	UpdateChatBot        = "/update-bot"
	StartChatWithChatBot = "/chat"
	EndChatWithChatBot   = "/end"

	SaveFile        = "/upload-file"
	GetFileMetaData = "/file-metadata"
	RemoveFile      = "/remove-file"
)

// Stages of multi-action commands that needs to be performed like Login, Register, CreateChatBot etc.
const (
	NoStage = Stage(iota)

	LoginNew
	LoginEmailEntered
	LoginPasswordEntered

	RegisterNew
	RegisterEmailEntered
	RegisterPasswordEntered

	SaveChatBotNew
	SaveChatBotNameEntered
	SaveChatBotInstructionsEntered

	GetChatBotFilesNew
	GetChatBotFilesIDEntered

	GetChatBotNew
	GetChatBotIDEntered

	RemoveChatBotNew
	RemoveChatBotIDEntered

	UpdateChatBotNew
	UpdateChatBotIDEntered
	UpdateChatBotNameEntered
	UpdateChatBotInstructionsEntered

	SaveFileNew
	SaveFileChatBotIDEntered
	SaveFileSent

	GetFileNew
	GetFileIDEntered

	RemoveFileNew
	RemoveFileIDEntered

	UpdatePlanNew
	UpdatePlanEntered
)

type Command string
type Stage int

type UserAction struct {
	Command Command
	Stage   Stage
	State   State
}

type State interface {
	State() string
}
