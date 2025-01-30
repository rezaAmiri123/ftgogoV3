package ddd

import (
	"fmt"

	"github.com/rezaAmiri123/ftgogoV3/internal/registry"
)

type IDSetter interface {
	SetID(id string)
}

func SetID(id string)registry.BuildOption{
	return func(v any) error {
		if e,ok := v.(IDSetter);ok{
			e.SetID(id)
			return nil
		}
		return fmt.Errorf("%T does not have the method setID(string)")
	}
}

type NameSetter interface {
	SetName(name string)
}

func SetName(name string)registry.BuildOption{
	return func(v any) error {
		if e,ok := v.(NameSetter);ok{
			e.SetName(name)
			return nil
		}
		return fmt.Errorf("%T does not have the method setName(string)")
	}
}
