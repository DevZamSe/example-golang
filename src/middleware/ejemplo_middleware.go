package middleware

import (
	"ejemplo/src/entities"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//ver que hacemos los servicios
func HolaMundo(rw http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")

	var world entities.HolaMundo
	_ = json.NewDecoder(r.Body).Decode(&world)

	fmt.Println(world.Dato)

}

func Suma(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	datos := entities.HolaMundo{"texto"}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(datos)
}

func DatosGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	parametros := mux.Vars(r)

	id := parametros["id"]
	token := parametros["token"]

	datos := entities.ResponseDatos{id, token}

	json.NewEncoder(w).Encode(datos)

}
