package repository

import (
	"database/sql"
	"errors"

	"github.com/devfullcycle/imersao22/go-gateway/internal/domain"
)

type InvoiceRepository struct {
	db *sql.DB
}

func NewInvoiceRepository(db *sql.DB) *InvoiceRepository {
	return &InvoiceRepository{db: db}
}

// Save salva uma fatura no banco de dados
func (r *InvoiceRepository) Save(invoice *domain.Invoice) error {
	stmt, err := r.db.Prepare("INSERT INTO invoices (id, account_id, amount, status, description, payment_type, card_last_digits, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		invoice.ID,
		invoice.AccountID,
		invoice.Amount,
		invoice.Status,
		invoice.Description,
		invoice.PaymentType,
		invoice.CardLastDigits,
		invoice.CreatedAt,
		invoice.UpdatedAt,
	)
	return err
}

// FindByID busca uma fatura pelo ID
func (r *InvoiceRepository) FindByID(id string) (*domain.Invoice, error) {
	stmt, err := r.db.Prepare("SELECT id, account_id, amount, status, description, payment_type, card_last_digits, created_at, updated_at FROM invoices WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var invoice domain.Invoice
	err = stmt.QueryRow(id).Scan(&invoice.ID, &invoice.AccountID, &invoice.Amount, &invoice.Status, &invoice.Description, &invoice.PaymentType, &invoice.CardLastDigits, &invoice.CreatedAt, &invoice.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, domain.ErrInvoiceNotFound
	}
	if err != nil {
		return nil, err
	}

	return &invoice, nil
}

// FindByAccountID busca todas as faturas de um determinado accountID
func (r *InvoiceRepository) FindByAccountID(accountID string) ([]*domain.Invoice, error) {
	stmt, err := r.db.Prepare("SELECT id, account_id, amount, status, description, payment_type, card_last_digits, created_at, updated_at FROM invoices WHERE account_id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	invoices := []*domain.Invoice{}
	for rows.Next() {
		var invoice domain.Invoice
		err = rows.Scan(&invoice.ID, &invoice.AccountID, &invoice.Amount, &invoice.Status, &invoice.Description, &invoice.PaymentType, &invoice.CardLastDigits, &invoice.CreatedAt, &invoice.UpdatedAt)
		if err != nil {
			return nil, err
		}
		invoices = append(invoices, &invoice)
	}

	return invoices, nil
}

// Update atualiza uma fatura no banco de dados
func (r *InvoiceRepository) UpdateStatus(invoice *domain.Invoice) error {
	row, err := r.db.Exec("UPDATE invoices SET status = $1, updated_at = $2 WHERE id = $3", invoice.Status, invoice.UpdatedAt, invoice.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := row.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("invoice not found")
	}
	return nil
}

func (r *InvoiceRepository) Update(invoice *domain.Invoice) error {
	return r.UpdateStatus(invoice)
}
