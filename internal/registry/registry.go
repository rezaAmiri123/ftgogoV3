package registry

import (
	"sync"
)

type (
	Registrable interface {
		Key() string
	}

	Serializer   func(v any) ([]byte, error)
	Deserializer func(d []byte, v any) error

	Registry interface {
		Serialize(key string, v any) ([]byte, error)
		MustSerialize(key string, v any) []byte
		Build(key string, options ...BuildOption) (any, error)
		MustBuild(key string, options ...BuildOption) any
		Deserialize(key string, data []byte, options ...BuildOption) (any, error)
		MustDeserialize(key string, data []byte, options ...BuildOption) any
		register(key string, fn func() any, s Serializer, d Deserializer, o []BuildOption) error
	}
)

type registered struct {
	factory      func() any
	serializer   Serializer
	deserializer Deserializer
	options      []BuildOption
}

var _ Registry = (*registry)(nil)

type registry struct {
	registered map[string]registered
	mu         sync.RWMutex
}

func New() *registry {
	return &registry{
		registered: make(map[string]registered),
	}
}

func (r *registry) Serialize(key string, v any) ([]byte, error) {
	reg, exists := r.registered[key]
	if !exists {
		return nil, UnregisteredKey(key)
	}
	return reg.serializer(v)
}

func (r *registry) MustSerialize(key string, v any) []byte {
	data, err := r.Serialize(key, v)
	if err != nil {
		panic(err)
	}
	return data
}

func (r *registry) Build(key string, options ...BuildOption) (any, error) {
	reg, exists := r.registered[key]
	if !exists {
		return nil, UnregisteredKey(key)
	}

	v := reg.factory()
	uos := append(r.registered[key].options, options...)

	for _, option := range uos {
		err := option(v)
		if err != nil {
			return nil, err
		}
	}
	return v, nil
}

func (r *registry) MustBuild(key string, options ...BuildOption) any {
	v, err := r.Build(key, options...)
	if err != nil {
		panic(err)
	}
	return v
}

func (r *registry) Deserialize(key string, data []byte, options ...BuildOption) (any, error) {
	v, err := r.Build(key, options...)
	if err != nil {
		return nil, err
	}

	err = r.registered[key].deserializer(data, v)
	if err != nil {
		return nil, err
	}

	return v, err
}

func (r *registry) MustDeserialize(key string, data []byte, options ...BuildOption) any {
	v, err := r.Deserialize(key, data, options...)
	if err != nil {
		panic(err)
	}
	return v
}

func (r *registry) register(key string, fn func() any, s Serializer, d Deserializer, o []BuildOption) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.registered[key]; exists {
		return AlreadyRegisteredKey(key)
	}

	r.registered[key] = registered{
		factory:      fn,
		serializer:   s,
		deserializer: d,
		options:      o,
	}

	return nil
}
