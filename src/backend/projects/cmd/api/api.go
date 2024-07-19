package api

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/Inteli-College/2024-1B-T09-ES06-G03/project/docs"
	"github.com/Inteli-College/2024-1B-T09-ES06-G03/project/service/healthcheck"
	"github.com/Inteli-College/2024-1B-T09-ES06-G03/project/service/project"
	"github.com/Inteli-College/2024-1B-T09-ES06-G03/project/service/synergy"
	update "github.com/Inteli-College/2024-1B-T09-ES06-G03/project/service/updates"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type ApiServer struct {
	addr string
	db   *sql.DB
}

func NewApiServer(addr string, db *sql.DB) *ApiServer {
	return &ApiServer{
		addr: addr,
		db:   db,
	}
}

func (s *ApiServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/projects/v1").Subrouter()

	healthcheck := healthcheck.NewHandler()
	healthcheck.RegisterRoutes(subrouter)

	projectStore := project.NewStore(s.db)
	projects := project.NewHandler(projectStore)
	projects.RegisterRoutes(subrouter)

	synergyStore := synergy.NewStore(s.db)
	synergyHandler := synergy.NewHandler(synergyStore)
	synergyHandler.RegisterRoutes(subrouter)

	updateStore := update.NewStore(s.db)
	updateHandler := update.NewHandler(updateStore)
	updateHandler.RegisterRoutes(subrouter)

	subrouter.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	log.Println("listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
