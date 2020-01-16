package mesnajes

import (
	model "chatGolang/models"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)
//Nuevo Mensaje 
func Insertar(db *sql.DB) echo.HandlerFunc {
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
//Eliminar mensaje 
func Eliminar(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		// Use our new model to delete a task
		_, err := model.DeleteMensage(db, id)
		// Return a JSON response on success
		if err == nil {
			return c.JSON(http.StatusOK, id)
			// Handle errors
		} else {
			return err
		}
	}
}