package e2e

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/cucumber/godog"
	"github.com/rezaAmiri123/ftgogoV3/customer-web/customerapi"
	"github.com/stackus/errors"
)

type consumerIDKey struct{}
type tokenKey struct{}

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
	// deleteTable := func(tableName string) {
	// 	_, _ = c.db.Exec(fmt.Sprintf("DELETE from %s", tableName))
	// }
	// deleteTable("consumer.consumers")
}

func (c *consumerFeature) register(ctx *godog.ScenarioContext) {
	ctx.Step(`^I expect the consumr is created$`, c.iExpectTheConsumrIsCreated)
	ctx.Step(`^I register a new consumer as "([^"]*)"$`, c.iRegisterANewConsumerAs)
	ctx.Step(`^no consumer named "([^"]*)" exists$`, c.noConsumerNamedExists)
	ctx.Step(`^I add address to consumer$`, c.iAddAddressToConsumer)
	ctx.Step(`^I am a registered consumer$`, c.iAmARegisteredConsumer)
	ctx.Step(`^I sign in$`, c.iSignIn)
}

func (c *consumerFeature) iExpectTheConsumrIsCreated(ctx context.Context) error {
	if err := lastResponseWas(ctx, &customerapi.RegisterConsumerResponse{}); err != nil {
		return err
	}
	return nil
}

func (c *consumerFeature) iRegisterANewConsumerAs(ctx context.Context, name string) context.Context {
	name = withRandomString(name)
	ctx, _ = c.registerConsumer(ctx, name)
	return ctx
}

func (c *consumerFeature) noConsumerNamedExists(name string) error {
	name = withRandomString(name)
	var consumerID string
	row := c.db.QueryRow("SELECT id FROM consumer.consumers WHERE name = $1", name)
	err := row.Scan(&consumerID)
	if err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		return err
	}
	return errors.ErrAlreadyExists.Msgf("the consumer `%s` already exist", name)
}

func (c *consumerFeature) registerConsumer(ctx context.Context, name string) (context.Context, error) {
	name = withRandomString(name)
	response, err := c.client.RegisterConsumer(ctx, customerapi.RegisterConsumerJSONRequestBody{
		Name: name,
	})
	ctx = setLastResponseError(ctx, response, err)
	if err != nil {
		return ctx, err
	}

	resp, err := customerapi.ParseRegisterConsumerResponse(response)
	ctx = setLastResponseError(ctx, resp, err)
	if err != nil {
		return ctx, err
	}
	return context.WithValue(ctx, consumerIDKey{}, resp.JSON201.Id), nil
}

func tokenHeader() func(ctx context.Context, req *http.Request) error {
	return func(ctx context.Context, req *http.Request) error {
		token, err := lastToken(ctx)
		if err != nil {
			return err
		}
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
		return nil
	}
}
func (c *consumerFeature) iAddAddressToConsumer(ctx context.Context) context.Context {
	name := withRandomString("1")
	response, err := c.client.AddConsumerAddress(ctx, customerapi.AddConsumerAddressJSONRequestBody{
		Name: name,
		Address: customerapi.Address{
			Street1: "street1",
		},
	}, tokenHeader())

	ctx = setLastResponseError(ctx, response, err)
	if err != nil {
		return ctx
	}

	resp, err := customerapi.ParseAddConsumerAddressResponse(response)
	ctx = setLastResponseError(ctx, resp, err)
	if err != nil {
		return ctx
	}
	return ctx

}

func (c *consumerFeature) iAmARegisteredConsumer(ctx context.Context) context.Context {
	name := withRandomString("name")
	ctx, _ = c.registerConsumer(ctx, name)
	return ctx
}

func (c *consumerFeature) iSignIn(ctx context.Context) context.Context {
	consumerID, err := lastConsumerID(ctx)
	if err != nil {
		return ctx
	}
	response, err := c.client.SignInConsumer(ctx, customerapi.SignInConsumerJSONRequestBody{
		ConsumerId: consumerID,
	})
	ctx = setLastResponseError(ctx, response, err)
	if err != nil {
		return ctx
	}

	resp, err := customerapi.ParseSignInConsumerResponse(response)
	ctx = setLastResponseError(ctx, resp, err)
	if err != nil {
		return ctx
	}
	return context.WithValue(ctx, tokenKey{}, resp.JSON200.Token)
}

func lastConsumerID(ctx context.Context) (string, error) {
	v := ctx.Value(consumerIDKey{})
	if v == nil {
		return "", errors.ErrNotFound.Msg("no consumer ID to work with")
	}
	return v.(string), nil
}

func lastToken(ctx context.Context) (string, error) {
	v := ctx.Value(tokenKey{})
	if v == nil {
		return "", errors.ErrNotFound.Msg("no token to work with")
	}
	return v.(string), nil
}
