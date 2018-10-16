package receta

import (
	"net/http"

	. "github.com/JuanJoseGonGi/recetas-restapi/logger"
	"github.com/gorilla/mux"
)

var controller = &Controller{Repository: Repository{}}

type Ruta struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
type Rutas []Ruta

var rutas = Rutas{
	Ruta{
		"Index",
		"GET",
		"/recetas/",
		controller.Index,
	},
	Ruta{
		"Index",
		"GET",
		"/recetas",
		controller.Index,
	},
	Ruta{
		"GetReceta",
		"GET",
		"/recetas/{nombre}",
		controller.GetReceta,
	},
	Ruta{
		"AddReceta",
		"POST",
		"/recetas",
		controller.AddReceta,
	},
	Ruta{
		"UpdateReceta",
		"PUT",
		"/recetas",
		controller.UpdateReceta,
	},
	Ruta{
		"DeleteReceta",
		"DELETE",
		"/recetas/{nombre}",
		controller.DeleteReceta,
	},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, ruta := range rutas {
		var handler http.Handler
		handler = ruta.HandlerFunc
		handler = Logger(handler, ruta.Name)
		router.
			Methods(ruta.Method).
			Path(ruta.Pattern).
			Name(ruta.Name).
			Handler(handler)
	}
	return router
}
