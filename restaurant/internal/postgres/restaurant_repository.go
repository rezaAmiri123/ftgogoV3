package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/rezaAmiri123/ftgogoV3/restaurant/internal/domain"
	"github.com/stackus/errors"
)

type RestaurantRepository struct {
	tableName string
	db        *sql.DB
}

var _ domain.RestaurantRepository = (*RestaurantRepository)(nil)

func NewRestaurantRepository(tableName string, db *sql.DB) RestaurantRepository {
	return RestaurantRepository{
		tableName: tableName,
		db:        db,
	}
}

func (r RestaurantRepository) Save(ctx context.Context, restaurant *domain.Restaurant) error {
	const query = "INSERT INTO %s (id, name, address, menu_items) VALUES ($1, $2, $3, $4)"

	menuItems, err := json.Marshal(restaurant.MenuItems)
	if err != nil {
		return errors.ErrInternalServerError.Err(err)
	}
	address, err := json.Marshal(restaurant.Address)
	if err != nil {
		return errors.ErrInternalServerError.Err(err)
	}
	_, err = r.db.ExecContext(ctx, r.table(query), restaurant.ID, restaurant.Name, address, menuItems)

	return err
}

func (r RestaurantRepository) Find(ctx context.Context, restaurantID string) (*domain.Restaurant, error) {
	const query = "SELECT name, address, menu_items from %s where id = $1 LIMIT 1"

	restaurant := &domain.Restaurant{
		ID: restaurantID,
	}

	var address []byte
	var menuItems []byte

	err := r.db.QueryRowContext(ctx, r.table(query), restaurantID).Scan(
		&restaurant.Name,
		&address,
		&menuItems,
	)

	err = json.Unmarshal(address, &restaurant.Address)
	if err != nil {
		return nil, errors.ErrInternalServerError.Err(err)
	}
	err = json.Unmarshal(menuItems, &restaurant.MenuItems)
	if err != nil {
		return nil, errors.ErrInternalServerError.Err(err)
	}

	return restaurant, err
}

func (r RestaurantRepository) Update(ctx context.Context, restaurant *domain.Restaurant) error {
	const query = "UPDATE %s SET name = $2, address = $3, menu_items = $4 WHERE id = $1"

	menuItems, err := json.Marshal(restaurant.MenuItems)
	if err != nil {
		return errors.ErrInternalServerError.Err(err)
	}
	address, err := json.Marshal(restaurant.Address)
	if err != nil {
		return errors.ErrInternalServerError.Err(err)
	}

	_, err = r.db.ExecContext(ctx, r.table(query), restaurant.ID, restaurant.Name, address, menuItems)

	return err
}

func (r RestaurantRepository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}
