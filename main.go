package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type Password struct {
	Name         string    `json:"name"`
	Value        string    `json:"value"`    // password value
	Category     string    `json:"category"` // social, finance
	CreatedAt    time.Time `json:"CreatedAt"`
	LastModified time.Time `json:"lastModified"`
}

func NewPassword(name, value, category string) Password {
	now := time.Now().UTC()
	return Password{
		Name:         name,
		Value:        value,
		Category:     category,
		CreatedAt:    now,
		LastModified: now,
	}
}

type PasswordManager struct {
	passwords   map[string]Password
	masterKey   []byte
	filepath    string
	isInitilzed bool `json:"-"`
}

func NewPasswordManager(filepath string) *PasswordManager {
	return &PasswordManager{
		passwords:   make(map[string]Password),
		masterKey:   make([]byte, 0),
		filepath:    filepath,
		isInitilzed: false,
	}
}

func main() {
	password := NewPassword("github.com", "password123", "development")

	out, err := json.Marshal(password)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(out))

	passwordManager := NewPasswordManager("passwords.dat")
	fmt.Printf("isInitilzed: %t,\nFile path: %s,\n", passwordManager.isInitilzed, passwordManager.filepath)
}
