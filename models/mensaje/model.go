package model

import (
	"database/sql"
	"strconv"
	_ "github.com/mattn/go-sqlite3"
)
type MensajeJ struct {
	IDMensaje		int    	`json:"id_mensaje"`
	IDConversacion  int		`json:"id_conversacion"`
	Mensaje   		string 	`json:"mensaje"`
	IDUsuario      	int 	`json:"id_usuario"`
	FechaCreacion 	string 	`json:"fecha_creacion"`
}

// Creamos colleción de la esntrada del Json
type mensajesCollection struct {
	MensajesJ []MensajeJ `json:"mensajes"`
}
//Crear una conversación 
func Crear(db *sql.DB, id_conversacion int, mensaje string, id_usuario int) (int64, error) {
	sql := "INSERT INTO mensajes( id_conversacion, mensaje, id_usuario) VALUES(?,?,?)"

	// preparar la declaración
	stmt, err := db.Prepare(sql)
	// sale si existe algun problema
	if err != nil {
		panic(err)
	}
	//Asegurar de de limpiar cuando salga del programa
	defer stmt.Close()

	// Rellena los datos con el array
	result, err2 := stmt.Exec(id_conversacion, mensaje, id_usuario)
	// sale si existe un error
	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()
}
func Obtener(db *sql.DB, id int) mensajesCollection {
	t := strconv.Itoa(id)
	sql := "SELECT * FROM mensajes WHERE id_conversacion = "+t+" ORDER BY fecha_creacion ASC"
	rows, err := db.Query(sql)
	// Si por alguna razon no funcionno la consulta
	if err != nil {
		panic(err)
	}
	// cierra la conexion y limpia
	defer rows.Close()

	result := mensajesCollection{}
	for rows.Next() {
		mensajesArray := MensajeJ{}
		err2 := rows.Scan(&mensajesArray.IDMensaje, &mensajesArray.IDConversacion, &mensajesArray.Mensaje, &mensajesArray.IDUsuario, &mensajesArray.FechaCreacion)
		// sale si existe algun error
		if err2 != nil {
			panic(err2)
		}
		result.MensajesJ = append(result.MensajesJ, mensajesArray)
	}
	return result
}
func Borrar(db *sql.DB, id int) (int64, error) {
	sql := "DELETE FROM mensajes WHERE id_mensaje = ?"

	// Create a prepared SQL statement
	stmt, err := db.Prepare(sql)
	// Exit if we get an error
	if err != nil {
		panic(err)
	}

	// Replace the '?' in our prepared statement with 'id'
	result, err2 := stmt.Exec(id)
	// Exit if we get an error
	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()
}