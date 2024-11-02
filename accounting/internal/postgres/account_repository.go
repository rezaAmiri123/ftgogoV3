package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/rezaAmiri123/ftgogoV3/accounting/internal/domain"
	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
)

type AccountReopsitory struct {
	tableName string
	db        *sql.DB
}

var _ domain.AccountRepository = (*AccountReopsitory)(nil)

func NewAccountReopsitory(tableName string, db *sql.DB) AccountReopsitory {
	return AccountReopsitory{
		tableName: tableName,
		db:        db,
	}
}

func (r AccountReopsitory) Save(ctx context.Context, account *domain.Account) error {
	const query = "INSERT INTO %s (id, name, enabled) VALUES ($1, $2, $3)"

	_, err := r.db.ExecContext(ctx, r.table(query), account.ID, account.Name, account.Enabled)

	return err
}

func (r AccountReopsitory) Find(ctx context.Context, accountID string) (*domain.Account, error) {
	const query = "SELECT name, enabled from %s where id = $1 LIMIT 1"

	account := &domain.Account{
		AggregateBase: ddd.AggregateBase{ID: accountID},
	}

	err := r.db.QueryRowContext(ctx, r.table(query), accountID).Scan(
		&account.Name,
		&account.Enabled,
	)

	return account, err
}

func (r AccountReopsitory) Update(ctx context.Context, account *domain.Account) error {
	const query = "UPDATE %s SET name = $2, enabled = $3 WHERE id = $1"

	_, err := r.db.ExecContext(ctx, r.table(query), account.ID, account.Name, account.Enabled)

	return err
}

func (r AccountReopsitory) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}
