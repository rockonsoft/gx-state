package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-pg/pg/v10"
	"rockonsoft.com/gx-state-api/machine"
)

// start api with the pgdb and return a chi router
func StartAPI(pgdb *pg.DB, service *machine.MachineService) *chi.Mux {
	//get the router
	r := chi.NewRouter()
	//add middleware
	//in this case we will store our DB to use it later
	// r.Use(middleware.RequestID)
	// r.Use(middleware.Logger)
	// r.Use(middleware.Recoverer)
	// r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Use(middleware.Logger, middleware.WithValue("DB", pgdb))
	r.Use(middleware.Logger, middleware.WithValue("Service", service))

	// r.Get("/articles/{slug}", getArticle)
	// r.Get("/machines/{slug}", getMachine)

	//routes for our service
	r.Route("/api/machine-def", func(r chi.Router) {
		r.Post("/", createMachineDef)
		// r.Get("/", getMachineDefs)
		r.Get("/{machineName}", getMachineDefsByTyneName)
		// r.Put("/{machineDefID}", updateMachineDefID)
		// r.Delete("/{machineDefID}", deleteMachineDefID)
	})
	// r.Route("/api/machine-def", func(r chi.Router) {
	// 	r.Get("/{machineSlug:[a-z-]+}", getMachineDefsByTyneName)
	// })

	// routes to instantiate a state machine
	r.Route("/api/machine", func(r chi.Router) {
		r.Post("/", createPersistedMachineInstance)
		// r.Get("/", getMachineDefs)
		// r.Get("/{machineDefID}", getMachineDefID)
		// r.Put("/{machineDefID}", updateMachineDefID)
		// r.Delete("/{machineDefID}", deleteMachineDefID)
	})
	// route to send messages to a state machine
	r.Route("/api/message", func(r chi.Router) {
		r.Post("/", createMessage)
		// r.Get("/", getMachineDefs)
		// r.Get("/{machineDefID}", getMachineDefID)
		// r.Put("/{machineDefID}", updateMachineDefID)
		// r.Delete("/{machineDefID}", deleteMachineDefID)
	})

	//test route to make sure everything works
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("up and running"))
	})

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		route = strings.Replace(route, "/*/", "/", -1)
		fmt.Printf("%s %s\n", method, route)
		return nil
	}

	if err := chi.Walk(r, walkFunc); err != nil {
		fmt.Printf("Logging err: %s\n", err.Error())
	}

	return r
}

// -- handle routes
func getArticle(w http.ResponseWriter, r *http.Request) {
	// dateParam := chi.URLParam(r, "date")
	slugParam := chi.URLParam(r, "slug")

	fmt.Println(fmt.Sprintf("error fetching article %s", slugParam))
	// article, err := database.GetArticle(date, slug)

	// if err != nil {
	//   w.WriteHeader(422)
	//   w.Write([]byte(fmt.Sprintf("error fetching article %s-%s: %v", dateParam, slugParam, err)))
	//   return
	// }

	// if article == nil {
	//   w.WriteHeader(404)
	//   w.Write([]byte("article not found"))
	//   return
	// }
	// w.Write([]byte(article.Text()))
}
