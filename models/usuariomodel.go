package usuariomodel

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Estructura con tipos de datos del Json
type UsuarioJ struct {
	IDUsuario  int    `json:"id_usuario"`
	Nombre     string `json:"nombre"`
	Edad       string `json:"edad"`
	Email      string `json:"email"`
	Contrasena string `json:"password`
	PictureURL string `json:"picture_url"`
}

// Creamos colleción de la esntrada del Json
type usuarioCollection struct {
	UsuariosJ []UsuarioJ `json:"usuarios"`
}

// Obtener usuarios registrados
func GetUsuario(db *sql.DB) usuarioCollection {
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
		err2 := rows.Scan(&usuarioArray.IDUsuario, &usuarioArray.Nombre, &usuarioArray.Edad, &usuarioArray.Email, &usuarioArray.Contrasena, &usuarioArray.PictureURL)
		// sale si existe algun error
		if err2 != nil {
			panic(err2)
		}
		result.UsuariosJ = append(result.UsuariosJ, usuarioArray)
	}
	return result
}
func PutUsuario(db *sql.DB, nombre string, edad string, email string, password string, picture_url string) (int64, error) {
	sql := "INSERT INTO usuario(nombre, edad, email, password, picture_url) VALUES(?,?,?,?,?)"

	// preparar la declaración
	stmt, err := db.Prepare(sql)
	// sale si existe algun problema
	if err != nil {
		panic(err)
	}
	//Asegurar de de limpiar cuando salga del programa
	defer stmt.Close()

	// Rellena los datos con el array
	result, err2 := stmt.Exec(nombre, edad, email, password, picture_url)
	// sale si existe un error
	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()
}