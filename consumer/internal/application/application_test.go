package application

import (
	"context"
	"errors"
	"testing"

	"github.com/rezaAmiri123/ftgogoV3/consumer/internal/domain"
	"github.com/rezaAmiri123/ftgogoV3/consumer/internal/domain/mocks"
	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	dddmocks "github.com/rezaAmiri123/ftgogoV3/internal/ddd/mocks"
)

func TestApplication_UpdateConsumerAddress(t *testing.T) {
	type Mocks struct {
		consumers *mocks.ConsumerRepository
		publisher *dddmocks.EventPublisher
	}
	type args struct {
		ctx    context.Context
		update UpdateConsumerAddress
	}
	tests := map[string]struct {
		args    args
		on      func(f Mocks)
		wantErr bool
	}{
		"OK": {
			args: args{
				ctx: context.Background(),
				update: UpdateConsumerAddress{
					ConsumerID: "consumer-id",
					AddressID:  "address-id",
					Address:    domain.Address{Street1: "new-street"},
				},
			},
			on: func(f Mocks) {
				consumer := &domain.Consumer{
					AggregateBase: ddd.AggregateBase{ID: "consumer-id"},
					Name:          "consumer-name",
					Addresses:     map[string]domain.Address{"address-id": domain.Address{Street1: "street"}},
				}
				f.consumers.On("Find", context.Background(), "consumer-id").Return(consumer, nil)
				
				consumer.Addresses["address-id"] = domain.Address{Street1: "new-street"}
				f.consumers.On("Update", context.Background(), consumer).Return(nil)
				// f.publisher.On("Publish", context.Background(), mock.AnythingOfType("[]ddd.event")).Return(nil)
			},
			wantErr: false,
		},
		"Save consumer failed": {
			args: args{
				ctx: context.Background(),
				update: UpdateConsumerAddress{
					ConsumerID: "consumer-id",
					AddressID:  "address-id",
					Address:    domain.Address{Street1: "new-street"},
				},
			},
			on: func(f Mocks) {
				f.consumers.On("Find", context.Background(), "consumer-id").Return(&domain.Consumer{
					AggregateBase: ddd.AggregateBase{ID: "consumer-id"},
					Name:          "consumer-name",
					Addresses:     map[string]domain.Address{"address-id": domain.Address{Street1: "street"}},
				}, nil)
				f.consumers.On("Update", context.Background(), &domain.Consumer{
					AggregateBase: ddd.AggregateBase{ID: "consumer-id"},
					Name:          "consumer-name",
					Addresses:     map[string]domain.Address{"address-id": domain.Address{Street1: "new-street"}},
				}).Return(errors.New("save failed"))
			},
			wantErr: true,
		},
		"Find consumer failed": {
			args: args{
				ctx: context.Background(),
				update: UpdateConsumerAddress{
					ConsumerID: "consumer-id",
					AddressID:  "address-id",
					Address:    domain.Address{Street1: "new-street"},
				},
			},
			on: func(f Mocks) {
				f.consumers.On("Find", context.Background(), "consumer-id").Return(&domain.Consumer{
					AggregateBase: ddd.AggregateBase{ID: "consumer-id"},
					Name:          "consumer-name",
					Addresses:     map[string]domain.Address{"address-id": domain.Address{Street1: "street"}},
				}, errors.New("find failed"))
			},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			m := Mocks{
				consumers: mocks.NewConsumerRepository(t),
				publisher: dddmocks.NewEventPublisher(t),
			}
			app := New(m.consumers, m.publisher)
			if tt.on != nil {
				tt.on(m)
			}
			if err := app.UpdateConsumerAddress(tt.args.ctx, tt.args.update); (err != nil) != tt.wantErr {
				t.Errorf("UpdateConsumerAddress() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
