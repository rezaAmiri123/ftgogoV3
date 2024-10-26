package rest

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
	"github.com/rezaAmiri123/ftgogoV3/customer-web/customerapi"
	"github.com/rezaAmiri123/ftgogoV3/customer-web/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/customer-web/internal/application/commands"
	"github.com/rezaAmiri123/ftgogoV3/customer-web/internal/application/queries"
	"github.com/rezaAmiri123/ftgogoV3/customer-web/internal/domain"
	"github.com/rezaAmiri123/ftgogoV3/internal/web"
	"github.com/stackus/errors"
)

const (
	jwtAudience    = "web"
	consumerCtxKey = "consumerID"
	jwtExpire      = time.Hour * 24
)

type Server struct {
	customerapi.Unimplemented
	app     application.App
	jwtAuth *jwtauth.JWTAuth
	router  chi.Router
}

var _ customerapi.ServerInterface = (*Server)(nil)

func NewServer(app application.App, secret string) *Server {
	jwtAuth := jwtauth.New(jwa.HS256.String(), []byte(secret), nil)

	return &Server{
		app:     app,
		jwtAuth: jwtAuth,
		router:  chi.NewRouter(),
	}
}

func (s *Server) Mount() http.Handler {
	// Public Routes
	s.router.Post("/signin", s.SignInConsumer)
	s.router.Post("/register", s.RegisterConsumer)

	// Protected Routes
	s.router.Group(func(r chi.Router) {
		// JWT Session Authentication
		r.Use(
			jwtauth.Verifier(s.jwtAuth),
			jwtauth.Authenticator(s.jwtAuth),
			s.decodeClaimsIntoContext,
		)

		r.Route("/consumer", func(r chi.Router) {
			r.Get("/", s.GetConsumer)
		})

		r.Route("/addresses", func(r chi.Router) {
			r.Post("/", s.AddConsumerAddress)
		})

		r.Route("/orders", func(r chi.Router) {
			r.Post("/", s.CreateOrder)
		})
	})
	return s.router
}

func (s Server) RegisterConsumer(w http.ResponseWriter, r *http.Request) {
	var request customerapi.RegisterConsumerJSONRequestBody
	err := render.Decode(r, &request)
	if err != nil {
		render.Render(w, r, web.NewErrorResponse(err))
		return
	}

	consumerID, err := s.app.RegisterConsumer(r.Context(), commands.RegisterConsumer{
		Name: request.Name,
	})
	if err != nil {
		render.Render(w, r, web.NewErrorResponse(err))
		return
	}

	render.Status(r, http.StatusCreated)
	render.Respond(w, r, customerapi.ConsumerIDResponse{
		Id: consumerID,
	})
}

func (s Server) SignInConsumer(w http.ResponseWriter, r *http.Request) {
	request := customerapi.SignInConsumerJSONRequestBody{}
	err := render.Decode(r, &request)
	if err != nil {
		render.Render(w, r, web.NewErrorResponse(err))
		return
	}

	// For Demo: "login" with any known consumerID
	consumer, err := s.app.GetConsumer(r.Context(), queries.GetConsumer{
		ConsumerID: request.ConsumerId,
	})
	if err != nil {
		render.Render(w, r, web.NewErrorResponse(err))
		return
	}

	var token string
	_, token, err = s.jwtAuth.Encode(map[string]interface{}{
		jwt.SubjectKey:    consumer.ConsumerID,
		jwt.AudienceKey:   jwtAudience,
		jwt.ExpirationKey: jwtauth.ExpireIn(jwtExpire),
	})
	if err != nil {
		render.Render(w, r, web.NewErrorResponse(errors.Wrap(errors.ErrInternal, "could not authenticate you at this time")))
		return
	}

	render.Respond(w, r, customerapi.SignInResponse{
		Token: token,
	})
}

func (s Server) decodeClaimsIntoContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, claims, err := jwtauth.FromContext(r.Context())
		if err != nil {
			render.Render(w, r, web.NewErrorResponse(err))
			return
		}
		var consumerID string
		if subject, exists := claims[jwt.SubjectKey]; !exists {
			render.Render(w, r, web.NewErrorResponse(errors.Wrap(errors.ErrUnauthenticated, "missing claims subject")))
			return
		} else {
			switch s := subject.(type) {
			case string:
				consumerID = s
			case []string:
				if len(s) == 0 {
					render.Render(w, r, web.NewErrorResponse(errors.Wrap(errors.ErrUnauthenticated, "invalid claims subject")))
					return
				}
				consumerID = s[0]
			default:
				render.Render(w, r, web.NewErrorResponse(errors.Wrap(errors.ErrUnauthenticated, "invalid claims subject")))
				return
			}
		}
		ctx := context.WithValue(r.Context(), consumerCtxKey, consumerID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Setting and fetching the consumerID into and from claims is strictly the responsibility of these web handlers!
//
// The application (commands, queries), the domain, and the adapters don't know and should not know where any
// session values might come from.
func (s Server) cosnumerID(ctx context.Context) string {
	v := ctx.Value(consumerCtxKey)
	switch c := v.(type) {
	case string:
		return c
	default:
		return ""
	}
}

func (s Server) GetConsumer(w http.ResponseWriter, r *http.Request) {
	consumer, err := s.app.GetConsumer(r.Context(), queries.GetConsumer{
		ConsumerID: s.cosnumerID(r.Context()),
	})
	if err != nil {
		render.Render(w, r, web.NewErrorResponse(err))
		return
	}

	render.Respond(w, r, customerapi.ConsumerResponse{
		ConsumerId: consumer.ConsumerID,
		Name:       consumer.Name,
	})
}

func (s Server) AddConsumerAddress(w http.ResponseWriter, r *http.Request) {
	var request customerapi.AddConsumerAddressJSONRequestBody
	err := render.Decode(r, &request)
	if err != nil {
		render.Render(w, r, web.NewErrorResponse(err))
		return
	}

	consumerID := s.cosnumerID(r.Context())
	err = s.app.AddConsumerAddress(r.Context(), commands.AddConsumerAddress{
		ConsumerID: consumerID,
		AddressID:  request.Name,
		Address:    s.toAddressDomain(request.Address),
	})
	if err != nil {
		render.Render(w, r, web.NewErrorResponse(err))
		return
	}

	render.Status(r, http.StatusCreated)
	render.Respond(w, r, customerapi.ConsumerAddressIDResponse{
		AddressId:  request.Name,
		ConsumerId: consumerID,
	})
}

func (s Server) toAddressDomain(address customerapi.Address) domain.Address {
	return domain.Address{
		Street1: address.Street1,
		Street2: *address.Street2,
		City:    address.City,
		State:   address.State,
		Zip:     address.Zip,
	}
}

func (s Server) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var request customerapi.CreateOrderJSONRequestBody
	err := render.Decode(r, &request)
	if err != nil {
		render.Render(w, r, web.NewErrorResponse(err))
		return
	}

	consumerID := s.cosnumerID(r.Context())
	orderID, err := s.app.CreateOrder(r.Context(),commands.CreateOrder{
		ConsumerID: consumerID,
		RestaurantID: request.RestaurantId,
		AddressID: request.AddressId,
		LineItems: domain.MenuItemQuantities(request.LineItems),
	})
	if err != nil {
		render.Render(w, r, web.NewErrorResponse(err))
		return
	}

	render.Status(r,http.StatusCreated)
	render.Respond(w,r,customerapi.OrderIDResponse{
		Id: orderID,
	})
}