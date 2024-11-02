package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/rezaAmiri123/ftgogoV3/delivery/internal/domain"
	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/stackus/errors"
)

type DeliveryRepository struct {
	tableName string
	db        *sql.DB
}

var _ domain.DeliveryRepository = (*DeliveryRepository)(nil)

func NewDeliveryRepository(tableName string, db *sql.DB) DeliveryRepository {
	return DeliveryRepository{
		tableName: tableName,
		db:        db,
	}
}

func (r DeliveryRepository) Save(ctx context.Context, delivery *domain.Delivery) error {
	const query = `INSERT INTO %s 
	(id, restaurant_id, assigned_courier_id, pick_up_address,delivery_address, 
	status, pick_up_time, ready_by) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	pickUpAddress, err := json.Marshal(delivery.PickUpAddress)
	if err != nil {
		return errors.ErrInternalServerError.Err(err)
	}
	deliveryAddress, err := json.Marshal(delivery.DeliveryAddress)
	if err != nil {
		return errors.ErrInternalServerError.Err(err)
	}

	_, err = r.db.ExecContext(ctx, r.table(query),
		delivery.ID,
		delivery.RestaurantID,
		delivery.AssignedCourierID,
		pickUpAddress,
		deliveryAddress,
		delivery.Status.String(),
		delivery.PickUpTime,
		delivery.ReadyBy,
	)

	return err
}

func (r DeliveryRepository) Find(ctx context.Context, deliveryID string) (*domain.Delivery, error) {
	const query = `SELECT restaurant_id, assigned_courier_id, pick_up_address,delivery_address, 
	status, pick_up_time, ready_by
	from %s where id = $1 LIMIT 1`

	delivery := &domain.Delivery{
		AggregateBase: ddd.AggregateBase{ID: deliveryID},
	}

	var pickUpAddress []byte
	var deliveryAddress []byte
	var status string

	err := r.db.QueryRowContext(ctx, r.table(query), deliveryID).Scan(
		&delivery.RestaurantID,
		&delivery.AssignedCourierID,
		&pickUpAddress,
		&deliveryAddress,
		&status,
		&delivery.PickUpTime,
		&delivery.ReadyBy,
	)

	err = json.Unmarshal(pickUpAddress, &delivery.PickUpAddress)
	if err != nil {
		return nil, errors.ErrInternalServerError.Err(err)
	}
	err = json.Unmarshal(deliveryAddress, &delivery.DeliveryAddress)
	if err != nil {
		return nil, errors.ErrInternalServerError.Err(err)
	}

	delivery.Status = domain.ToDeliveryStatus(status)

	return delivery, err
}

func (r DeliveryRepository) Update(ctx context.Context, delivery *domain.Delivery) error {
	const query = `UPDATE %s SET 
	restaurant_id = $2, assigned_courier_id = $3, pick_up_address = $4, delivery_address = $5,
	status = $6, pick_up_time = $7, ready_by = $8
	WHERE id = $1`

	pickUpAddress, err := json.Marshal(delivery.PickUpAddress)
	if err != nil {
		return errors.ErrInternalServerError.Err(err)
	}
	deliveryAddress, err := json.Marshal(delivery.DeliveryAddress)
	if err != nil {
		return errors.ErrInternalServerError.Err(err)
	}

	_, err = r.db.ExecContext(ctx, r.table(query),
		delivery.ID,
		delivery.RestaurantID,
		delivery.AssignedCourierID,
		pickUpAddress,
		deliveryAddress,
		delivery.Status.String(),
		delivery.PickUpTime,
		delivery.ReadyBy,
	)

	return err
}

func (r DeliveryRepository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}
