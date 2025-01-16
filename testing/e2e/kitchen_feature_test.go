package e2e

import (
	"context"
	"database/sql"
	"time"

	"github.com/cucumber/godog"
	"github.com/google/uuid"
	"github.com/rezaAmiri123/ftgogoV3/store-web/storeapi"
	"github.com/stackus/errors"
)

type ticketIDKey struct{}

type kitchenFeature struct {
	client *storeapi.Client
	db     *sql.DB
}

func (f *kitchenFeature) init(cfg featureConfig) (err error) {
	f.db, err = sql.Open("pgx", "postgres://ftgogo_user:ftgogo_pass@localhost:5432/ftgogo?sslmode=disable")
	if err != nil {
		return err
	}
	f.client, err = storeapi.NewClient("http://localhost:8000/store/v1")
	return
}

func (f *kitchenFeature) reset() {
	// deleteTable := func(tableName string) {
	// 	_, _ = f.db.Exec(fmt.Sprintf("DELETE from %s", tableName))
	// }
	// deleteTable("kitchen.tickets")
}

func (f *kitchenFeature) register(ctx *godog.ScenarioContext) {
	ctx.Step(`^I accept a ticket$`, f.iAcceptATicket)
	ctx.Step(`^I expect the ticket is accepted$`, f.iExpectTheTicketIsAccepted)
	ctx.Step(`^no accepted ticket is exists$`, f.noAcceptedTicketIsExists)
}
func (f *kitchenFeature) iAcceptATicket(ctx context.Context) context.Context {
	orderID, err := lastOrderID(ctx)
	if err != nil {
		return ctx
	}
	orderUUID, _ := uuid.Parse(orderID)
	response, err := f.client.AcceptTicket(ctx, orderUUID, storeapi.AcceptTicketJSONRequestBody{
		ReadyBy: time.Now().Add(time.Hour),
	}, tokenHeader())
	ctx = setLastResponseError(ctx, response, err)
	if err != nil {
		return ctx
	}

	resp, err := storeapi.ParseAcceptTicketResponse(response)
	ctx = setLastResponseError(ctx, resp, err)
	if err != nil {
		return ctx
	}

	return ctx
}

func (f *kitchenFeature) iExpectTheTicketIsAccepted(ctx context.Context) error {
	if err := lastResponseWas(ctx, &storeapi.AcceptTicketResponse{}); err != nil {
		return err
	}
	return nil
}

func (f *kitchenFeature) noAcceptedTicketIsExists(ctx context.Context) context.Context {
	return ctx
}

func lastTicketID(ctx context.Context) (string, error) {
	v := ctx.Value(restaurantIDKey{})
	if v == nil {
		return "", errors.ErrNotFound.Msg("no restaurant ID to work with")
	}
	return v.(string), nil
}
