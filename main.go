package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)
// структура для хранения одной записи пароля
type Password struct {
	Name         string    `json:"name"`
	Value        string    `json:"value"`    // password value
	Category     string    `json:"category"` // social, finance
	CreatedAt    time.Time `json:"createdAt"`
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

//ядро енеджера
type PasswordManager struct {
	passwords   map[string]Password
	masterKey   []byte
	filepath    string
	isInitialized bool `json:"-"`
}

// конструктор
func NewPasswordManager(filepath string) *PasswordManager {
	return &PasswordManager{
		passwords:   make(map[string]Password),
		masterKey:   make([]byte, 0),
		filepath:    filepath,
		isInitialized: false,
	}
}

//метод
func(pm *PasswordManager)SetMasterPassword(masterPassword string) error{
	if len(masterPassword) < 8{
		return fmt.Errorf("password is too weak")
	}
	keyBuffer := make([]byte, 32)
	copy(keyBuffer, []byte(masterPassword))
	pm.masterKey = keyBuffer
	pm.isInitialized = true
	return nil
}

//добавление новых паролей в хранилище 
func (pm *PasswordManager)SavePassword(name, value, category string) error{
	if !pm.isInitialized{
		return fmt.Errorf("password manager is not initialized")
	}

	if _, ok := pm.passwords[name]; ok{
		return fmt.Errorf("password already exists")
	}

	password := NewPassword(name, value, category)
	pm.passwords[name] = password
	return nil
}

//получение пароля
func (pm *PasswordManager)GetPassword(name string) (Password, error){
	if !pm.isInitialized{
		return Password{}, fmt.Errorf("password manager does noi initialized")
	}

	if password, ok := pm.passwords[name];ok{
		return NewPassword(name, password.Value, password.Category), nil
	}else{
		return Password{}, fmt.Errorf("password not found")
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
	fmt.Printf("isInitialized: %t,\nFile path: %s,\n", passwordManager.isInitialized, passwordManager.filepath)
}
