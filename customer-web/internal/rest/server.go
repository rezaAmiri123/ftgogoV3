package rest

import (
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
	"github.com/stackus/errors"

	"github.com/go-chi/chi/v5"
	"github.com/rezaAmiri123/ftgogoV3/consumer/consumerpb"
	"github.com/rezaAmiri123/ftgogoV3/customer-web/customerapi"
	"github.com/rezaAmiri123/ftgogoV3/customer-web/internal/application"
)

type Server struct{
	app application.App
	customerapi.Unimplemented
}

var _ customerapi.ServerInterface = (*Server)(nil)

func NewServer(app application.App)*Server{
	return &Server{app: app}
}

func RegisterServer(server *Server, router *chi.Mux)error{
	// Public Routes
	router.Get("/register", server.RegisterConsumer)
	return nil
}

func(s Server)RegisterConsumer(w http.ResponseWriter, r *http.Request){
	var request customerapi.RegisterConsumerJSONRequestBody
	// err := render
}