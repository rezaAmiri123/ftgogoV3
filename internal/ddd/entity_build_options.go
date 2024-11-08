package ddd

import (
	"fmt"

	"github.com/rezaAmiri123/ftgogoV3/internal/registry"
)

type IDSetter interface {
	setID(id string)
}

func SetID(id string)registry.BuildOption{
	return func(v any) error {
		if e,ok := v.(IDSetter);ok{
			e.setID(id)
			return nil
		}
		return fmt.Errorf("%T does not have the method setID(string)")
	}
}

type NameSetter interface {
	setName(name string)
}

func SetName(name string)registry.BuildOption{
	return func(v any) error {
		if e,ok := v.(NameSetter);ok{
			e.setName(name)
			return nil
		}
		return fmt.Errorf("%T does not have the method setName(string)")
	}
}
