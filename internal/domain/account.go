package domain

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
)

var ErrDuplicateAccount = errors.New("account already exists")

type Account struct {
	ID        string
	Name      string
	Email     string
	ApiKey    string
	Balance   float64
	mu        sync.Mutex
	CreatedAt time.Time
	UpdatedAt time.Time
}

func generateApiKey() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

func NewAccount(name, email string) *Account {
	account := &Account{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		ApiKey:    generateApiKey(),
		Balance:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return account
}
func (a *Account) AddBalance(amount float64) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Balance += amount
	a.UpdatedAt = time.Now()
	a.mu.Unlock()
}
