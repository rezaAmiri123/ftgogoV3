package logging

import "github.com/rs/zerolog"

type Application struct{
	application.App
	logger zerolog.Logger
}