package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/rezaAmiri123/ftgogoV3/kitchen/internal/domain"
	"github.com/stackus/errors"
)

type TicketReopsitory struct {
	tableName string
	db        *sql.DB
}

var _ domain.TicketRepository = (*TicketReopsitory)(nil)

func NewTicketReopsitory(tableName string, db *sql.DB) TicketReopsitory {
	return TicketReopsitory{
		tableName: tableName,
		db:        db,
	}
}

func (r TicketReopsitory) Save(ctx context.Context, ticket *domain.Ticket) error {
	const query = `INSERT INTO %s 
	(id, restaurant_id, line_items, accepted_at,preparing_time,ready_for_pick_up_at, 
	picked_up_at,status, pervious_status) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	lineItems, err := json.Marshal(ticket.LineItems)
	if err != nil {
		return errors.ErrInternalServerError.Err(err)
	}

	_, err = r.db.ExecContext(ctx, r.table(query), ticket.ID, ticket.RestaurantID, lineItems,
		ticket.AcceptedAt, ticket.PreparingTime, ticket.ReadyForPickUpAt, ticket.PickedUpAt,
		ticket.Status.String(), ticket.PerviousStatus.String())

	return err
}

func (r TicketReopsitory) Find(ctx context.Context, ticketID string) (*domain.Ticket, error) {
	const query = `SELECT restaurant_id, line_items, accepted_at, preparing_time, ready_for_pick_up_at, 
	picked_up_at, status, pervious_status
	from %s where id = $1 LIMIT 1`

	ticket := &domain.Ticket{
		AggregateBase: ddd.AggregateBase{ID: ticketID},
	}

	var lineItems []byte
	var status string
	var perviousStatus string

	err := r.db.QueryRowContext(ctx, r.table(query), ticketID).Scan(
		&ticket.RestaurantID,
		&lineItems,
		&ticket.AcceptedAt,
		&ticket.PreparingTime,
		&ticket.ReadyForPickUpAt,
		&ticket.PickedUpAt,
		&status,
		&perviousStatus,
	)

	err = json.Unmarshal(lineItems, &ticket.LineItems)
	if err != nil {
		return nil, errors.ErrInternalServerError.Err(err)
	}
	ticket.Status = domain.ToTicketStatus(status)
	ticket.PerviousStatus = domain.ToTicketStatus(perviousStatus)

	return ticket, err
}

func (r TicketReopsitory) Update(ctx context.Context, ticket *domain.Ticket) error {
	const query = `UPDATE %s SET 
	restaurant_id = $2, line_items = $3, accepted_at = $4, preparing_time = $5,
	ready_for_pick_up_at = $6, picked_up_at = $7, status = $8, pervious_status = $9
	WHERE id = $1`

	lineItems, err := json.Marshal(ticket.LineItems)
	if err != nil {
		return errors.ErrInternalServerError.Err(err)
	}

	_, err = r.db.ExecContext(ctx, r.table(query), ticket.ID, ticket.RestaurantID, lineItems,
		ticket.AcceptedAt, ticket.PreparingTime, ticket.ReadyForPickUpAt, ticket.PickedUpAt,
		ticket.Status.String(), ticket.PerviousStatus.String(),
	)

	return err
}

func (r TicketReopsitory) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}
