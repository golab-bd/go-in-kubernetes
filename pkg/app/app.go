package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Run - starts the application with necessary bootstrapping
func Run() {
	fmt.Println("Waiting for request to serve...")

	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/users", getUsers).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", myRouter))

}

func getUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I am responding to your API call")
}
