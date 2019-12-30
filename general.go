package main

import (
	"database/sql"
	"chatGolang/apis"
	//"chatGolang/contacto"

	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db := initDB("chat.db")
	migrate(db)
	// Crea instancia de framework Echo
	e := echo.New()
	//Web page
	e.File("/", "public/index.html")
	//Apis usuarios
	e.GET("/obtener-usuarios", controlador.UsuariosObtener(db))
	e.GET("/obtener-usuario/:id", controlador.UsuarioObtener(db))
	e.POST("/guardar-usuario", controlador.UsuariosGuardar(db))
	e.PUT("/actualizar-usuario",controlador.UsuariosActualizar(db))
	//Apis contacto
	e.POST("/crear-contacto",controlador.ContactosCrear(db))
	e.GET("/obtener-contactos/:id", controlador.ContactosObtener(db))
	//Apis conversacion 
	e.POST("/crear-conversacion", controlador.ConversacionCrear(db))
	e.GET("/mensajes-conversacion/:id", controlador.ObtenerMensajes(db))
	//consulta para obtener conversaciones relacionadas 
	//Apis mensajes 
	e.POST("/nuevo-mensaje", controlador.InsertarMensaje(db))
	//Apis Grupos 
	e.POST("/nuevo-grupo", controlador.GrupoCrear(db))
	// Iniciar el servidor
	e.Start(":8000")
}
func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	// Control de errores de db
	if err != nil {
		panic(err)
	}

	// si no existen errores en la coneccion realiza la conexión a la base de datos
	if db == nil {
		panic("db nil")
	}
	return db
}

func migrate(db *sql.DB) {
	sql := `
    CREATE TABLE IF NOT EXISTS usuario(
        id_usuario INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		nombre TEXT NOT NULL,
		edad TEXT,
		email TEXT NOT NULL UNIQUE,
		contrasenia TEXT NOT NULL
	);
	CREATE TABLE IF NOT EXISTS contacto(
		id_contacto INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		id_usuario INTEGER NOT NULL,
		id_usuario_contacto INTEGER NOT NULL,
		nombre TEXT NOT NULL,
		telefono TEXT,
		direccion TEXT,
		FOREIGN KEY (id_usuario) REFERENCES usuario (id_usuario),
		FOREIGN KEY (id_usuario_contacto) REFERENCES usuario (id_usuario)
	);
	CREATE TABLE IF NOT EXISTS grupo(
		id_grupo INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		id_usuario INTEGER NOT NULL,
		nombre_grupo TEXT NOT NULL,
		FOREIGN KEY (id_usuario) REFERENCES usuario (id_usuario)
	);
	CREATE TABLE IF NOT EXISTS conversacion(
		id_conversacion INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		id_contacto INTEGER,
		nombre_grupo TEXT,
		administrador INTEGER NOT NULL, 
		fecha_creacion TEXT DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (id_contacto) REFERENCES contacto (id_contacto),
		FOREIGN KEY (nombre_grupo) REFERENCES grupo (nombre_grupo)
	);
	CREATE TABLE IF NOT EXISTS mensajes(
		id_mensaje INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		id_conversacion INTEGER NOT NULL,
		mensaje TEXT NOT NULL,
		id_usuario INTEGER NOT NULL,
		fecha_creacion TEXT DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (id_conversacion) REFERENCES conversacion (id_conversacion) 
	);
    `

	_, err := db.Exec(sql)
	// Si hay alguna falla no retorna nada
	if err != nil {
		panic(err)
	}
}
