package main

import (
	"log"
	"log/slog"
	"net/http"
	"project/config"
	h "project/handlers"
	m "project/handlers/middleware"
	"project/services/auth"
	"project/services/store"
	
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	db, err := store.NewSQLStorage(store.Sqlconfig())
	if err != nil {
		log.Fatal(err)
	}
	s := store.NewStore(db)
	//store.InitStorage(db) //SET UP YOUR DB BEFORE COMMENTING OUT
	store.DB = s
	r := chi.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, //CHANGE THIS IN PRODUCTION
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           15,
	})
	sessionStore := auth.NewCookieStore(auth.SessionOptions{
		CookiesKey: config.Envs.CookiesAuthSecret,
		MaxAge:     config.Envs.CookiesAuthAgeInSeconds,
		HttpOnly:   config.Envs.CookiesAuthIsHttpOnly,
		Secure:     config.Envs.CookiesAuthIsSecure,
	})
	authService := auth.NewAuthService(sessionStore)

	authHandler := h.New(authService)
	r.Use(c.Handler)
	r.Use(middleware.Logger)
	r.Handle("/*", public())
	r.Get("/", m.LogErr(h.HandleHome))
	
	r.Get("/auth/{provider}", m.LogErr(authHandler.HandleProviderLogin))
	r.Get("/auth/{provider}/callback", m.LogErr(authHandler.HandleAuthCallback))
	r.Get("/auth/logout/{provider}", m.LogErr(authHandler.HandleAuthLogout))
	r.Get("/login", m.LogErr(h.HandleLogin))
	listenAddr := ":" + config.Envs.Port
	slog.Info("HTTP server started", "address", "http://localhost"+listenAddr)
	if err := http.ListenAndServe(listenAddr, r); err != nil {
		log.Println(err.Error())
	}
}