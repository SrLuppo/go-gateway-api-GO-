package domain

type AccountRepository interface {
	Save(account *Account) error
	FindByApiKey(apiKey string) (*Account, error)
	FindByID(id string) (*Account, error)
	Update(account *Account) error
}

type InvoiceRepository interface {
	Save(invoice *Invoice) error
	FindByID(id string) (*Invoice, error)
	FindByAccountID(accountID string) ([]*Invoice, error)
	Update(invoice *Invoice) error
}
