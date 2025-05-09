package domain

import (
	"errors"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	StatusPending  Status = "pending"
	StatusApproved Status = "approved"
	StatusRejected Status = "rejected"
)

var (
	ErrInvalidStatus   = errors.New("invalid status")
	ErrUnauthorized    = errors.New("unauthorized")
	ErrInvoiceNotFound = errors.New("invoice not found")
)

type Invoice struct {
	ID             string
	AccountID      string
	Amount         float64
	Status         Status
	Description    string
	PaymentType    string
	CardLastDigits string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type CreditCard struct {
	Number         string
	CVV            string
	ExpiryMonth    string
	ExpiryYear     string
	CardholderName string
}

func NewInvoice(accountID string, amount float64, description string, paymentType string, card *CreditCard) (*Invoice, error) {
	if amount <= 0 {
		return nil, errors.New("amount must be greater than 0")
	}
	lastDigits := card.Number[len(card.Number)-4:]

	invoice := &Invoice{
		ID:             uuid.New().String(),
		AccountID:      accountID,
		Amount:         amount,
		Status:         StatusPending,
		Description:    description,
		PaymentType:    paymentType,
		CardLastDigits: lastDigits,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	return invoice, nil
}

func (i *Invoice) Process() error {
	if i.Amount > 1000 {
		return nil
	}

	randomSource := rand.New(rand.NewSource(time.Now().UnixNano()))

	var newStatus Status

	if randomSource.Float64() <= 0.7 {
		newStatus = StatusApproved
	} else {
		newStatus = StatusRejected
	}

	i.Status = newStatus
	return nil
}

func (i *Invoice) UpdateStatus(status Status) error {
	if status != StatusApproved && status != StatusRejected {
		return ErrInvalidStatus
	}

	i.Status = status
	i.UpdatedAt = time.Now()
	return nil
}
