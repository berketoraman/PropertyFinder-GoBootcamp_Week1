package main

// Call http://localhost:8080/Berke
import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", GreetingServer)
	http.ListenAndServe(":8080", nil)
}

func GreetingServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
