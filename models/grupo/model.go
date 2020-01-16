package model

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)
type GrupoJ struct {
	IDGrupo		int    	`json:"id_grupo"`
	IDUsuario  	int		`json:"id_usuario"`
	NombreGrupo string 	`json:"nombre_grupo"`
}

// Creamos colleción de la esntrada del Json
type grupoCollection struct {
	GruposJ []GrupoJ `json:"grupo"`
}
func Crear(db *sql.DB, id_usuario int, nombre_grupo string) (int64, error) {
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
func Borrar(db *sql.DB, nombre_grupo string) (int64, error) {
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
func Actualizar(db *sql.DB, nombre_grupo_debe_decir string, nombre_grupo_dice string) (int64, error) {
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