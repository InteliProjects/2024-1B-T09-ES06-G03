package api

import (
	"database/sql"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/Inteli-College/2024-1B-T09-ES06-G03/config"
	_ "github.com/Inteli-College/2024-1B-T09-ES06-G03/docs"
	"github.com/Inteli-College/2024-1B-T09-ES06-G03/service/category"
	"github.com/Inteli-College/2024-1B-T09-ES06-G03/service/geolocation"
	"github.com/Inteli-College/2024-1B-T09-ES06-G03/service/healthcheck"
	"github.com/Inteli-College/2024-1B-T09-ES06-G03/service/subcategory"
	"github.com/Inteli-College/2024-1B-T09-ES06-G03/service/user"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
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

func newReverseProxyProject(target string) *httputil.ReverseProxy {
	url, _ := url.Parse(target)
	proxy := httputil.NewSingleHostReverseProxy(url)

	proxy.ModifyResponse = func(resp *http.Response) error {
		return nil
	}

	proxy.Director = func(req *http.Request) {
		req.URL.Scheme = url.Scheme
		req.URL.Host = url.Host
		req.URL.Path = singleJoiningSlash(url.Path, req.URL.Path[len("/api/v1/projects/"):])
		req.Host = url.Host

		// Log request details
		log.Println("Proxying request to:", req.URL.String())
		log.Println("Request Headers:", req.Header)

		// Extract and verify the JWT
		authHeader := req.Header.Get("Authorization")
		if authHeader != "" {
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")

			// Decode and verify the token
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				// Make sure that the token method conforms to "alg" value.
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, jwt.ErrInvalidKey
				}
				return []byte(config.Envs.JWTSecret), nil
			})

			if err != nil || !token.Valid {
				log.Println("Invalid JWT Token")
				return
			}

			// Extract claims
			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				log.Println("JWT Claims:", claims)
			} else {
				log.Println("Invalid JWT Claims")
			}
		}
	}

	return proxy
}

func newReverseProxyCEO(target string) *httputil.ReverseProxy {
	url, _ := url.Parse(target)
	proxy := httputil.NewSingleHostReverseProxy(url)

	proxy.ModifyResponse = func(resp *http.Response) error {
		return nil
	}

	proxy.Director = func(req *http.Request) {
		req.URL.Scheme = url.Scheme
		req.URL.Host = url.Host
		req.URL.Path = singleJoiningSlash(url.Path, req.URL.Path[len("/api/v1/ceo/"):])
		req.Host = url.Host

		// Log request details
		log.Println("Proxying request to:", req.URL.String())
		log.Println("Request Headers:", req.Header)

		// Extract and verify the JWT
		authHeader := req.Header.Get("Authorization")
		if authHeader != "" {
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")

			// Decode and verify the token
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				// Make sure that the token method conforms to "alg" value.
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, jwt.ErrInvalidKey
				}
				return []byte(config.Envs.JWTSecret), nil
			})

			if err != nil || !token.Valid {
				log.Println("Invalid JWT Token")
				return
			}

			// Extract claims
			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				log.Println("JWT Claims:", claims)
			} else {
				log.Println("Invalid JWT Claims")
			}
		}
	}

	return proxy
}

func singleJoiningSlash(a, b string) string {
	if a == "" {
		return b
	}
	if b == "" {
		return a
	}
	aSlash := strings.HasSuffix(a, "/")
	bSlash := strings.HasPrefix(b, "/")
	switch {
	case aSlash && bSlash:
		return a + b[1:]
	case !aSlash && !bSlash:
		return a + "/" + b
	}
	return a + b
}

func (s *ApiServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	healthcheck := healthcheck.NewHandler()
	healthcheck.RegisterRoutes(subrouter)

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	categoryStore := category.NewStore(s.db)
	categoryHandler := category.NewHandler(categoryStore, userStore)
	categoryHandler.RegisterRoutes(subrouter)

	subcategoryStore := subcategory.NewStore(s.db)
	subcategoryHandler := subcategory.NewHandler(subcategoryStore, userStore)
	subcategoryHandler.RegisterRoutes(subrouter)

	geoHandler := geolocation.NewGeoHandler()
	geoHandler.RegisterRoutes(subrouter)

	ceoTarget := "http://localhost:8083/ceo/v1"
	ceoProxy := newReverseProxyCEO(ceoTarget)

	projectTarget := "http://localhost:8082/projects/v1"
	projectProxy := newReverseProxyProject(projectTarget)

	subrouter.PathPrefix("/ceo/healthcheck").Handler(ceoProxy)
	subrouter.PathPrefix("/ceo/notifications").Handler(ceoProxy)
	subrouter.PathPrefix("/ceo/ratings").Handler(ceoProxy)

	subrouter.PathPrefix("/projects/healthcheck").Handler(projectProxy)
	subrouter.PathPrefix("/projects/projects").Handler(projectProxy)
	subrouter.PathPrefix("/projects/updates").Handler(projectProxy)
	subrouter.PathPrefix("/projects/synergies").Handler(projectProxy)
	subrouter.PathPrefix("/projects/projects/predict").Handler(projectProxy)

	subrouter.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // You can change this to your allowed origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: true,
		Debug:            true,
	})

	c.Handler(router)

	log.Println("listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
