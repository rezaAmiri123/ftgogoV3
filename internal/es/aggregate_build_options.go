package es

import (
	"fmt"

	"github.com/rezaAmiri123/ftgogoV3/internal/registry"
)

type VersionSetter interface{
	SetVersion(int)
}

func SetVersion(version int)registry.BuildOption{
	return func(v any) error {
		if agg,ok := v.(VersionSetter);ok{
			agg.SetVersion(version)
			return nil
		}
		return fmt.Errorf("%T does not have the method SetVersion(int)", v)
	}
}