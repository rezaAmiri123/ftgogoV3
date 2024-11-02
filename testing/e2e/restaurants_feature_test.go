package e2e

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/cucumber/godog"
	"github.com/google/uuid"
	"github.com/rezaAmiri123/ftgogoV3/store-web/storeapi"
	"github.com/stackus/errors"
)

type restaurantIDKey struct{}

type restaurantFeature struct {
	client *storeapi.Client
	db *sql.DB
}

func (f *restaurantFeature) init(cfg featureConfig) (err error) {
	f.db, err = sql.Open("pgx", "postgres://ftgogo_user:ftgogo_pass@localhost:5432/ftgogo?sslmode=disable")
	if err != nil {
		return err
	}
	f.client, err = storeapi.NewClient("http://localhost:8000/store/v1")
	return
}

func (f *restaurantFeature) reset() {
	deleteTable := func(tableName string) {
		_, _ = f.db.Exec(fmt.Sprintf("DELETE from %s", tableName))
	}
	deleteTable("restaurant.restaurants")
}

func (f *restaurantFeature) register(ctx *godog.ScenarioContext) {
	ctx.Step(`^I create a new restarant$`, f.iCreateANewRestarant)
	ctx.Step(`^I update the restarant menu$`, f.iUpdateTheRestarantMenu)
}

func (f *restaurantFeature) iCreateANewRestarant(ctx context.Context) context.Context {
	// const query = `INSERT INTO restaurant.restaurants (id, name, address, menu_items) VALUES  (
	// 'ea7ac1aa-b14b-4904-bfb7-bb91e2fe6920', 'name', 
	// decode('7b2253747265657431223a2231222c2253747265657432223a2231222c2243697479223a2231222c225374617465223a2231222c225a6970223a2231227d', 'hex'),
	// decode('7b7d', 'hex')
	// )`
	// response, err := f.db.Exec(query)
	// ctx = setLastResponseError(ctx, response, err)
	// if err != nil {
	// 	return ctx
	// }

	// return context.WithValue(ctx, restaurantIDKey{}, "ea7ac1aa-b14b-4904-bfb7-bb91e2fe6920")
	response, err := f.client.CreateRestaurant(ctx, storeapi.CreateRestaurantJSONRequestBody{
		Name: "name",
		Address: storeapi.Address{Street1: "street1"},
	},tokenHeader())
	ctx = setLastResponseError(ctx, response, err)
	if err != nil {
		fmt.Println("iCreateANewRestarant: ",err.Error())
		return ctx
	}

	resp, err := storeapi.ParseCreateRestaurantResponse(response)
	ctx = setLastResponseError(ctx, resp, err)
	if err != nil {
		return ctx
	}
	return context.WithValue(ctx, restaurantIDKey{}, resp.JSON201.Id)

}
func (f *restaurantFeature) iUpdateTheRestarantMenu(ctx context.Context) context.Context {
	restaurantID,_ := lastRestaurantID(ctx)
	restaurantUUID , err := uuid.Parse(restaurantID)
	response, err := f.client.UpdateRestaurantMenu(ctx, restaurantUUID,storeapi.UpdateRestaurantMenuJSONRequestBody{
		Menu: struct{MenuItems []storeapi.MenuItem "json:\"menu_items\""}{MenuItems: []storeapi.MenuItem{storeapi.MenuItem{
			Id: "1",
			Name: "name",
			Price: 1,
		}}},
	},tokenHeader())
	ctx = setLastResponseError(ctx, response, err)
	if err != nil {
		return ctx
	}

	resp, err := storeapi.ParseUpdateRestaurantMenuResponse(response)
	ctx = setLastResponseError(ctx, resp, err)
	if err != nil {
		return ctx
	}
	return ctx

}

func lastRestaurantID(ctx context.Context) (string, error) {
	v := ctx.Value(restaurantIDKey{})
	if v == nil {
		return "", errors.ErrNotFound.Msg("no restaurant ID to work with")
	}
	return v.(string), nil
}
