package e2e

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/cucumber/godog"
	"github.com/rezaAmiri123/ftgogoV3/customer-web/customerapi"
	"github.com/stackus/errors"
)

type consumerIDKey struct{}

type consumerFeature struct {
	client *customerapi.Client
	db     *sql.DB
}

func (c *consumerFeature) init(cfg featureConfig) (err error) {
	c.db, err = sql.Open("pgx", "postgres://ftgogo_user:ftgogo_pass@localhost:5432/ftgogo?sslmode=disable")
	if err != nil {
		return err
	}
	c.client, err = customerapi.NewClient("http://localhost:8000/api/v1")
	return
}

func (c *consumerFeature) reset() {
	deleteTable := func(tableName string){
		_,_ = c.db.Exec(fmt.Sprintf("DELETE from %s", tableName))
	}
	deleteTable("consumer.consumers")
}

func (c *consumerFeature) register(ctx *godog.ScenarioContext) {
	ctx.Step(`^I expect the consumr is created$`, c.iExpectTheConsumrIsCreated)
	ctx.Step(`^I register a new consumer as "([^"]*)"$`, c.iRegisterANewConsumerAs)
	ctx.Step(`^no consumer named "([^"]*)" exists$`, c.noConsumerNamedExists)
}

func (c *consumerFeature) iExpectTheConsumrIsCreated(ctx context.Context) error {
	if err := lastResponseWas(ctx, &customerapi.RegisterConsumerResponse{}); err != nil {
		return err
	}
	return nil
}

func (c *consumerFeature) iRegisterANewConsumerAs(ctx context.Context, name string) context.Context {
	response, err := c.client.RegisterConsumer(context.Background(), customerapi.RegisterConsumerJSONRequestBody{
		Name: name,
	})
	ctx = setLastResponseError(ctx, response, err)
	if err != nil {
		return ctx
	}

	resp, err := customerapi.ParseRegisterConsumerResponse(response)
	ctx = setLastResponseError(ctx, resp, err)
	if err != nil {
		return ctx
	}
	return context.WithValue(ctx, consumerIDKey{}, resp.JSON201.Id)
}

func (c *consumerFeature) noConsumerNamedExists(name string) error {
	var consumerID string
	row := c.db.QueryRow("SELECT id FROM consumer.consumers WHERE name = $1", name)
	err := row.Scan(&consumerID)
	if err== sql.ErrNoRows{
		return nil
	}else if err != nil{
		return err
	}
	return errors.ErrAlreadyExists.Msgf("the consumer `%s` already exist", name)
}
