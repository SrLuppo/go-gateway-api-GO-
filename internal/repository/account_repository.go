package repository

import (
	"database/sql"

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

	err := r.db.QueryRow
	("SELECT id, name, email, api_key, balance, created_at, updated_at FROM accounts WHERE api_key = $1", apiKey).Scan(
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

