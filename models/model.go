package model

import (
	"database/sql"
	"strconv"
	_ "github.com/mattn/go-sqlite3"
)
//---------------------Usuarios-----------------
// Estructura con tipos de datos del Json
type UsuarioJ struct {
	IDUsuario  int    `json:"id_usuario"`
	Nombre     string `json:"nombre"`
	Edad       string `json:"edad"`
	Email      string `json:"email"`
	Contrasenia string `json:"contrasenia"`
}

// Creamos colleción de la esntrada del Json
type usuarioCollection struct {
	UsuariosJ []UsuarioJ `json:"usuarios"`
}
// Obtener usuarios registrados
func GetUsuarios(db *sql.DB) usuarioCollection {
	sql := "SELECT * FROM usuario"
	rows, err := db.Query(sql)
	// Si por alguna razon no funcionno la consulta
	if err != nil {
		panic(err)
	}
	// cierra la conexion y limpia
	defer rows.Close()

	result := usuarioCollection{}
	for rows.Next() {
		usuarioArray := UsuarioJ{}
		err2 := rows.Scan(&usuarioArray.IDUsuario, &usuarioArray.Nombre, &usuarioArray.Edad, &usuarioArray.Email, &usuarioArray.Contrasenia)
		// sale si existe algun error
		if err2 != nil {
			panic(err2)
		}
		result.UsuariosJ = append(result.UsuariosJ, usuarioArray)
	}
	return result
}
func GetUsuario(db *sql.DB, id int) usuarioCollection {
	t := strconv.Itoa(id)
	sql := "SELECT * FROM usuario WHERE id_usuario = "+t
	rows, err := db.Query(sql)
	// Si por alguna razon no funcionno la consulta
	if err != nil {
		panic(err)
	}
	// cierra la conexion y limpia
	defer rows.Close()

	result := usuarioCollection{}
	for rows.Next() {
		usuarioArray := UsuarioJ{}
		err2 := rows.Scan(&usuarioArray.IDUsuario, &usuarioArray.Nombre, &usuarioArray.Edad, &usuarioArray.Email, &usuarioArray.Contrasenia)
		// sale si existe algun error
		if err2 != nil {
			panic(err2)
		}
		result.UsuariosJ = append(result.UsuariosJ, usuarioArray)
	}
	return result
}
func PutUsuario(db *sql.DB, nombre string, edad string, email string, contrasenia string) (int64, error) {
	sql := "INSERT INTO usuario(nombre, edad, email, contrasenia) VALUES(?,?,?,?)"

	// preparar la declaración
	stmt, err := db.Prepare(sql)
	// sale si existe algun problema
	if err != nil {
		panic(err)
	}
	//Asegurar de de limpiar cuando salga del programa
	defer stmt.Close()

	// Rellena los datos con el array
	result, err2 := stmt.Exec(nombre, edad, email, contrasenia)
	// sale si existe un error
	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()
}
func UpdateUsuario(db *sql.DB, nombre string, edad string, email string, contrasenia string, id_usuario int) (int64, error) {
	sql := "UPDATE usuario SET nombre = ?, edad = ?, email = ?, contrasenia = ? WHERE id_usuario = ?"

	// preparar la declaración
	stmt, err := db.Prepare(sql)
	// sale si existe algun problema
	if err != nil {
		panic(err)
	}
	//Asegurar de de limpiar cuando salga del programa
	defer stmt.Close()

	// Rellena los datos con el array
	result, err2 := stmt.Exec(nombre, edad, email, contrasenia,id_usuario)
	// sale si existe un error
	if err2 != nil {
		panic(err2)
	}
	return result.RowsAffected()
}
//**********************************************Contactos********************************
// Estructura con tipos de datos del Json
type ContactoJ struct {
	IDContacto  		int    	`json:"id_contacto"`
	IDUsuario   		int		`json:"id_usuario"`
	IDUsuarioContacto   int 	`json:"id_usuario_contacto"`
	Nombre      		string 	`json:"nombre"`
	Telefono 			string 	`json:"telefono"`
	Direccion 			string	`json:"direccion"`
}

// Creamos colleción de la esntrada del Json
type contactoCollection struct {
	ContactosJ []ContactoJ `json:"contactos"`
}
// Obtener usuarios registrados
func GetContactos(db *sql.DB, id int) contactoCollection {
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
