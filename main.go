package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Middleware
	r.Use(loggingMiddleware)

	// Routes
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/about", aboutHandler).Methods("GET")
	r.HandleFunc("/user/{username}", userHandler).Methods("GET")

	// Start the server
	http.Handle("/", r)
	fmt.Println("Server started at :8080")
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
