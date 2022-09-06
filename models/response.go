package models

import (
	"time"
)

type ErrorMsg struct {
	Message string `json:"message"`
}

type ResponseUserAccounts struct {
	Message      string        `json:"message"`
	UserAccounts []UserAccount `json:"data"`
}

type ResponseUserAccount struct {
	Message     string      `json:"message"`
	UserAccount UserAccount `json:"data"`
}

type UserAccount struct {
	ID         int       `json:"id"`
	Status     string    `json:"status"`
	InputDate  time.Time `json:"inputdate"`
	Nama       string    `json:"nama"`
	Keterangan string    `json:"keterangan"`
	Email      string    `json:"email"`
	Pass       string    `json:"pass"`
	UserNumber string    `json:"usernumber"`
}
