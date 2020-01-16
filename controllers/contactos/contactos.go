package contactos

import (
	model "chatGolang/models"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

//mostrar contactos por id de usuario
func Obtener(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		return c.JSON(http.StatusOK, model.GetContactos(db, id))
	}
}

//Crear un contacto por usuario
func Crear(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Nueva instacia del modelo
		var contacto model.ContactoJ
		// mapea el Json con el modelo
		c.Bind(&contacto)
		// Add a task using our new model
		guardado, err := model.CreateContacto(db, contacto.IDUsuario, contacto.IDUsuarioContacto, contacto.Nombre, contacto.Apodo, contacto.Telefono, contacto.Direccion)
		// Return a JSON response if successful
		if err != nil {
			return c.JSON(http.StatusUnauthorized, "Error al guardar")
		} else {
			return c.JSON(http.StatusCreated, guardado)
		}
	}
}

//Contacto actualizar
func Actualizar(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var contacto model.ContactoJ
		// mapea el Json con el modelo
		c.Bind(&contacto)
		actualizado, err := model.UpdateContacto(db, contacto.IDUsuario, contacto.IDUsuarioContacto, contacto.Nombre, contacto.Apodo, contacto.Telefono, contacto.Direccion, contacto.IDContacto)
		if err == nil {
			return c.JSON(http.StatusOK, actualizado)
		} else {
			return c.JSON(http.StatusUnauthorized, "Error al actualizar")
		}
	}
}
