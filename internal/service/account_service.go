package service

import (
	"database/sql"

	"github.com/devfullcycle/imersao22/go-gateway/internal/domain"
	"github.com/devfullcycle/imersao22/go-gateway/internal/dto"
	"github.com/devfullcycle/imersao22/go-gateway/internal/repository"
)

type AccountService struct {
	repository *repository.AccountRepository
}

func NewAccountService(repository *repository.AccountRepository) *AccountService {
	return &AccountService{repository: repository}
}

func (s *AccountService) CreateAccount(input *dto.CreateAccountRequest) (*dto.AccountOutput, error) {
	account := dto.ToAccount(input)

	existingAccount, err := s.repository.FindByApiKey(account.ApiKey)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if existingAccount != nil {
		return nil, domain.ErrDuplicateAccount
	}

	err = s.repository.Save(account)
	if err != nil {
		return nil, err
	}
	return dto.FromAccount(account), nil
}

func (s *AccountService) UpdateBalance(apiKey string, amount float64) (*dto.AccountOutput, error) {
	account, err := s.repository.FindByApiKey(apiKey)
	if err != nil {
		return nil, err
	}

	account.Balance += amount

	err = s.repository.UpdateBalance(account)
	if err != nil {
		return nil, err
	}

	return dto.FromAccount(account), nil
}

func (s *AccountService) FindAccountByApiKey(apiKey string) (*dto.AccountOutput, error) {
	account, err := s.repository.FindByApiKey(apiKey)
	if err != nil {
		return nil, err
	}
	return dto.FromAccount(account), nil
}

func (s *AccountService) FindAccountById(id string) (*dto.AccountOutput, error) {
	account, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return dto.FromAccount(account), nil
}
