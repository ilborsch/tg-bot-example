package backend

import (
	"encoding/json"
	"fmt"
	"tg-bot/internal/schemas"
)

type userClient interface {
	UpdatePlan(token string, newPlan string) (schemas.UpdatePlanResponse, error)
	User(token string) (schemas.User, error)
}

type UserClient struct {
	baseURL string
}

func (c UserClient) UpdatePlan(token string, newPlan string) (schemas.UpdatePlanResponse, error) {
	url := fmt.Sprintf("%s/user/", c.baseURL)

	body, err := json.Marshal(map[string]string{"plan": newPlan})
	if err != nil {
		return schemas.UpdatePlanResponse{}, err
	}

	response, err := sendPutRequest(url, body, token)
	if err != nil {
		return schemas.UpdatePlanResponse{}, err
	}

	var updateResponse schemas.UpdatePlanResponse
	if err := json.Unmarshal(response, &updateResponse); err != nil {
		return schemas.UpdatePlanResponse{}, err
	}
	return updateResponse, err
}

func (c UserClient) User(token string) (schemas.User, error) {
	url := fmt.Sprintf("%s/user/", c.baseURL)

	response, err := sendGetRequest(url, token)
	if err != nil {
		return schemas.User{}, err
	}

	var user schemas.User
	if err := json.Unmarshal(response, &user); err != nil {
		return schemas.User{}, err
	}
	return user, err
}
