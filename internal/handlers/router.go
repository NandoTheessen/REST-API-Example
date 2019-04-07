package handlers

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/buzzbird/go-service-helpers/pkg/output"
	"github.com/sirupsen/logrus"
)

type handlerShared struct {
	log *logrus.Logger
}

func (h *handlerShared) helloWorld(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	output.WriteResponse(ctx, h.log, w, http.StatusOK, "Hello World!")
}

// NewRouter instantiates a router and adds routes and respective handlers
func NewRouter(log *logrus.Logger) *chi.Mux {
	h := &handlerShared{
		log: log,
	}

	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Get("/", h.helloWorld)
	})

	return r
}
