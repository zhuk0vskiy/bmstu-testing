package e2e_tests

import (
	"backend/src/config"
	"backend/src/internal/app"
	"backend/src/pkg/logger"
	v1 "backend/src/web/v1"
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"net/http"
	"os"
	//"ppo/internal/storage"
)

// var tokenAuth *jwtauth.JWTAuth
var TestDbInstance *pgxpool.Pool

func RunTheApp(db *pgxpool.Pool, done chan os.Signal, ok chan struct{}) {
	ctx := context.Background()
	c, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	// Create logger
	fmt.Println(1)
	loggerFile, err := os.OpenFile(
		c.Logger.File,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer func(loggerFile *os.File) {
		err := loggerFile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(loggerFile)
	fmt.Println(1)
	l := logger.New(c.Logger.Level, loggerFile)

	//tokenAuth := jwtauth.New("HS256", []byte(c.JwtKey), nil)
	fmt.Println(c.JwtKey)
	db, err = newConn(ctx, &c.Database)
	if err != nil {
		l.Fatalf("failed to connect to database: %v", err)
	}

	a := app.NewApp(db, c, l)
	fmt.Println(1)
	mux := chi.NewMux()
	mux.Use(cors.Handler(cors.Options{

		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	//mux.Use(middleware.Logger)
	fmt.Println(2)
	mux.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Post("/login", v1.LoginHandler(a))
			r.Post("/signin", v1.SignInHandler(a))

			r.Get("/validation", v1.ValidationHandler(a))

			r.Route("/reserves", func(r chi.Router) {
				r.Post("/", v1.AddReserveHandler(a))
				r.Delete("/{id}", v1.DeleteReserveHandler(a))
			})

			r.Route("/studios", func(r chi.Router) {
				r.Post("/", v1.AddStudioHandler(a))

				r.Get("/{id}", v1.GetStudioHandler(a))
				r.Get("/{id}/rooms", v1.GetRoomsByStudioHandler(a))
				r.Get("/{id}/producers", v1.GetProducerHandler(a))
				r.Get("/{id}/instrumentalists", v1.GetInstrumentalistHandler(a))
				r.Get("/{id}/equipments", v1.GetEquipmentHandler(a))

				r.Patch("/{id}", v1.UpdateStudioHandler(a))
				r.Delete("/{id}", v1.DeleteStudioHandler(a))
				r.Post("/", v1.AddStudioHandler(a))

			})

			r.Route("/rooms", func(r chi.Router) {

				r.Get("/{id}", v1.GetRoomHandler(a))

				r.Post("/", v1.AddRoomHandler(a))
				r.Patch("/{id}", v1.UpdateRoomHandler(a))
				r.Delete("/{id}", v1.DeleteRoomHandler(a))

			})

			r.Route("/producers", func(r chi.Router) {
				r.Group(func(r chi.Router) {

					r.Get("/{id}", v1.GetProducerHandler(a))
				})

				r.Group(func(r chi.Router) {

					r.Post("/", v1.AddProducerHandler(a))
					r.Patch("/{id}", v1.UpdateProducerHandler(a))
					r.Delete("/{id}", v1.DeleteProducerHandler(a))
				})
			})

			r.Route("/instrumentalists", func(r chi.Router) {
				r.Group(func(r chi.Router) {

					r.Get("/{id}", v1.GetInstrumentalistHandler(a))
				})
				r.Group(func(r chi.Router) {

					r.Post("/", v1.AddInstrumentalistHandler(a))
					r.Patch("/{id}", v1.UpdateInstrumentalistHandler(a))
					r.Delete("/{id}", v1.DeleteInstrumentalistHandler(a))
				})
			})

			r.Route("/equipments", func(r chi.Router) {
				r.Group(func(r chi.Router) {

					r.Get("/{id}", v1.GetEquipmentHandler(a))
				})
				r.Group(func(r chi.Router) {

					r.Post("/", v1.AddEquipmentHandler(a))
					r.Patch("/{id}", v1.UpdateEquipmentHandler(a))
					r.Delete("/{id}", v1.DeleteEquipmentHandler(a))
				})
			})

			r.Route("/users", func(r chi.Router) {
				r.Group(func(r chi.Router) {
					r.Get("/{id}/reserves", v1.GetUserReservesHandler(a))
				})
			})
		})
	})

	go func() {
		//serverAddress := fmt.Sprintf("%s:8083", cfg.Server.ServerHost)
		//fmt.Printf("сервер прослушивает адрес: %s\n", serverAddress)
		//logger.Infof("сервер прослушивает адрес: %s\n", serverAddress)

		ok <- struct{}{}
		fmt.Println("Len", len(ok))
		http.ListenAndServe(":8081", mux)

	}()

	//<-done
}

func newConn(ctx context.Context, cfg *config.DatabaseConfig) (pool *pgxpool.Pool, err error) {
	connStr := fmt.Sprintf("%s://%s:%s@%s:%d/%s",
		cfg.Postgres.Driver,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.Database,
	)

	pool, err = pgxpool.New(ctx, connStr)
	if err != nil {
		return nil, fmt.Errorf("подключение к БД: %w", err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("пинг БД: %w", err)
	}

	return pool, nil
}
