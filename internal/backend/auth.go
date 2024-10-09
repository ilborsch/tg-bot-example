package backend

import (
	"encoding/json"
	"fmt"
	"tg-bot/internal/schemas"
)

type authClient interface {
	Login(request schemas.LoginRequest) (token string, err error)
	Register(request schemas.RegisterRequest) (schemas.RegisterResponse, error)
}

type AuthClient struct {
	baseURL string
}

func (c AuthClient) Login(request schemas.LoginRequest) (token string, err error) {
	url := fmt.Sprintf("%s/login", c.baseURL)

	body, err := json.Marshal(request)
	if err != nil {
		return "", err
	}

	response, err := sendPostRequest(url, body, "")
	if err != nil {
		return "", err
	}

	var loginResponse schemas.LoginResponse
	if err := json.Unmarshal(response, &loginResponse); err != nil {
		return "", err
	}
	return loginResponse.Token, err
}

func (c AuthClient) Register(request schemas.RegisterRequest) (schemas.RegisterResponse, error) {
	url := fmt.Sprintf("%s/login", c.baseURL)

	body, err := json.Marshal(request)
	if err != nil {
		return schemas.RegisterResponse{}, err
	}

	response, err := sendPostRequest(url, body, "")
	if err != nil {
		return schemas.RegisterResponse{}, err
	}

	var registerResponse schemas.RegisterResponse
	if err := json.Unmarshal(response, &registerResponse); err != nil {
		return schemas.RegisterResponse{}, err
	}
	return registerResponse, err
}
