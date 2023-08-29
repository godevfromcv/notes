package main

import (
	"backend/database"
	"backend/handlers"
	"log"
	"net/http"
)

func main() {

	// Инициализируем базу данных
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}
	defer db.Close()

	http.HandleFunc("/register", logRequest(handlers.RegisterHandler))
	http.HandleFunc("/login", logRequest(handlers.LoginHandler))
	http.HandleFunc("/notes", logRequest(handlers.GetNotesHandler))
	http.HandleFunc("/notes/add", logRequest(handlers.AddNoteHandler))

	log.Println("Server started on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

// logRequest логирует все входящие запросы
func logRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received %s request for %s", r.Method, r.URL.Path)
		next(w, r)
	}
}
