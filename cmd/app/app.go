package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"go.uber.org/zap"

	_ "gestia-auth/docs"
	"gestia-auth/internal/app/gestia/handlers"

	httpSwagger "github.com/swaggo/http-swagger"
)

type App struct {
	router *chi.Mux
}

func NewApp(logger *zap.Logger) (*App, error) {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	r.Use(NewZapMiddleware(logger))

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           500, // Maximum value not ignored by any of major browsers
	}))

	authHandler := handlers.NewRootHandler()
	r.Route("/api/v1/auth", func(r chi.Router) {
		r.Post("/register", authHandler.RegisterHandler)
		r.Post("/login", authHandler.LoginHandler)
		r.Post("/refresh", authHandler.RefreshTokenHandler)
	})
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	return &App{
		router: r,
	}, nil
}

func (a *App) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, a.router)
}
