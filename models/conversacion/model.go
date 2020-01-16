package model

import (
	"database/sql"
	"strconv"
	_ "github.com/mattn/go-sqlite3"
)
type ConversacionJ struct {
	IDConversacion		int    	`json:"id_conversacion"`
	IDContacto   		int		`json:"id_contacto"`
	NombreGrupo   		string 	`json:"nombre_grupo"`
	Administrador      	int 	`json:"administrador"`
	FechaCreacion 		string 	`json:"fecha_creacion"`
}

// Creamos colleción de la esntrada del Json
type conversacionCollection struct {
	ConversacionesJ []ConversacionJ `json:"conversacion"`
}
//Crear una conversación 
func Crear(db *sql.DB, id_contacto int, nombre_grupo string, administrador int) (int64, error) {
	sql := "INSERT INTO conversacion( id_contacto, nombre_grupo, administrador) VALUES(?,?,?)"

	// preparar la declaración
	stmt, err := db.Prepare(sql)
	// sale si existe algun problema
	if err != nil {
		panic(err)
	}
	//Asegurar de de limpiar cuando salga del programa
	defer stmt.Close()

	// Rellena los datos con el array
	result, err2 := stmt.Exec(id_contacto, nombre_grupo, administrador)
	// sale si existe un error
	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()
}
func ConversacionGrupos(db *sql.DB, id int) conversacionCollection {
	t := strconv.Itoa(id)
	sql := "SELECT a.id_conversacion,a.id_contacto,a.nombre_grupo,a.administrador,a.fecha_creacion FROM conversacion a LEFT JOIN grupo b ON a.nombre_grupo = b.nombre_grupo WHERE b.id_usuario = "+t+" ORDER BY fecha_creacion ASC"
	rows, err := db.Query(sql)
	// Si por alguna razon no funcionno la consulta
	if err != nil {
		panic(err)
	}
	// cierra la conexion y limpia
	defer rows.Close()

	result := conversacionCollection{}
	for rows.Next() {
		conversacion := ConversacionJ{}
		err2 := rows.Scan(&conversacion.IDConversacion, &conversacion.IDContacto, &conversacion.NombreGrupo, &conversacion.Administrador, &conversacion.FechaCreacion)
		// sale si existe algun error
		if err2 != nil {
			panic(err2)
		}
		result.ConversacionesJ = append(result.ConversacionesJ, conversacion)
	}
	return result
}
func ConversacionUsuario(db *sql.DB, id int) conversacionCollection {
	t := strconv.Itoa(id)
	sql := "SELECT a.id_conversacion,a.id_contacto,a.nombre_grupo,a.administrador,a.fecha_creacion FROM conversacion a LEFT JOIN contacto b ON a.id_contacto = b.id_contacto WHERE b.id_usuario = "+t+" OR b.id_usuario_contacto = "+t+" ORDER BY fecha_creacion ASC"
	rows, err := db.Query(sql)
	// Si por alguna razon no funcionno la consulta
	if err != nil {
		panic(err)
	}
	// cierra la conexion y limpia
	defer rows.Close()

	result := conversacionCollection{}
	for rows.Next() {
		conversacion := ConversacionJ{}
		err2 := rows.Scan(&conversacion.IDConversacion, &conversacion.IDContacto, &conversacion.NombreGrupo, &conversacion.Administrador, &conversacion.FechaCreacion)
		// sale si existe algun error
		if err2 != nil {
			panic(err2)
		}
		result.ConversacionesJ = append(result.ConversacionesJ, conversacion)
	}
	return result
}