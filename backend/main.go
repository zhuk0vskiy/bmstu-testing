package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	"log"
	"net/http"
	"os"
	app2 "ppo/internal/app"
	"ppo/internal/config"
	iLogger "ppo/pkg/logger"
	"ppo/pkg/storage/postgres"
	"ppo/web"
)

var tokenAuth *jwtauth.JWTAuth

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	logFile, err := os.OpenFile(cfg.Logger.File, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatalln(err)
	}
	defer func(logFile *os.File) {
		err := logFile.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(logFile)
	logger := iLogger.NewLogger(cfg.Logger.Level, logFile)

	tokenAuth = jwtauth.New("HS256", []byte(cfg.Jwt.Key), nil)

	fmt.Println("connecting to database...")
	pool, err := postgres.NewConn(context.Background(), cfg)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("successfully connected to database!")

	app := app2.NewApp(pool, cfg, logger)
	mux := chi.NewRouter()
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	mux.Use(middleware.Logger)

	mux.Post("/login", web.LoginHandler(app))
	mux.Post("/signup", web.RegisterHandler(app))

	mux.Route("/users", func(r chi.Router) {
		r.Get("/{id}", web.GetUser(app))
	})

	mux.Route("/salads", func(r chi.Router) {
		r.Get("/", web.GetSalads(app))
		r.Get("/{id}/rating", web.GetSaladRating(app))
		r.Get("/{id}", web.GetSaladById(app))

		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(tokenAuth))
			r.Use(jwtauth.Authenticator(tokenAuth))
			r.Use(web.ValidateUserRoleJWT)

			r.Post("/create", web.CreateSalad(app))
			r.Patch("/{id}/update", web.UpdateSalad(app))
			r.Get("/user", web.GetUserSalads(app))
			r.Get("/userRated", web.GetUserRatedSalads(app))
			r.Delete("/{id}/delete", web.DeleteSalad(app))
		})

		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(tokenAuth))
			r.Use(jwtauth.Authenticator(tokenAuth))
			r.Use(web.ValidateAdminRoleJWT)

			r.Get("/status", web.GetSaladsWithStatus(app))
		})
	})

	mux.Route("/recipe", func(r chi.Router) {
		r.Get("/{id}", web.GetSaladRecipe(app))
		r.Get("/{id}/rating", web.GetSaladRating(app))

		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(tokenAuth))
			r.Use(jwtauth.Authenticator(tokenAuth))
			r.Use(web.ValidateUserRoleJWT)

			r.Post("/create", web.CreateRecipe(app))
			r.Patch("/{id}/update", web.UpdateRecipe(app))
		})
	})

	mux.Route("/recipeSteps", func(r chi.Router) {
		r.Get("/{id}", web.GetRecipeSteps(app))

		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(tokenAuth))
			r.Use(jwtauth.Authenticator(tokenAuth))
			r.Use(web.ValidateUserRoleJWT)

			r.Post("/create", web.CreateRecipeStep(app))
			r.Patch("/{id}/update", web.UpdateRecipeStep(app))
			r.Delete("/{id}/delete", web.DeleteRecipeStep(app))
		})
	})

	mux.Route("/ingredients", func(r chi.Router) {
		r.Get("/{id}", web.GetRecipeIngredients(app))
		r.Get("/", web.GetIngredientsByPage(app))

		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(tokenAuth))
			r.Use(jwtauth.Authenticator(tokenAuth))
			r.Use(web.ValidateUserRoleJWT)

			r.Post("/link", web.LinkIngredientToSalad(app))
			r.Patch("/unlink", web.UnlinkIngredientFromSalad(app))
		})
	})

	mux.Route("/types", func(r chi.Router) {
		r.Get("/{id}", web.GetSaladTypes(app))
		r.Get("/", web.GetSaladTypesByPage(app))

		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(tokenAuth))
			r.Use(jwtauth.Authenticator(tokenAuth))
			r.Use(web.ValidateUserRoleJWT)

			r.Post("/link", web.LinkTypeToSalad(app))
			r.Patch("/unlink", web.UnlinkTypeFromSalad(app))
		})
	})

	mux.Route("/ingredientTypes", func(r chi.Router) {
		r.Get("/{id}", web.GetIngredientType(app))
	})

	mux.Route("/measurements", func(r chi.Router) {
		r.Get("/", web.GetMeasurementByRecipe(app))
		r.Get("/all", web.GetAllMeasurements(app))
		//r.Get("/{id}", web.)
	})

	mux.Route("/comments", func(r chi.Router) {
		r.Get("/", web.GetCommentsBySalad(app))
		r.Get("/userSalad", web.GetUserComment(app))
		r.Get("/{id}", web.GetCommentById(app))

		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(tokenAuth))
			r.Use(jwtauth.Authenticator(tokenAuth))
			r.Use(web.ValidateUserRoleJWT)

			r.Post("/create", web.CreateComment(app))
			r.Patch("/{id}/update", web.UpdateComment(app))
			r.Delete("/{id}/delete", web.DeleteComment(app))
		})

	})

	fmt.Println("server was started")
	http.ListenAndServe(":8081", mux)
}
