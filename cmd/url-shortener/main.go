package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/ydrobot/golang-url_shortener/internal/config"
	"github.com/ydrobot/golang-url_shortener/internal/http-server/handlers/redirect"
	"github.com/ydrobot/golang-url_shortener/internal/http-server/handlers/url/save"
	"github.com/ydrobot/golang-url_shortener/internal/http-server/middleware/logger"
	"github.com/ydrobot/golang-url_shortener/internal/lib/logger/sl"
	"github.com/ydrobot/golang-url_shortener/internal/storage/sqlite"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// init config: cleanenv читает из всех популярных источников конфигурацию, можно задавть required, default, env, prefix, ignoreUnknown, tag
	cfg := config.MustLoad()

	// init logger: slog - появится в стандартной библиотеке, 1.21
	log := setupLogger(cfg.Env)

	log.Info("starting url-shortener", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	// init storage: sqlite
	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		// slog.Error("failed to init storage", err)
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	// init router: chi, "chi render" - лучший роутер, но нет поддержки http2, можно использовать fasthttp
	router := chi.NewRouter()

	// middleware
	router.Use(middleware.RequestID) // добавляет RequestID в контекст
	router.Use(middleware.Logger)
	router.Use(logger.New(log))
	router.Use(middleware.Recoverer) // если в одном хендлере случилась паника, то приложение не должно падать целиком
	router.Use(middleware.URLFormat) // передает расширение curl из пути запроса и сохраняет его в контексте в виде строки под ключом (будет жесткая привязка к chi)

	router.Route("/url", func(r chi.Router) {
		r.Use(middleware.BasicAuth("url-shortener", map[string]string{
			cfg.HTTPServer.User: cfg.HTTPServer.Password,
		}))

		r.Post("/", save.New(log, storage))
		// TODO: add DELETE /url/{id}
	})

	router.Get("/{alias}", redirect.New(log, storage))

	// run server
	log.Info("starting server", slog.String("address", cfg.Address))
	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,     // время жизни на прочитать запрос
		WriteTimeout: cfg.HTTPServer.Timeout,     // время жизни на написание ответа
		IdleTimeout:  cfg.HTTPServer.IdleTimeout, // время жизни соединение сервиса
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server")
		os.Exit(1)
	}

	log.Error("server stopped")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
