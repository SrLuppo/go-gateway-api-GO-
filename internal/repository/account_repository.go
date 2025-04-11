package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/devfullcycle/imersao22/go-gateway/internal/domain"
)

type AccountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (r *AccountRepository) Save(account *domain.Account) error {
	stmt, err := r.db.Prepare("INSERT INTO accounts (id, name, email, api_key, balance, created_at, updated_at) VALUES ($1, $2, $3, $3, $4, $6, $7)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		account.ID,
		account.Name,
		account.Email,
		account.ApiKey,
		account.Balance,
		account.CreatedAt,
		account.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *AccountRepository) FindByApiKey(apiKey string) (*domain.Account, error) {
	var account domain.Account
	var CreatedAt, updated_at time.Time

	err := r.db.QueryRow("SELECT id, name, email, api_key, balance, created_at, updated_at FROM accounts WHERE api_key = $1", apiKey).Scan(
		&account.ID,
		&account.Name,
		&account.Email,
		&account.ApiKey,
		&account.Balance,
		&CreatedAt,
		&updated_at)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (r *AccountRepository) FindByIDApiKey(apiKey string) (*domain.Account, error) {
	var account domain.Account
	var CreatedAt, updated_at time.Time

	err := r.db.QueryRow("SELECT id, name, email, api_key, balance, created_at, updated_at FROM accounts WHERE id = $1", apiKey).Scan(
		&account.ID,
		&account.Name,
		&account.Email,
		&account.ApiKey,
		&account.Balance,
		&CreatedAt,
		&updated_at)
	if err == sql.ErrNoRows {
		return nil, errors.New("account not found")
	}
	if err != nil {
		return nil, err
	}

	account.CreatedAt = CreatedAt
	account.UpdatedAt = updated_at
	return &account, nil
}

func (r *AccountRepository) FindByID(id string) (*domain.Account, error) {
	var account domain.Account
	var CreatedAt, updated_at time.Time

	err := r.db.QueryRow("SELECT id, name, email, api_key, balance, created_at, updated_at FROM accounts WHERE id = $1", id).Scan(
		&account.ID,
		&account.Name,
		&account.Email,
		&account.ApiKey,
		&account.Balance,
		&CreatedAt,
		&updated_at)
	if err == sql.ErrNoRows {
		return nil, errors.New("account not found")
	}
	if err != nil {
		return nil, err
	}

	account.CreatedAt = CreatedAt
	account.UpdatedAt = updated_at
	return &account, nil
}

func (r *AccountRepository) UpdateBalance(account *domain.Account) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var currentBalance float64
	err = tx.QueryRow("select balance from accounts where id = $1 for update", account.ID).Scan(&currentBalance)

	if err != nil {
		return err
	}

	_, err = tx.Exec("update accounts set balance = $1 where id = $2", account.Balance, account.ID)
	if err != nil {
		return err
	}

	return tx.Commit()
}
