package api

import (
	"net/http"
	"strconv"
	"time"
	"zerotoherochallenge/config"

	"github.com/go-chi/chi"

	"go.uber.org/zap"

	"github.com/go-chi/chi/v5/middleware"
)

// HTTPServer Holds HttpServer configuration
type HTTPServer struct {
	Logger *zap.SugaredLogger
	sc     config.ServerConfigurations
	Router *chi.Mux
}

// NewHTTPServer Initializes a new http server
func NewHTTPServer(logger *zap.SugaredLogger, serverConf config.ServerConfigurations) *HTTPServer {
	router := chi.NewRouter()

	// A good base middleware stack
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Use(middleware.AllowContentType("application/json", "multipart/form-data"))

	// Set a timeout value on the request models (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	router.Use(middleware.Timeout(60 * time.Second))

	return &HTTPServer{
		Logger: logger,
		sc:     serverConf,
		Router: router,
	}
}

// Start Fires the http server
func (r *HTTPServer) Start() {
	listeningAddr := ":" + strconv.Itoa(r.sc.Port)
	r.Logger.Infof("Server listening on port %s", listeningAddr)

	err := http.ListenAndServe(listeningAddr, r.Router)
	if err != nil {
		r.Logger.Fatalf("Failed to start http server. %v", err)
	}
}
