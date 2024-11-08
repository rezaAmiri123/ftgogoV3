package serdes

import (
	"encoding/json"

	"github.com/rezaAmiri123/ftgogoV3/internal/registry"
)

type JsonSerde struct {
	r registry.Registry
}

func NewJsonSerde(r registry.Registry) *JsonSerde {
	return &JsonSerde{r: r}
}

var _ registry.Serde = (*JsonSerde)(nil)

func (c JsonSerde) Register(v registry.Registrable, options ...registry.BuildOption) error {
	return registry.Register(c.r, v, c.serialize, c.deserialize, options)
}
func (c JsonSerde) RegisterKey(key string, v any, options ...registry.BuildOption) error {
	return registry.RegisterKey(c.r, key, v, c.serialize, c.deserialize, options)
}
func (c JsonSerde) RegisterFactory(key string, fn func() any, options ...registry.BuildOption) error {
	return registry.RegisterFactory(c.r, key, fn, c.serialize, c.deserialize, options)
}

func (JsonSerde) serialize(v any) ([]byte, error) {
	return json.Marshal(v)
}
func (JsonSerde) deserialize(d []byte, v any) error {
	return json.Unmarshal(d, v)
}
