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

	return result.LastInsertId()
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
	Apodo      			string 	`json:"apodo"`
	Telefono 			string 	`json:"telefono"`
	Direccion 			string	`json:"direccion"`
}

// Creamos colleción de la esntrada del Json
type contactoCollection struct {
	ContactosJ []ContactoJ `json:"contactos"`
}
// Obtener contactos por usuario
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
//Crear un nuevo contacto 
func CreateContacto(db *sql.DB, id_usuario int, id_usuario_contacto int, nombre string,  apodo string, telefono string, direccion string) (int64, error) {
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
func UpdateContacto(db *sql.DB, id_usuario int, id_usuario_contacto int, nombre string, apodo string, telefono string, direccion string, id_contacto int) (int64, error) {
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
//*********************************Conversacion****************************************************************
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
func CreateConversacion(db *sql.DB, id_contacto int, nombre_grupo string, administrador int) (int64, error) {
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
func GetConversacionGrupos(db *sql.DB, id int) conversacionCollection {
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
func GetConversacionUsuario(db *sql.DB, id int) conversacionCollection {
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
//**********************************Mensajes ******************************************
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
func CreateMensaje(db *sql.DB, id_conversacion int, mensaje string, id_usuario int) (int64, error) {
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
func GetMensajes(db *sql.DB, id int) mensajesCollection {
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
func DeleteMensage(db *sql.DB, id int) (int64, error) {
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
//***************************************Grupos***************************************
type GrupoJ struct {
	IDGrupo		int    	`json:"id_grupo"`
	IDUsuario  	int		`json:"id_usuario"`
	NombreGrupo string 	`json:"nombre_grupo"`
}

// Creamos colleción de la esntrada del Json
type grupoCollection struct {
	GruposJ []GrupoJ `json:"grupo"`
}
func CreateGrupo(db *sql.DB, id_usuario int, nombre_grupo string) (int64, error) {
	sql := "INSERT INTO grupo( id_usuario, nombre_grupo) VALUES(?,?)"

	// preparar la declaración
	stmt, err := db.Prepare(sql)
	// sale si existe algun problema
	if err != nil {
		panic(err)
	}
	//Asegurar de de limpiar cuando salga del programa
	defer stmt.Close()

	// Rellena los datos con el array
	result, err2 := stmt.Exec(id_usuario, nombre_grupo)
	// sale si existe un error
	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()
}
func DeleteGroup(db *sql.DB, nombre_grupo string) (int64, error) {
	sql := "DELETE FROM grupo WHERE nombre_grupo = ?"

	// preparar la declaración
	stmt, err := db.Prepare(sql)
	// sale si existe algun problema
	if err != nil {
		panic(err)
	}
	//Asegurar de de limpiar cuando salga del programa
	defer stmt.Close()

	// Rellena los datos con el array
	result, err2 := stmt.Exec(nombre_grupo)
	// sale si existe un error
	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()
}
type ActualizarGrupoJ struct {
	NombreGrupoDice		 	string    	`json:"nombre_grupo_dice"`
	NombreGrupoDebeDecir  	string		`json:"nombre_grupo_debe_decir"`
}
func UpdateGroup(db *sql.DB, nombre_grupo_debe_decir string, nombre_grupo_dice string) (int64, error) {
	sql := "UPDATE grupo SET nombre_grupo = ? WHERE nombre_grupo = ?"

	// preparar la declaración
	stmt, err := db.Prepare(sql)
	// sale si existe algun problema
	if err != nil {
		panic(err)
	}
	//Asegurar de de limpiar cuando salga del programa
	defer stmt.Close()

	// Rellena los datos con el array
	result, err2 := stmt.Exec(nombre_grupo_debe_decir, nombre_grupo_dice)
	// sale si existe un error
	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()
}