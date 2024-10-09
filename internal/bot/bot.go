package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"tg-bot/internal/domain"
	"tg-bot/internal/handler"
	"time"
)

type Bot struct {
	botAPI       *tgbotapi.BotAPI
	updateConfig tgbotapi.UpdateConfig
	Storage      *botStorage
	*handler.Handler
}

func New(baseURL, token string, timeout int) *Bot {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	u := tgbotapi.NewUpdate(0)
	u.Timeout = timeout

	log.Printf("Authorized on account %s", bot.Self.UserName)

	return &Bot{
		botAPI:       bot,
		updateConfig: u,
		Storage:      newBotStorage(1 * time.Hour),
		Handler:      handler.New(baseURL),
	}
}

func (b *Bot) MustRun() {
	updates, err := b.botAPI.GetUpdatesChan(b.updateConfig)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("telegram bot started!")
	for update := range updates {
		go b.HandleUpdate(update)
	}
}

func (b *Bot) HandleUpdate(update tgbotapi.Update) {
	username := update.Message.From.UserName
	log.Printf("Received a message from %s: %s", username, update.Message.Text)

	latestAction, ok := b.Storage.GetLatestAction(username)
	if !ok {
		latestAction = domain.UserAction{
			Command: domain.NoCommand,
			Stage:   domain.NoStage,
			State:   nil,
		}
		b.Storage.SetLatestAction(username, latestAction)
	}

	if update.Message != nil {
		b.HandleMessage(username, latestAction, update.Message)
	}

}

func (b *Bot) HandleMessage(username string, latestAction domain.UserAction, message *tgbotapi.Message) {
	token, _ := b.Storage.CheckToken(username)
	//if !isAuthenticated {
	//	action, response := b.Handler.AuthHandler.LoginNew()
	//	b.Storage.SetLatestAction(username, action)
	//
	//	msg := tgbotapi.NewMessage(message.Chat.ID, response)
	//	b.botAPI.Send(msg)
	//	return
	//}

	documentData := make([]byte, 0)
	documentFileName := ""
	if message.Document != nil {
		documentFileName = message.Document.FileName
		documentData = getDocumentBytes(b.botAPI, message.Document)
	}

	newUserAction := domain.UserAction{}
	botResponse := ""

	command := extractMessageCommand(message)
	if command == "" {
		switch latestAction.Stage {
		case domain.LoginNew:
			newUserAction, botResponse = b.Handler.AuthHandler.LoginEmailEntered(message.Text)
		case domain.LoginEmailEntered:
			state := latestAction.State.(domain.LoginState)
			newUserAction, token, botResponse = b.Handler.AuthHandler.LoginPasswordEntered(state.Email, message.Text)
			b.Storage.SetToken(username, token)

		case domain.RegisterNew:
			newUserAction, botResponse = b.Handler.AuthHandler.RegisterEmailEntered(message.Text)
		case domain.RegisterEmailEntered:
			state := latestAction.State.(domain.RegisterState)
			newUserAction, botResponse = b.Handler.AuthHandler.RegisterPasswordEntered(state.Email, message.Text)

		case domain.SaveChatBotNew:
			newUserAction, botResponse = b.Handler.ChatBotHandler.SaveNameEntered(message.Text)
		case domain.SaveChatBotNameEntered:
			state := latestAction.State.(domain.SaveChatBotState)
			newUserAction, botResponse = b.Handler.ChatBotHandler.SaveInstructionsEntered(token, state.Name, message.Text)

		case domain.GetChatBotFilesNew:
			newUserAction, botResponse = b.Handler.ChatBotHandler.ChatBotFilesIDEntered(token, message.Text)

		case domain.GetChatBotNew:
			newUserAction, botResponse = b.Handler.ChatBotHandler.ChatBotIDEntered(token, message.Text)

		case domain.RemoveChatBotNew:
			newUserAction, botResponse = b.Handler.ChatBotHandler.RemoveIDEntered(token, message.Text)

		case domain.UpdateChatBotNew:
			newUserAction, botResponse = b.Handler.ChatBotHandler.UpdateIDEntered(message.Text)
		case domain.UpdateChatBotIDEntered:
			state := latestAction.State.(domain.UpdateChatBotState)
			newUserAction, botResponse = b.Handler.ChatBotHandler.UpdateNameEntered(state.ID, message.Text)
		case domain.UpdateChatBotNameEntered:
			state := latestAction.State.(domain.UpdateChatBotState)
			newUserAction, botResponse = b.Handler.ChatBotHandler.UpdateInstructionsEntered(token, state.ID, state.Name, message.Text)

		case domain.SaveFileNew:
			newUserAction, botResponse = b.Handler.FileHandler.SaveIDEntered(message.Text)
		case domain.SaveFileChatBotIDEntered:
			state := latestAction.State.(domain.SaveFileState)
			newUserAction, botResponse = b.Handler.FileHandler.SaveFileSent(token, state.ChatBotID, documentFileName, documentData)

		case domain.GetFileNew:
			newUserAction, botResponse = b.Handler.FileHandler.FileIDEntered(token, message.Text)

		case domain.RemoveFileNew:
			newUserAction, botResponse = b.Handler.FileHandler.RemoveIDEntered(token, message.Text)

			//case domain.UpdatePlanNew:
			//	newUserAction, botResponse = b.Handler.UserHandler.UpdatePlanEntered(token, message.Text)
		default:
			log.Println("Unknown stage")

		}
	} else {
		switch command {
		case domain.Start:
			newUserAction, botResponse = b.Handler.GeneralHandler.Start()
		case domain.Menu:
			newUserAction, botResponse = b.Handler.GeneralHandler.Menu()
		case domain.CancelLatestAction:
			newUserAction, botResponse = b.Handler.GeneralHandler.Cancel()
		case domain.GetPlanInfo:
			newUserAction, botResponse = b.Handler.GeneralHandler.Plans()

		case domain.Login:
			newUserAction, botResponse = b.Handler.AuthHandler.LoginNew()
		case domain.Register:
			newUserAction, botResponse = b.Handler.AuthHandler.RegisterNew()

		case domain.GetUserInfo:
			newUserAction, botResponse = b.Handler.UserHandler.User(token)
		//case domain.UpdateUserPlan:
		//	newUserAction, botResponse = b.Handler.UserHandler.UpdatePlanNew()
		case domain.GetAllUserChatBots:
			newUserAction, botResponse = b.Handler.UserHandler.GetUserChatBots(token)

		case domain.SaveChatBot:
			newUserAction, botResponse = b.Handler.ChatBotHandler.SaveNew()
		case domain.GetChatBotFiles:
			newUserAction, botResponse = b.Handler.ChatBotHandler.ChatBotFilesNew()
		case domain.GetChatBot:
			newUserAction, botResponse = b.Handler.ChatBotHandler.ChatBotNew()
		case domain.RemoveChatBot:
			newUserAction, botResponse = b.Handler.ChatBotHandler.RemoveNew()
		case domain.UpdateChatBot:
			newUserAction, botResponse = b.Handler.ChatBotHandler.UpdateNew()

		case domain.SaveFile:
			newUserAction, botResponse = b.Handler.FileHandler.SaveNew()
		case domain.GetFileMetaData:
			newUserAction, botResponse = b.Handler.FileHandler.FileNew()
		case domain.RemoveFile:
			newUserAction, botResponse = b.Handler.FileHandler.RemoveNew()

		default:
			newUserAction, botResponse = b.Handler.GeneralHandler.UnknownCommand()
		}
	}

	b.Storage.SetLatestAction(username, newUserAction)

	msg := tgbotapi.NewMessage(message.Chat.ID, botResponse)
	b.botAPI.Send(msg)
}
