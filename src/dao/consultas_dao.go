package dao

import (
	"database/sql"
	"ejemplo/src/entities"
	"fmt"
)

type NewsConsulta struct {
	Db *sql.DB
}

func (newsConsulta NewsConsulta) InsertDatos(datos entities.Consultas) error {
	result, err := newsConsulta.Db.Exec(
		"insert into test(nombre, apellido) values (?,?)",
		datos.Nombre, datos.Apellido)

	fmt.Println("si paso la query")

	fmt.Println(result)
	fmt.Println(err)

	if err != nil {
		return err
	} else {
		fmt.Println(result)
		return nil
	}

}
