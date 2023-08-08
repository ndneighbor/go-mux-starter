package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func DomainHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hello from your handler!")
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", DomainHandler)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://go-mux-railway-app-starter-production.up.railway.app/"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
	})

	handler := c.Handler(r)
	http.Handle("/", handler)
	// Middleware
	r.Use(loggingMiddleware)

	// Routes
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/about", aboutHandler).Methods("GET")
	r.HandleFunc("/user/{username}", userHandler).Methods("GET")

	// Start the server
	http.Handle("/", r)
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}

// Handlers
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the home page!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the about page!")
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	fmt.Fprintf(w, "Hi %s! Welcome to your user page", username)
}

// Middleware
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log request information
		logInfo := fmt.Sprintf("Logged connection from middleware %s %s %s", r.RemoteAddr, r.Method, r.URL)
		log.Println(logInfo)

		// Write the same information to the browser as well
		w.Write([]byte(logInfo + "\n"))
		next.ServeHTTP(w, r)
	})
}
