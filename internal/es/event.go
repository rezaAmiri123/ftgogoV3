package es

import (
	"fmt"

	"github.com/rezaAmiri123/ftgogoV3/internal/ddd"
)

type EventApplier interface{
	ApplyEvent(ddd.Event)error
}

type EventCommiter interface{
	CommitEvents()
}

func LoadEvent(v any, event ddd.AggregateEvent)error{
	type loader interface{
		EventApplier
		VersionSetter
	}

	agg,ok := v.(loader)
	if !ok{
		return fmt.Errorf("%T does not have the methods implemented to load events",v)
	}
	if err:= agg.ApplyEvent(event);err!= nil{
		return err
	}
	agg.setVersion(event.AggregateVersion())

	return nil
}