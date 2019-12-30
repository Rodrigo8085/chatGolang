package contactomodel

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

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
func GetContactos(db *sql.DB, id_usuario int) contactoCollection {
	sql := "SELECT * FROM contacto WHERE id_usuario = ?"

	rows := db.QueryRow(sql,id_usuario)

	result := contactoCollection{}
		contactoArray := ContactoJ{}
		err2 := rows.Scan(&contactoArray.IDContacto, &contactoArray.IDUsuario, &contactoArray.IDUsuarioContacto, &contactoArray.Nombre, &contactoArray.Telefono, &contactoArray.Direccion)
		// sale si existe algun error
		if err2 != nil {
			panic(err2)
		}
		result.ContactosJ = append(result.ContactosJ, contactoArray)
	return result
}
