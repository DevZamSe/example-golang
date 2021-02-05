package middleware

import (
	"ejemplo/src/config"
	"ejemplo/src/dao"
	"ejemplo/src/entities"
	"encoding/json"
	"fmt"
	"net/http"
)

func InsertDatos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var datos entities.Consultas
	_ = json.NewDecoder(r.Body).Decode(&datos)

	db, err := config.GetMySQLDB()

	fmt.Println(db)
	fmt.Println(err)

	if err != nil {
		fmt.Println("la conexión fallo")
		w.WriteHeader(http.StatusBadGateway)
		json.NewEncoder(w).Encode(err)
	} else {
		newsConsulta := dao.NewsConsulta{
			Db: db,
		}

		err := newsConsulta.InsertDatos(datos)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(err)
		} else {
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode("si se pudo")
		}
	}

}
