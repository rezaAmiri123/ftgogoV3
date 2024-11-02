package e2e

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"testing"

	"github.com/cucumber/godog"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/rdumont/assistdog"
	"github.com/stackus/errors"
)
var assist = assistdog.NewDefault()
type lastResponseKey struct{}
type lastErrorKey struct{}

type featureConfig struct {
}

type feature interface {
	init(cfg featureConfig) error
	register(ctx *godog.ScenarioContext)
	reset()
}

func TestEndToEnd(t *testing.T) {
	assist.RegisterComparer(float64(0.0), func(raw string, actual interface{}) error {
		af, ok := actual.(float64)
		if !ok {
			return fmt.Errorf("%v is not a float64", actual)
		}
		ef, err := strconv.ParseFloat(raw, 64)
		if err != nil {
			return err
		}

		if ef != af {
			return fmt.Errorf("expected %v, but got %v", ef, af)
		}

		return nil
	})
	assist.RegisterParser(float64(0.0), func(raw string) (interface{}, error) {
		return strconv.ParseFloat(raw, 64)
	})

	cfg := featureConfig{}

	features, err := func(fs ...feature) ([]feature, error) {
		features := make([]feature, len(fs))
		for i, f := range fs {
			err := f.init(cfg)
			if err != nil {
				return features, err
			}
			features[i] = f
		}
		return features, nil
	}(
		&consumerFeature{},
		&restaurantFeature{},
		&orderFeature{},
		&kitchenFeature{},
	)
	if err != nil {
		t.Fatal(err)
	}

	featurePaths := []string{
		"features/consumers",
		"features/orders",
		"features/kitchens",
	}

	suite := godog.TestSuite{
		Name: "ftgogo",
		ScenarioInitializer: func(ctx *godog.ScenarioContext) {
			for _, f := range features {
				f.register(ctx)
			}
			ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
				for _, f := range features {
					f.reset()
				}
				return ctx, nil
			})
		},
		Options: &godog.Options{
			Format:    "pretty",
			Paths:     featurePaths,
			Randomize: -1,
		},
	}
	if status := suite.Run(); status != 0 {
		t.Error("end to end feature test failed with status:", status)
	}
}

func setLastResponseError(ctx context.Context, resp any, err error) context.Context {
	ctx = context.WithValue(ctx, lastResponseKey{}, resp)
	ctx = context.WithValue(ctx, lastErrorKey{}, err)
	return ctx
}

func lastResponseWas(ctx context.Context, resp any)error{
	r := ctx.Value(lastResponseKey{})
	if reflect.ValueOf(r).Kind() == reflect.Ptr&&reflect.ValueOf(r).IsNil(){
		e := ctx.Value(lastErrorKey{})
		if e==nil{
return errors.ErrUnknown.Msg("no last response of error")
		}
		return e.(error)
	}
	if reflect.TypeOf(r) == reflect.TypeOf(resp){
		return nil
	}
	return errors.ErrBadRequest.Msgf("last response was `%v`", r)
}