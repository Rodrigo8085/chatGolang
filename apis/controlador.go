package controlador

import (
	
	"database/sql"
	"chatGolang/models"
	"net/http"
	"strconv"
	
	"github.com/labstack/echo"
)
//mostrar todos los usuarios
func UsuariosObtener(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, model.GetUsuarios(db))

	}
}
//mostrar al usuario
func UsuarioObtener(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
			return c.JSON(http.StatusOK, model.GetUsuario(db, id))
	}
}
// Registrar un usuario
func UsuariosGuardar(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Nueva instacia del modelo
		var usuario model.UsuarioJ
		// mapea el Json con el modelo
		c.Bind(&usuario)
		// Add a task using our new model
		guardado, err := model.PutUsuario(db, usuario.Nombre, usuario.Edad, usuario.Email, usuario.Contrasenia)
		// Return a JSON response if successful
		if err != nil {
			return c.JSON(http.StatusUnauthorized, "Error al guardar")
		} else {
			return c.JSON(http.StatusCreated, guardado)
		}
	}
}
//Actualizar al usuario
func UsuariosActualizar(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error{
		var usuario model.UsuarioJ
		// mapea el Json con el modelo
		c.Bind(&usuario)
		actualizado, err := model.UpdateUsuario(db, usuario.Nombre, usuario.Edad, usuario.Email, usuario.Contrasenia, usuario.IDUsuario)
		if err == nil {
			return c.JSON(http.StatusOK, actualizado)
		} else {
			return c.JSON(http.StatusUnauthorized, "Error al actualizar")
		}
	}
}
//***************************************contactos ******************************************
//mostrar contactos por id de usuario
func ContactosObtener(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		return c.JSON(http.StatusOK, model.GetContactos(db, id))
	}
}
//Crear un contacto por usuario 
func ContactosCrear(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Nueva instacia del modelo
		var contacto model.ContactoJ
		// mapea el Json con el modelo
		c.Bind(&contacto)
		// Add a task using our new model
		guardado, err := model.CreateContacto(db, contacto.IDUsuario, contacto.IDUsuarioContacto, contacto.Nombre, contacto.Telefono, contacto.Direccion)
		// Return a JSON response if successful
		if err != nil {
			return c.JSON(http.StatusUnauthorized, "Error al guardar")
		} else {
			return c.JSON(http.StatusCreated, guardado)
		}
	}
}
//**********************************conversacion**********************************************
//Crear una concersacion 
func ConversacionCrear(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Nueva instacia del modelo
		var conversacion model.ConversacionJ
		// mapea el Json con el modelo
		c.Bind(&conversacion)
		// Add a task using our new model
		guardado, err := model.CreateConversacion(db, conversacion.IDContacto, conversacion.NombreGrupo, conversacion.Administrador)
		// Return a JSON response if successful
		if err != nil {
			return c.JSON(http.StatusUnauthorized, "Error al guardar")
		} else {
			return c.JSON(http.StatusCreated, guardado)
		}
	}
}
func ObtenerMensajes(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		return c.JSON(http.StatusOK, model.GetMensajes(db, id))
	}
}
//*********************************mensajes*************************************************
//Nuevo Mensaje 
func InsertarMensaje(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Nueva instacia del modelo
		var mensaje model.MensajeJ
		// mapea el Json con el modelo
		c.Bind(&mensaje)
		// Add a task using our new model
		guardado, err := model.CreateMensaje(db, mensaje.IDConversacion, mensaje.Mensaje, mensaje.IDUsuario)
		// Return a JSON response if successful
		if err != nil {
			return c.JSON(http.StatusUnauthorized, "Error al guardar")
		} else {
			return c.JSON(http.StatusCreated, guardado)
		}
	}
}
//***************************grupo**************************************************
func GrupoCrear((db *sql.DB) echo.HandlerFunc {
	
}