package rest

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
	"github.com/rezaAmiri123/ftgogoV3/internal/web"
	"github.com/rezaAmiri123/ftgogoV3/store-web/internal/application"
	"github.com/rezaAmiri123/ftgogoV3/store-web/internal/application/commands"
	"github.com/rezaAmiri123/ftgogoV3/store-web/internal/domain"
	"github.com/rezaAmiri123/ftgogoV3/store-web/storeapi"
	"github.com/stackus/errors"
)

const (
	jwtAudience    = "web"
	consumerCtxKey = "consumerID"
)

type Server struct {
	storeapi.Unimplemented
	app     application.App
	jwtAuth *jwtauth.JWTAuth
	router  chi.Router
}

var _ storeapi.ServerInterface = (*Server)(nil)

func NewServer(app application.App, secret string) *Server {
	jwtAuth := jwtauth.New(jwa.HS256.String(), []byte(secret), nil)

	return &Server{
		app:     app,
		jwtAuth: jwtAuth,
		router:  chi.NewRouter(),
	}
}

func (s *Server) Mount() http.Handler {

	// Protected Routes
	s.router.Group(func(r chi.Router) {
		// JWT Session Authentication
		r.Use(
			jwtauth.Verifier(s.jwtAuth),
			jwtauth.Authenticator(s.jwtAuth),
			s.decodeClaimsIntoContext,
		)

		r.Route("/restaurants", func(r chi.Router) {
			r.Post("/", s.CreateRestaurant)
			r.Route("/{restaurantID}", func(r chi.Router) {
				r.Put("/menu", s.withRestaurantID(s.UpdateRestaurantMenu))
			})

		})

	})
	return s.router
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

func (s Server) CreateRestaurant(w http.ResponseWriter, r *http.Request) {
	var request storeapi.CreateRestaurantJSONRequestBody
	err := render.Decode(r, &request)
	if err != nil {
		render.Render(w, r, web.NewErrorResponse(err))
		return
	}

	restaurantID, err := s.app.CreateRestaurant(r.Context(), commands.CreateRestaurant{
		Name:    request.Name,
		Address: s.toAddressDomain(request.Address),
	})
	if err != nil {
		render.Render(w, r, web.NewErrorResponse(err))
		return
	}

	render.Status(r, http.StatusCreated)
	render.Respond(w, r, storeapi.RestaurantIDResponse{
		Id: restaurantID,
	})
}
func (s Server) toMenuItemDomain(req []storeapi.MenuItem) []domain.MenuItem {
	menuItems := make([]domain.MenuItem, len(req))
	for i, menuItem := range req {
		menuItems[i] = domain.MenuItem{
			ID:    menuItem.Id,
			Name:  menuItem.Name,
			Price: menuItem.Price,
		}
	}
	return menuItems
}

func (s Server) toAddressDomain(address storeapi.Address) domain.Address {
	return domain.Address{
		Street1: address.Street1,
		// TODO fix street2
		// Street2: *address.Street2,
		City:  address.City,
		State: address.State,
		Zip:   address.Zip,
	}
}
func (s Server) withRestaurantID(next func(http.ResponseWriter, *http.Request, storeapi.RestaurantID)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restaurantID, err := uuid.Parse(chi.URLParam(r, "restaurantID"))
		if err != nil {
			render.Render(w, r, web.NewErrorResponse(err))
			return
		}
		next(w, r, storeapi.RestaurantID(restaurantID))
	}
}

func (s Server) UpdateRestaurantMenu(w http.ResponseWriter, r *http.Request, restaurantID storeapi.RestaurantID) {
	var request storeapi.UpdateRestaurantMenuJSONRequestBody
	err := render.Decode(r, &request)
	if err != nil {
		render.Render(w, r, web.NewErrorResponse(err))
		return
	}

	err = s.app.UpdateRestaurantMenu(r.Context(), commands.UpdateRestaurantMenu{
		RestaurantID: restaurantID.String(),
		MenuItems: s.toMenuItemDomain(request.Menu.MenuItems),
	})
	if err != nil {
		render.Render(w, r, web.NewErrorResponse(err))
		return
	}
}

func (s Server) toMenuItemsDomain(request []storeapi.MenuItem) []domain.MenuItem {
	menuItems := make([]domain.MenuItem, len(request))
	for i, menuItem := range request {
		menuItems[i] = domain.MenuItem{
			ID:    menuItem.Id,
			Name:  menuItem.Name,
			Price: menuItem.Price,
		}
	}
	return menuItems
}
