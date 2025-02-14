package grpc

import (
	"context"

	"github.com/rezaAmiri123/ftgogoV3/accounting/accountingpb"
	"github.com/rezaAmiri123/ftgogoV3/accounting/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/internal/errorsotel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
)

type server struct {
	app application.App
	accountingpb.UnimplementedAccountingServiceServer
}

var _ accountingpb.AccountingServiceServer = (*server)(nil)

func RegisterServer(app application.App, register grpc.ServiceRegistrar) error {
	accountingpb.RegisterAccountingServiceServer(register, server{app: app})
	return nil
}

func (s server) GetAccount(ctx context.Context, request *accountingpb.GetAccountRequest) (*accountingpb.GetAccountResponse, error) {
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		attribute.String("AccountID", request.GetAccountID()),
	)

	account, err := s.app.GetAccount(ctx, application.GetAccount{
		ID: request.GetAccountID(),
	})
	if err != nil {
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	return &accountingpb.GetAccountResponse{
		AccountID: account.ID(),
		Enabled:   account.Enabled,
	}, nil
}
func (s server) DisableAccount(ctx context.Context, request *accountingpb.DisableAccountRequest) (*accountingpb.DisableAccountResponse, error) {
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		attribute.String("AccountID", request.GetAccountID()),
	)

	err := s.app.DisableAccount(ctx, application.DisableAccount{
		ID: request.GetAccountID(),
	})
	if err != nil {
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	return &accountingpb.DisableAccountResponse{}, nil
}

func (s server) EnableAccount(ctx context.Context, request *accountingpb.EnableAccountRequest) (*accountingpb.EnableAccountResponse, error) {
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		attribute.String("AccountID", request.GetAccountID()),
	)

	err := s.app.EnableAccount(ctx, application.EnableAccount{
		ID: request.GetAccountID(),
	})
	if err != nil {
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	return &accountingpb.EnableAccountResponse{}, nil
}

func (s server) CreateAccount(ctx context.Context, request *accountingpb.CreateAccountRequest) (*accountingpb.CreateAccountResponse, error) {
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		attribute.String("ID", request.GetID()),
		attribute.String("Name", request.GetName()),
	)

	err := s.app.RegisterAccount(ctx, application.RegisterAccount{
		ID:   request.GetID(),
		Name: request.GetName(),
	})
	if err != nil {
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	return &accountingpb.CreateAccountResponse{AccountID: request.GetID()}, nil
}

func (s server) AuthorizeOrderByAccount(ctx context.Context, request *accountingpb.AuthorizeOrderByAccountRequest) (*accountingpb.AuthorizeOrderByAccountResponse, error) {
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		attribute.String("AccountID", request.GetAccountID()),
		attribute.String("OrderID", request.GetOrderID()),
	)

	err := s.app.AuthorizeOrderByAccount(ctx, application.AuthorizeOrderByAccount{
		ID:         request.GetAccountID(),
		OrderID:    request.GetOrderID(),
		OrderTotal: int(request.GetOrderTotal()),
	})
	if err != nil {
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}
	return &accountingpb.AuthorizeOrderByAccountResponse{}, nil
}
