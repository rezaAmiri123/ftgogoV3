package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/rezaAmiri123/ftgogoV3/consumer/internal/domain"
	"github.com/stackus/errors"
)

type ConsumerReopsitory struct {
	tableName string
	db        *sql.DB
}

var _ domain.ConsumerRepository = (*ConsumerReopsitory)(nil)

func NewConsumerReopsitory(tableName string, db *sql.DB) ConsumerReopsitory {
	return ConsumerReopsitory{
		tableName: tableName,
		db:        db,
	}
}

func (r ConsumerReopsitory) Save(ctx context.Context, consumer *domain.Consumer) error {
	const query = "INSERT INTO %s (id, name, addresses) VALUES ($1, $2, $3)"

	addresses, err := json.Marshal(consumer.Addresses)
	if err != nil {
		return errors.ErrInternalServerError.Err(err)
	}
	_, err = r.db.ExecContext(ctx, r.table(query), consumer.ID, consumer.Name, addresses)

	return err
}

func (r ConsumerReopsitory) Find(ctx context.Context, consumerID string) (*domain.Consumer, error) {
	const query = "SELECT name, addresses from %s where id = $1 LIMIT 1"

	consumer := &domain.Consumer{
		ID: consumerID,
	}

	var addresses []byte

	err := r.db.QueryRowContext(ctx, r.table(query), consumerID).Scan(
		&consumer.Name,
		&addresses,
	)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(addresses, &consumer.Addresses)
	if err != nil {
		return nil, errors.ErrInternalServerError.Err(err)
	}
	return consumer, err
}

func (r ConsumerReopsitory) Update(ctx context.Context, consumer *domain.Consumer) error {
	const query = "UPDATE %s SET name = $2, addresses = $3 WHERE id = $1"

	addresses, err := json.Marshal(consumer.Addresses)
	if err != nil {
		return errors.ErrInternalServerError.Err(err)
	}

	_, err = r.db.ExecContext(ctx, r.table(query), consumer.ID, consumer.Name, addresses)

	return err
}

func (r ConsumerReopsitory) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}
