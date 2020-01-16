package usuarios

import (
	model "chatGolang/models"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

//mostrar todos los usuarios
func ObtenerTodos(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, model.GetUsuarios(db))

	}
}

//mostrar al usuario
func Obtener(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		return c.JSON(http.StatusOK, model.GetUsuario(db, id))
	}
}

// Registrar un usuario
func Guardar(db *sql.DB) echo.HandlerFunc {
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
func Actualizar(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
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
