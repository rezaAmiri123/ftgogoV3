package serdes

import (
	"fmt"
	"reflect"

	"github.com/rezaAmiri123/ftgogoV3/internal/registry"
	"google.golang.org/protobuf/proto"
)

type ProtoSerde struct {
	r registry.Registry
}

var _ registry.Serde = (*ProtoSerde)(nil)
var protoT = reflect.TypeOf((*proto.Message)(nil)).Elem()

func NewProtoSerde(r registry.Registry) *ProtoSerde {
	return &ProtoSerde{r: r}
}

func (c *ProtoSerde) Register(v registry.Registrable, options ...registry.BuildOption) error {
	if !reflect.TypeOf(v).Implements(protoT) {
		return fmt.Errorf("%T does not implement proto.Message", v)
	}
	return registry.Register(c.r, v, c.serialize, c.deserialize, options)
}

func (c *ProtoSerde) RegisterKey(key string, v any, options ...registry.BuildOption) error {
	if !reflect.TypeOf(v).Implements(protoT) {
		return fmt.Errorf("%T does not implement proto.Message", v)
	}
	return registry.RegisterKey(c.r, key, v, c.serialize, c.deserialize, options)
}

func (c *ProtoSerde) RegisterFactory(key string, fn func() any, options ...registry.BuildOption) error {
	if v := fn(); v == nil {
		return fmt.Errorf("%s factory returned a nil value", key)
	} else if !reflect.TypeOf(v).Implements(protoT) {
		return fmt.Errorf("%T does not implement proto.Message", v)
	}
	return registry.RegisterFactory(c.r, key, fn, c.serialize, c.deserialize, options)
}

func (ProtoSerde) serialize(v any) ([]byte, error) {
	return proto.Marshal(v.(proto.Message))
}
func (ProtoSerde) deserialize(d []byte, v any) error {
	return proto.Unmarshal(d, v.(proto.Message))
}
