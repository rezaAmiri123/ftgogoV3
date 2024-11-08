package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/rezaAmiri123/ftgogoV3/delivery/internal/domain"
	"github.com/stackus/errors"
)

type CourierRepository struct {
	tableName string
	db        *sql.DB
}

var _ domain.CourierRepository = (*CourierRepository)(nil)

func NewCourierRepository(tableName string, db *sql.DB) CourierRepository {
	return CourierRepository{
		tableName: tableName,
		db:        db,
	}
}

func (r CourierRepository) Find(ctx context.Context, courierID string) (*domain.Courier, error) {
	const query = `SELECT plan, available FROM %s WHERE id = $1`
	
	courier := domain.NewCourier(courierID)

	var plan []byte

	err := r.db.QueryRowContext(ctx, r.table(query), courierID).Scan(
		&plan,
		courier.Available,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrCourierNotFound
		}
	}

	err = json.Unmarshal(plan, &courier.Plan)
	if err != nil {
		return nil, errors.ErrInternalServerError.Err(err)
	}

	return courier, nil
}

func (r CourierRepository) FindOrCreate(ctx context.Context, courierID string) (*domain.Courier, error) {
	courier, err := r.Find(ctx, courierID)
	if err != nil {
		if errors.Is(err, domain.ErrCourierNotFound) {
			courier := domain.NewCourier(courierID)
			courier.Plan = domain.Plan{}
			courier.Available = true
			err = r.Save(ctx, courier)
			if err != nil {
				return nil, err
			}
			return courier, nil
		} else {
			return nil, err
		}
	}
	return courier, nil
}

func (r CourierRepository) FindFirstAvailable(ctx context.Context) (*domain.Courier, error) {
	// "SELECT id, plan, available FROM %s WHERE available ORDER BY modified_at DESC LIMIT 1"
	const query = `SELECT id, plan, available FROM %s 
		WHERE available ORDER BY updated_at DESC LIMIT 1`

	var plan []byte
	var id string
	var available bool

	err := r.db.QueryRowContext(ctx, r.table(query)).Scan(&id, &plan, &available)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// for the demo only; in the real world you can't count on instant courier (just add water!)
			return r.FindOrCreate(ctx, uuid.New().String())
		}
	}
	courier := domain.NewCourier(id)
	courier.Available = available
	
	err = json.Unmarshal(plan, &courier.Plan)
	if err != nil {
		return nil, errors.ErrInternalServerError.Err(err)
	}

	return courier, nil

}
func (r CourierRepository) Save(ctx context.Context, courier *domain.Courier) error {
	const query = `INSERT INTO %s (id, plan, available) VALUES ($1, $2, $3)`
	plan, err := json.Marshal(courier.Plan)
	if err != nil {
		return err
	}
	_, err = r.db.ExecContext(ctx, r.table(query), courier.ID(), plan, courier.Available)
	return err
}

func (r CourierRepository) Update(ctx context.Context, courier *domain.Courier) error {
	const query = `UPDATE %s SET
		plan = $2, available = $3 WHERE id = $1`

	plan, err := json.Marshal(courier.Plan)
	if err != nil {
		return errors.ErrInternalServerError.Err(err)
	}

	_, err = r.db.ExecContext(ctx, r.table(query), courier.ID(), plan, courier.Available)
	return err
}

func (r CourierRepository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}
