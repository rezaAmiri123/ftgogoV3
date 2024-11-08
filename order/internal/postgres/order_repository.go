package postgres

// import (
// 	"context"
// 	"database/sql"
// 	"encoding/json"
// 	"fmt"

// 	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
// 	"github.com/rezaAmiri123/ftgogoV3/order/internal/domain"
// 	"github.com/stackus/errors"
// )

// type OrderRepository struct {
// 	tableName string
// 	db        *sql.DB
// }

// var _ domain.OrderRepository = (*OrderRepository)(nil)

// func NewOrderRepository(tableName string, db *sql.DB) OrderRepository {
// 	return OrderRepository{
// 		tableName: tableName,
// 		db:        db,
// 	}
// }

// func (r OrderRepository) Save(ctx context.Context, order *domain.Order) error {
// 	const query = `INSERT INTO %s 
// 	(id, consumer_id, restaurant_id, ticket_id, line_items, status, deliver_at, deliver_to) 
// 	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

// 	lineItems, err := json.Marshal(order.LineItems)
// 	if err != nil {
// 		return errors.ErrInternalServerError.Err(err)
// 	}
// 	deliverTo, err := json.Marshal(order.DeliverTo)
// 	if err != nil {
// 		return errors.ErrInternalServerError.Err(err)
// 	}
// 	_, err = r.db.ExecContext(ctx, r.table(query),
// 		order.ID,
// 		order.ConsumerID,
// 		order.RestaurantID,
// 		order.TicketID,
// 		lineItems,
// 		order.Status.String(),
// 		order.DeliverAt,
// 		deliverTo,
// 	)
// 	return err
// }

// func (r OrderRepository) Find(ctx context.Context, orderID string) (*domain.Order, error) {
// 	const query = `SELECT consumer_id, restaurant_id, ticket_id, line_items, status, 
// 	deliver_at, deliver_to
// 	from %s where id = $1 LIMIT 1`

// 	order := &domain.Order{
// 		AggregateBase: ddd.AggregateBase{ID: orderID},
// 	}

// 	var lineItems []byte
// 	var deliverTo []byte
// 	var status string

// 	err := r.db.QueryRowContext(ctx, r.table(query), orderID).Scan(
// 		&order.ConsumerID,
// 		&order.RestaurantID,
// 		&order.TicketID,
// 		&lineItems,
// 		&status,
// 		&order.DeliverAt,
// 		&deliverTo,
// 	)

// 	err = json.Unmarshal(lineItems, &order.LineItems)
// 	if err != nil {
// 		return nil, errors.ErrInternalServerError.Err(err)
// 	}
// 	err = json.Unmarshal(deliverTo, &order.DeliverTo)
// 	if err != nil {
// 		return nil, errors.ErrInternalServerError.Err(err)
// 	}

// 	order.Status = domain.ToOrderStatus(status)

// 	return order, err
// }

// func (r OrderRepository) Update(ctx context.Context, order *domain.Order) error {
// 	const query = `UPDATE %s SET 
// 	consumer_id = $2, restaurant_id = $3, ticket_id = $4, line_items = $5,
// 	status = $6, deliver_at = $7, deliver_to = $8
// 	WHERE id = $1`

// 	lineItems, err := json.Marshal(order.LineItems)
// 	if err != nil {
// 		return errors.ErrInternalServerError.Err(err)
// 	}
// 	deliverTo, err := json.Marshal(order.DeliverTo)
// 	if err != nil {
// 		return errors.ErrInternalServerError.Err(err)
// 	}

// 	_, err = r.db.ExecContext(ctx, r.table(query),
// 		order.ID,
// 		order.ConsumerID,
// 		order.RestaurantID,
// 		order.TicketID,
// 		lineItems,
// 		order.Status.String(),
// 		order.DeliverAt,
// 		deliverTo,
// 	)

// 	return err
// }

// func (r OrderRepository) table(query string) string {
// 	return fmt.Sprintf(query, r.tableName)
// }
