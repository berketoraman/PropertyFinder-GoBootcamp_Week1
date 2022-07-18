package main

// Call http://localhost:8080/greet/Berke
import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/greet/{name}", greet).Methods(http.MethodGet)
	http.ListenAndServe(":8080", r)
}

func greet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	var name = vars["name"]
	w.Write([]byte("Hello " + name + "!"))
}
