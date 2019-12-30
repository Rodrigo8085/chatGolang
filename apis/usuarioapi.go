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
		// Fetch tasks using our new model
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
		guardado, err := usuariomodel.PutUsuario(db, usuario.Nombre, usuario.Edad, usuario.Email, usuario.Contrasena, usuario.PictureURL)
		// Return a JSON response if successful
		if err == nil {
			return c.JSON(http.StatusCreated, guardado)
			// Handle any errors
		} else {
			return err
		}
	}
}
