package api

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/Inteli-College/2024-1B-T09-ES06-G03/ceo/docs"
	"github.com/Inteli-College/2024-1B-T09-ES06-G03/ceo/service/healthcheck"
	"github.com/Inteli-College/2024-1B-T09-ES06-G03/ceo/service/notifications"
	"github.com/Inteli-College/2024-1B-T09-ES06-G03/ceo/service/rating"
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
	subrouter := router.PathPrefix("/ceo/v1").Subrouter()

	healthcheck := healthcheck.NewHandler()
	healthcheck.RegisterRoutes(subrouter)

	ratingStore := rating.NewStore(s.db)
	rating := rating.NewHandler(ratingStore)
	rating.RegisterRoutes(subrouter)

	notificationStore := notifications.NewStore(s.db)
	notifications := notifications.NewHandler(notificationStore)
	notifications.RegisterRoutes(subrouter)

	subrouter.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	log.Println("listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
