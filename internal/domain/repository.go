package domain

type AccountRepository interface {
	Save(account *Account) error
	FindByApiKey(apiKey string) (*Account, error)
	FindById(id string) (*Account, error)
	Update(account *Account) error
}
