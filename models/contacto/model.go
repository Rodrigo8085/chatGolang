package model

import (
	"database/sql"
	"strconv"
	_ "github.com/mattn/go-sqlite3"
)
// Estructura con tipos de datos del Json
type ContactoJ struct {
	IDContacto  		int    	`json:"id_contacto"`
	IDUsuario   		int		`json:"id_usuario"`
	IDUsuarioContacto   int 	`json:"id_usuario_contacto"`
	Nombre      		string 	`json:"nombre"`
	Apodo      			string 	`json:"apodo"`
	Telefono 			string 	`json:"telefono"`
	Direccion 			string	`json:"direccion"`
}

// Creamos colleción de la esntrada del Json
type contactoCollection struct {
	ContactosJ []ContactoJ `json:"contactos"`
}
// Obtener contactos por usuario
func Contactos(db *sql.DB, id int) contactoCollection {
	t := strconv.Itoa(id)
	sql := "SELECT * FROM contacto WHERE id_usuario = "+t
	rows, err := db.Query(sql)
	// Si por alguna razon no funcionno la consulta
	if err != nil {
		panic(err)
	}
	// cierra la conexion y limpia
	defer rows.Close()

	result := contactoCollection{}
	for rows.Next() {
		contactoArray := ContactoJ{}
		err2 := rows.Scan(&contactoArray.IDContacto, &contactoArray.IDUsuario, &contactoArray.IDUsuarioContacto, &contactoArray.Nombre, &contactoArray.Telefono, &contactoArray.Direccion)
		// sale si existe algun error
		if err2 != nil {
			panic(err2)
		}
		result.ContactosJ = append(result.ContactosJ, contactoArray)
	}
	return result
}
//Crear un nuevo contacto 
func Crear(db *sql.DB, id_usuario int, id_usuario_contacto int, nombre string,  apodo string, telefono string, direccion string) (int64, error) {
	sql := "INSERT INTO contacto(id_usuario, id_usuario_contacto, nombre, apodo, telefono, direccion) VALUES(?,?,?,?,?,?)"

	// preparar la declaración
	stmt, err := db.Prepare(sql)
	// sale si existe algun problema
	if err != nil {
		panic(err)
	}
	//Asegurar de de limpiar cuando salga del programa
	defer stmt.Close()

	// Rellena los datos con el array
	result, err2 := stmt.Exec(id_usuario, id_usuario_contacto, nombre, apodo, telefono, direccion)
	// sale si existe un error
	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()
}
func Actualizar(db *sql.DB, id_usuario int, id_usuario_contacto int, nombre string, apodo string, telefono string, direccion string, id_contacto int) (int64, error) {
	sql := "UPDATE usuario SET id_usuario = ?, id_usuario_contacto = ?, nombre = ?, apodo = ?, telefono = ?, direccion = ? WHERE id_contacto = ?"

	// preparar la declaración
	stmt, err := db.Prepare(sql)
	// sale si existe algun problema
	if err != nil {
		panic(err)
	}
	//Asegurar de de limpiar cuando salga del programa
	defer stmt.Close()

	// Rellena los datos con el array
	result, err2 := stmt.Exec(id_usuario, id_usuario_contacto, nombre, apodo,telefono, direccion, id_contacto)
	// sale si existe un error
	if err2 != nil {
		panic(err2)
	}
	return result.RowsAffected()
}