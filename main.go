package main

import (
	"ejemplo/src/middleware"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	// Init Router
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/dev", middleware.HolaMundo).Methods("POST")
	r.HandleFunc("/datos", middleware.Suma).Methods("POST")
	r.HandleFunc("/datos/{id}/{token}", middleware.DatosGet).Methods("GET")

	//Consultas Api
	r.HandleFunc("/api/insert", middleware.InsertDatos).Methods("POST")

	var port string
	port = os.Getenv("PORT")
	if port == "" {
		fmt.Println("No existe la variable")
		port = "8080"
	}

	fmt.Println("ya están los puertos abiertos")

	http.ListenAndServe(":"+port, r)

}
