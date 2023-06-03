package main

import (
	"fmt"
	"github.com/cosmicray001/Organising-Database-Access/models"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Use the InitDB function to initialise the global variable.
	err = models.InitDB(os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal(err)
	}
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello word!"))
	})
	mux.HandleFunc("/books", booksIndex)
	log.Println("Listening and serving on port: 8000")
	err = http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func booksIndex(w http.ResponseWriter, r *http.Request) {
	bks, err := models.AllBooks()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}
}
