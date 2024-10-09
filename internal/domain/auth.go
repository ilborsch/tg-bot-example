package domain

import "fmt"

type LoginState struct {
	Email    string
	Password string
}

func (c LoginState) State() string {
	return fmt.Sprintf("Email: %s, Password: %s", c.Email, c.Password)
}

type RegisterState struct {
	Email    string
	Password string
}

func (c RegisterState) State() string {
	return fmt.Sprintf("Email: %s, Password: %s", c.Email, c.Password)
}
