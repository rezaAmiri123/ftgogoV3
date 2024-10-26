package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConsumer_GetAddress(t *testing.T) {
	type fields struct {
		ID        string
		Name      string
		Addresses map[string]Address
	}
	type args struct{
		id string
	}
	type want struct{
		Address Address
		err error
	}
	tests := map[string]struct{
		fields fields
		args args
	}{}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			c := &Consumer{
				ID: tt.fields.ID,
				Name: tt.fields.Name,
				Addresses: tt.fields.Addresses,
			}
			got, err := c.GetAddress(tt.args.id)
			
		})
	}
}
