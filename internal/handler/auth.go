package handler

import (
	"tg-bot/internal/backend"
	"tg-bot/internal/domain"
	"tg-bot/internal/schemas"
)

type AuthHandler interface {
	LoginHandler
	RegisterHandler
}

type LoginHandler interface {
	LoginNew() (domain.UserAction, string)
	LoginEmailEntered(email string) (domain.UserAction, string)
	LoginPasswordEntered(email, password string) (action domain.UserAction, token, response string)
}

type RegisterHandler interface {
	RegisterNew() (domain.UserAction, string)
	RegisterEmailEntered(email string) (domain.UserAction, string)
	RegisterPasswordEntered(email, password string) (domain.UserAction, string)
}

type authHandler struct {
	BackendClient *backend.Client
}

func (h authHandler) LoginNew() (domain.UserAction, string) {
	return domain.UserAction{
		Command: domain.Login,
		Stage:   domain.LoginNew,
	}, "If you would like to login, please, send me your email."
}

func (h authHandler) LoginEmailEntered(email string) (domain.UserAction, string) {
	return domain.UserAction{
		Command: domain.Login,
		Stage:   domain.LoginEmailEntered,
		State: domain.LoginState{
			Email: email,
		},
	}, "Please, send me your password in order to login."
}

func (h authHandler) LoginPasswordEntered(email, password string) (action domain.UserAction, token, response string) {
	request := schemas.LoginRequest{
		Email:    email,
		Password: password,
	}
	token, err := h.BackendClient.Login(request)
	if err != nil {
		return domain.UserAction{
			Command: domain.Login,
			Stage:   domain.LoginNew,
		}, "", "Unknown error happened, please, enter your email again."
	}
	return domain.UserAction{
		Command: domain.Login,
		Stage:   domain.LoginPasswordEntered,
		State: domain.LoginState{
			Email:    email,
			Password: password,
		},
	}, token, "Your login is successful!"
}

func (h authHandler) RegisterNew() (domain.UserAction, string) {
	return domain.UserAction{
		Command: domain.Register,
		Stage:   domain.RegisterNew,
	}, "If you would like to register, please, send me your email."
}

func (h authHandler) RegisterEmailEntered(email string) (domain.UserAction, string) {
	return domain.UserAction{
		Command: domain.Register,
		Stage:   domain.RegisterEmailEntered,
		State: domain.RegisterState{
			Email: email,
		},
	}, "Please, send me your password in order to register."
}

func (h authHandler) RegisterPasswordEntered(email, password string) (action domain.UserAction, response string) {
	request := schemas.RegisterRequest{
		Email:    email,
		Password: password,
	}
	_, err := h.BackendClient.Register(request)
	if err != nil {
		return domain.UserAction{
			Command: domain.Register,
			Stage:   domain.RegisterNew,
		}, "Unknown error happened, please, enter your email again."
	}
	return domain.UserAction{
		Command: domain.Register,
		Stage:   domain.RegisterPasswordEntered,
		State: domain.RegisterState{
			Email:    email,
			Password: password,
		},
	}, "Your registration is successful! Now you can login!"
}
