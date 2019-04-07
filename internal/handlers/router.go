package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/buzzbird/trading-service/pkg/api"
	"github.com/go-chi/chi"

	"github.com/sirupsen/logrus"
)

type handlerShared struct {
	log *logrus.Logger
}

func (h *handlerShared) helloWorld(w http.ResponseWriter, r *http.Request) {
	// Our logger is used to give us some confidence if the correct handlers is
	// used with our endpoint
	h.log.Log(logrus.InfoLevel, "Request received on '/'")

	// As a first step we are setting our 'Content-Type' header, since this
	// will be a REST API we won't need anything but the 'application/json' type
	w.Header().Set("Content-Type", "application/json")

	// Since JSON is not native to Go, we have to go the extra step of 'Marshalling'
	// our data - or decoding it into JSON
	// Below you can also see a simple approach to error handling, which is necessary
	// if we encounter an error, we will log this event and then exit our function
	// if this were used in an actual project, it should return a response to the user instead
	jsonData, err := json.MarshalIndent("Hello World", "", "    ")
	if err != nil {
		h.log.Log(logrus.ErrorLevel, "failed to MarshalIndent response data")
		return
	}

	// Set the HTTP status of our response
	w.WriteHeader(200)

	// And finally, send (write) the HTTP response to our user!
	_, err = w.Write(jsonData)
	if err != nil {
		h.log.Log(logrus.ErrorLevel, "failed to write response")
		return
	}

}

func writeError(log *logrus.Logger, w http.ResponseWriter, err error, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	response := &api.Error{
		Code:        status,
		Description: message,
	}

	if err != nil {
		response.Message = err.Error()
	}

	log.WithFields(logrus.Fields{
		"code":        response.Code,
		"message":     response.Message,
		"description": response.Description,
	}).Error("An error occured.")

	w.WriteHeader(status)
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
