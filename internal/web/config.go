package web

import (
	"fmt"
)

type WebConfig struct {
	Host string `default:"0.0.0.0"`
	Port string `default:":8000"`
}

func (c WebConfig) Address() string {
	return fmt.Sprintf("%s%s", c.Host, c.Port)
}
