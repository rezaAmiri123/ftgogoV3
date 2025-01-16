package domain

import (
	"reflect"
	"testing"
	"time"

	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
	"github.com/rezaAmiri123/ftgogoV3/internal/es"
	"github.com/stretchr/testify/assert"
)

func TestOrder_ApplyEvent(t *testing.T) {
	aggregate := es.NewAggregate("id", "aggregate")
	deliverAt := time.Now()
	type fields struct {
		ConsumerID   string
		RestaurantID string
		TicketID     string
		LineItems    []LineItem
		Status       OrderStatus
		DeliverAt    time.Time
		DeliverTo    Address
	}
	type args struct {
		event ddd.Event
	}
	tests := map[string]struct {
		fields  fields
		args    args
		want    fields
		wantErr bool
	}{
		"OrderCreatedEvent": {
			fields: fields{
				ConsumerID:   "ConsumerID-old",
				RestaurantID: "RestaurantID-old",
				LineItems:    make([]LineItem, 0),
				DeliverAt:    deliverAt,
				DeliverTo:    Address{},
			},
			want: fields{
				ConsumerID:   "ConsumerID-new",
				RestaurantID: "RestaurantID-new",
				LineItems:    make([]LineItem, 0),
				DeliverAt:    deliverAt.Add(time.Hour),
				DeliverTo:    Address{Street1: "street1"},
				Status:       ApprovalPending,
			},
			args: args{
				ddd.NewEvent(OrderCreatedEvent, &OrderCreated{
					ConsumerID:   "ConsumerID-new",
					RestaurantID: "RestaurantID-new",
					LineItems:    make([]LineItem, 0),
					DeliverAt:    deliverAt.Add(time.Hour),
					DeliverTo:    Address{Street1: "street1"},
				}),
			},
			wantErr: false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			o := &Order{
				Aggregate:    aggregate,
				ConsumerID:   tt.fields.ConsumerID,
				RestaurantID: tt.fields.RestaurantID,
				TicketID:     tt.fields.TicketID,
				LineItems:    tt.fields.LineItems,
				Status:       tt.fields.Status,
				DeliverAt:    tt.fields.DeliverAt,
				DeliverTo:    tt.fields.DeliverTo,
			}
			if err := o.ApplyEvent(tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("ApplyEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(t, o.ConsumerID, tt.want.ConsumerID)
			assert.Equal(t, o.RestaurantID, tt.want.RestaurantID)
			assert.Equal(t, o.TicketID, tt.want.TicketID)
			assert.Equal(t, o.LineItems, tt.want.LineItems)
			assert.Equal(t, o.Status, tt.want.Status)
			assert.Equal(t, o.DeliverAt, tt.want.DeliverAt)
			assert.Equal(t, o.DeliverTo, tt.want.DeliverTo)

		})
	}
}
func TestOrder_CreateOrder(t *testing.T) {
	aggregate := es.NewAggregate("id", "aggregate")
	deliverAt := time.Now()

	type args struct {
		consumerID   string
		restaurantID string
		lineItems    []LineItem
		deliverAt    time.Time
		deliverTo    Address
	}
	tests := map[string]struct {
		args        args
		orderStatus OrderStatus
		wantErr     error
	}{
		"OK": {
			args: args{
				consumerID:   "consumerID",
				restaurantID: "restaurantID",
				lineItems:    []LineItem{LineItem{}},
				deliverAt:    deliverAt,
				deliverTo:    Address{Street1: "Street1"},
			},
			orderStatus: UnknownOrderStatus,
			wantErr:     nil,
		},
		"Blank DliverTo": {
			args: args{
				consumerID:   "consumerID",
				restaurantID: "restaurantID",
				lineItems:    []LineItem{LineItem{}},
				deliverAt:    deliverAt,
				// deliverTo:    Address{Street1: "Street1"},
			},
			orderStatus: UnknownOrderStatus,
			wantErr:     ErrDeliverToCannotBeBlank,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			o := &Order{Aggregate: aggregate}
			o.Status = tt.orderStatus
			event, err := o.CreateOrder(
				tt.args.consumerID,
				tt.args.restaurantID,
				tt.args.lineItems,
				tt.args.deliverAt,
				tt.args.deliverTo,
			)
			if err != nil {
				assert.ErrorIs(t, err, tt.wantErr)
				return
			}
			assert.Nil(t, tt.wantErr)
			assert.NotNil(t, event)
			assert.Len(t, o.Events(), 1)
			assert.Equal(t, o.Events()[0].EventName(), OrderCreatedEvent)
		})
	}
}
func TestOrder_ToSnapshot(t *testing.T) {
	aggregate := es.NewAggregate("id", "aggregate")
	deliverAt := time.Now()
	type fields struct {
		ConsumerID   string
		RestaurantID string
		TicketID     string
		LineItems    []LineItem
		Status       OrderStatus
		DeliverAt    time.Time
		DeliverTo    Address
	}
	tests := map[string]struct {
		fields fields
		want   es.Snapshot
	}{
		"OK_V1": {
			fields: fields{
				ConsumerID: "id",
				LineItems:  make([]LineItem, 0),
				Status:     ApprovalPending,
				DeliverTo:  Address{Street1: "street1"},
				DeliverAt:  deliverAt,
			},
			want: &OrderV1{
				ConsumerID: "id",
				LineItems:  make([]LineItem, 0),
				Status:     ApprovalPending,
				DeliverTo:  Address{Street1: "street1"},
				DeliverAt:  deliverAt,
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			o := &Order{
				Aggregate:    aggregate,
				ConsumerID:   tt.fields.ConsumerID,
				RestaurantID: tt.fields.RestaurantID,
				TicketID:     tt.fields.TicketID,
				LineItems:    tt.fields.LineItems,
				Status:       tt.fields.Status,
				DeliverAt:    tt.fields.DeliverAt,
				DeliverTo:    tt.fields.DeliverTo,
			}

			if got := o.ToSnapshot(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToSnapshot() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestNewOrder(t *testing.T) {
	type args struct {
		id string
	}
	tests := map[string]struct {
		args args
		want *Order
	}{
		"OK": {
			args: args{id: "order-id"},
			want: &Order{
				Aggregate: es.NewAggregate("order-id", OrderAggregate),
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := NewOrder(tt.args.id)

			assert.Equal(t, tt.want.ID(), got.ID())
			assert.Equal(t, tt.want.AggregateName(), got.AggregateName())
		})
	}
}
