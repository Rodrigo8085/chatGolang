package endpoints 

import (
	contactos "chatGolang/controllers/contactos"
	usuarios "chatGolang/controllers/usuarios"
	conversacion "chatGolang/controllers/conversacion"
	mensaje "chatGolang/controllers/mensajes"
	grupos "chatGolang/controllers/grupos"
	"database/sql"

	"github.com/labstack/echo"
)
func Endpoints(db *sql.DB){
	// Crea instancia de framework Echo
	e := echo.New()
	//Web page
	e.File("/", "public/index.html")
	//Apis usuarios
	e.GET("/obtener-usuarios", usuarios.ObtenerTodos(db))
	e.GET("/obtener-usuario/:id", usuarios.Obtener(db))
	e.POST("/guardar-usuario", usuarios.Guardar(db))
	e.PUT("/actualizar-usuario", usuarios.Actualizar(db))
	//Apis contacto
	e.POST("/crear-contacto", contactos.Crear(db))
	e.PUT("/actualizar-contacto", contactos.Actualizar(db))
	e.GET("/obtener-contactos/:id", contactos.Obtener(db))
	//Apis conversacion
	e.POST("/crear-conversacion", conversacion.Crear(db))
	e.GET("/consultar-conversacion-contactos/:id", conversacion.ConversacionUsuario(db))
	e.GET("/consultar-conversacion-grupos/:id", conversacion.ConversacionGrupo(db))
	//Apis mensajes
	e.POST("/nuevo-mensaje", mensaje.Insertar(db))
	e.DELETE("eliminar-mensaje/:id", mensaje.Borrar(db))
	e.GET("/mensajes-conversacion/:id", mensaje.ObtenerMensajes(db))
	//Apis Grupos
	e.POST("/nuevo-usuario-grupo", grupos.Crear(db))
	e.PUT("/actualizar-grupo", grupos.Editar(db))
	e.DELETE("/borrar-grupo", grupos.Elimina(db))
	//imagen usuario
	/*
	e.POST("/guardar-imagen", controlador.upload(db))
	e.GET("/obtener-imagen/:id", controlador.getfoto(db))
	*/
	// Iniciar el servidor
	e.Start(":8000")
}