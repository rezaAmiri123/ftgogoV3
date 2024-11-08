package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConsumer_GetAddress(t *testing.T) {
	address := Address{Street1: "street1"}
	type fields struct {
		Addresses map[string]Address
	}
	type args struct {
		id string
	}
	type want struct {
		address Address
		err     error
	}
	tests := map[string]struct {
		fields fields
		args   args
		want   want
	}{
		"OK": {
			fields: fields{Addresses: map[string]Address{"id": address}},
			args:   args{id: "id"},
			want:   want{address: address, err: nil},
		},
		"Not Exists": {
			fields: fields{Addresses: map[string]Address{}},
			args:   args{id: "id"},
			want:   want{address: Address{}, err: ErrAddressDoesNotExist},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			c := &Consumer{Addresses: tt.fields.Addresses}
			got, err := c.GetAddress(tt.args.id)
			if tt.want.err != nil {
				assert.Error(t, err)
				assert.ErrorContains(t, err, tt.want.err.Error())
				assert.Empty(t, got)
				return
			}
			assert.Equal(t, got.Street1, tt.want.address.Street1)
		})
	}
}

func TestRegisterConsumer(t *testing.T) {
	consumer := NewConsumer("id")
	consumer.Name = "name"
	type args struct {
		id   string
		name string
	}
	type want struct {
		consumer *Consumer
		err      error
	}
	tests := map[string]struct {
		args args
		want want
	}{
		"OK": {
			args: args{id: "id", name: "name"},
			want: want{consumer: consumer, err: nil},
		},
		"ID is blank": {
			args: args{id: "", name: "name"},
			want: want{consumer: nil, err: ErrConsumerIDCannotBeBlank},
		},
		"Name is blank": {
			args: args{id: "id", name: ""},
			want: want{consumer: nil, err: ErrConsumerNameCannotBeBlank},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := RegisterConsumer(tt.args.id, tt.args.name)
			if tt.want.err != nil {
				assert.Error(t, err)
				assert.ErrorContains(t, err, tt.want.err.Error())
				assert.Nil(t, got)
				return
			}
			assert.Equal(t, got.ID, tt.want.consumer.ID)
			assert.Empty(t, got.Addresses)
		})
	}
}
