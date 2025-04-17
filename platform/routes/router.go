package routes

import (
	"auth0_example/platform/authenticator"
	"auth0_example/platform/routes/handlers"
	"encoding/gob"
	"net/http"
)

// New registers the routes and returns the router
func New(auth *authenticator.Authenticator) http.Handler {
	// Register custom types for cookie encoding
	gob.Register(map[string]interface{}{})

	mux := http.NewServeMux()

	// Server static fiels
	staticFs := http.FileServer(http.Dir("web/static"))
	mux.Handle("/public/", http.StripPrefix("/public/", staticFs))

	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/login", handlers.Login(auth))

	return mux
}
