package e2e

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/cucumber/godog"
	"github.com/rezaAmiri123/ftgogoV3/customer-web/customerapi"
)

type orderIDKey struct{}

type orderFeature struct {
	client *customerapi.Client
	db     *sql.DB
}

func (f *orderFeature) init(cfg featureConfig) (err error) {
	f.db, err = sql.Open("pgx", "postgres://ftgogo_user:ftgogo_pass@localhost:5432/ftgogo?sslmode=disable")
	if err != nil {
		return err
	}
	f.client, err = customerapi.NewClient("http://localhost:8000/api/v1")
	return
}

func (f *orderFeature) reset() {
	deleteTable := func(tableName string) {
		_, _ = f.db.Exec(fmt.Sprintf("DELETE from %s", tableName))
	}

	deleteTable("orders.orders")
}

func (f *orderFeature) register(ctx *godog.ScenarioContext) {
	ctx.Step(`^I create a new order$`, f.iCreateANewOrder)
	ctx.Step(`^I expect the order is created$`, f.iExpectTheOrderIsCreated)
	ctx.Step(`^no order for registered consumer exists$`, f.noOrderForRegisteredConsumerExists)
}

func (f *orderFeature) iCreateANewOrder(ctx context.Context) context.Context {
	consumerID, err := lastConsumerID(ctx)
	if err != nil {
		return ctx
	}
	restaurantID, err := lastRestaurantID(ctx)
	if err != nil {
		return ctx
	}

	response, err := f.client.CreateOrder(ctx, customerapi.CreateOrderJSONRequestBody{
		AddressId: "1",
		ConsumerId: consumerID,
		LineItems: customerapi.MenuItemQuantities{"1":1},
		RestaurantId: restaurantID,
	}, tokenHeader())

	ctx = setLastResponseError(ctx, response, err)
	if err != nil {
		return ctx
	}

	resp, err := customerapi.ParseCreateOrderResponse(response)
	ctx = setLastResponseError(ctx, resp, err)
	if err != nil {
		return ctx
	}
	return ctx
}

func (f *orderFeature) iExpectTheOrderIsCreated(ctx context.Context) error {
	if err := lastResponseWas(ctx, &customerapi.CreateOrderResponse{}); err != nil {
		return err
	}
	return nil
}

func (f *orderFeature) noOrderForRegisteredConsumerExists(ctx context.Context) context.Context {
	return ctx
}
