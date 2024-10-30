package rest

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/rezaAmiri123/ftgogoV3/store-web/storeapi"
)

func SwaggerHandler()http.Handler{
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		swagger,err := storeapi.GetSwagger()
		if err!= nil{
			render.Status(r,http.StatusInternalServerError)
			render.PlainText(w,r,fmt.Sprintf("error rendering swagger api: %s", err.Error()))
			return
		}
		render.JSON(w,r,swagger)
	})
	return r
}