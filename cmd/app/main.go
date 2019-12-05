package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/PhilLar/ToISD/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	server "github.com/PhilLar/ToISD/http"
	"github.com/go-chi/cors"
)

var port int
var db string


func init() {
	defPort := 3333
	var defDB string
	if portVar, ok := os.LookupEnv("PORT"); ok {
		if portValue, err := strconv.Atoi(portVar); err == nil {
			defPort = portValue
		}
	}
	if dbVar, ok := os.LookupEnv("DATABASE_URL"); ok {
		defDB = dbVar
	}
	flag.IntVar(&port, "port", defPort, "port to listen on")
	flag.StringVar(&db, "db", defDB, "database to connect to")
}

func main() {

	dbPsql, err := models.NewDB("postgres://store_user:store_password@localhost/store?sslmode=disable", "")
	if err != nil {
		log.Panic(err)
	}
	defer dbPsql.Close()

	env := &server.Env{Store: &models.Store{DB: dbPsql}}

	r := chi.NewRouter()
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	// r.With(middleware.AllowContentType("application/sql")).Post("/newsfeed", handlers.NewsfeedPost(feed))
	// r.With(middleware.AllowContentType("application/json")).Post("/newsfeedRIGHT", handlers.NewsfeedPost(feed))

	r.Use(cors.Handler)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)

	// without auth
	r.Post("/auth/register", env.RegisterUser)
	r.Post("/auth/login", env.AuthUser)
	// with auth
	// r.Get("/auth/logout", env.LogoutUser)
	// r.Route("/stocks", func(r chi.Router) {
	// 	r.Get("/", env.HandlerLogin().GetUserItems)
	// 	r.Get("/{ID}", env.GetSingleUserItem)
	// 	r.Post("/", env.CreateUserItem)
	// })

	http.ListenAndServe(":"+strconv.Itoa(port), r)
}
