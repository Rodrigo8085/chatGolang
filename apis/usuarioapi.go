package usuarioapi

import (
	
	"database/sql"
	"chatGolang/models"
	"net/http"

	"github.com/labstack/echo"
)
//mostrar usuarios
func UsuariosObtener(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, usuariomodel.GetUsuario(db))

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
			//return err
			return c.JSON(http.StatusUnauthorized, err)
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
			return err
		}
	}
}
