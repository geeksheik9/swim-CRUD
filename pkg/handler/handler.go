package handler

import (
	"net/http"

	"github.com/geeksheik9/swim-CRUD/pkg/api"
	"github.com/geeksheik9/swim-CRUD/pkg/model"
	"github.com/gorilla/mux"
)

//WorkoutDatabase is the interface for the database object
type WorkoutDatabase interface {
	Ping() error
}

//WorkoutService is the implementation of a service to access sets within a database and build a workout
type WorkoutService struct {
	Version  string
	Database WorkoutDatabase
}

//Routes sets up the routes for the RESTful interface
func (s *WorkoutService) Routes(r *mux.Router) *mux.Router {
	r.HandleFunc("/ping", s.PingCheck).Methods(http.MethodGet)
	r.Handle("/health", s.healthCheck(s.Database)).Methods(http.MethodGet)

	return r
}

//PingCheck checks that the app is running and returns 200, OK, version
func (s *WorkoutService) PingCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("OK, " + s.Version))
}

func (s *WorkoutService) healthCheck(database WorkoutDatabase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dbErr := database.Ping()
		var stringDBErr string

		if dbErr != nil {
			stringDBErr = dbErr.Error()
		}

		response := model.HealthCheckResponse{
			APIVersion: s.Version,
			DBError:    stringDBErr,
		}

		if dbErr != nil {
			api.RespondWithJSON(w, http.StatusFailedDependency, response)
			return
		}

		api.RespondWithJSON(w, http.StatusOK, response)
	})
}
