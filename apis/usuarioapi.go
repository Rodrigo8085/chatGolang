package usuarioapi

import (
	
	"database/sql"
	"chatGolang/models"
	"net/http"
	"strconv"
	
	"github.com/labstack/echo"
)
//mostrar usuarios
func UsuariosObtener(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, usuariomodel.GetUsuarios(db))

	}
}
func UsuarioObtener(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		actualizado := usuariomodel.GetUsuario(db, id)
			return c.JSON(http.StatusOK, actualizado)
	}
}
// Registrar un usuario
func UsuariosGuardar(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Nueva instacia del modelo
		var usuario usuariomodel.UsuarioJ
		// mapea el Json con el modelo
		c.Bind(&usuario)
		// Add a task using our new model
		guardado, err := usuariomodel.PutUsuario(db, usuario.Nombre, usuario.Edad, usuario.Email, usuario.Contrasenia)
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
		var usuario usuariomodel.UsuarioJ
		// mapea el Json con el modelo
		c.Bind(&usuario)
		actualizado, err := usuariomodel.UpdateUsuario(db, usuario.Nombre, usuario.Edad, usuario.Email, usuario.Contrasenia, usuario.IDUsuario)
		if err == nil {
			return c.JSON(http.StatusOK, actualizado)
		} else {
			return c.JSON(http.StatusUnauthorized, "Error al actualizar")
		}
	}
}
