package controller

import (
	model "chatGolang/models/mensaje"
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
		guardado, err := model.Crear(db, mensaje.IDConversacion, mensaje.Mensaje, mensaje.IDUsuario)
		// Return a JSON response if successful
		if err != nil {
			return c.JSON(http.StatusUnauthorized, "Error al guardar")
		} else {
			return c.JSON(http.StatusCreated, guardado)
		}
	}
}
//Eliminar mensaje 
func Borrar(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		// Use our new model to delete a task
		_, err := model.Borrar(db, id)
		// Return a JSON response on success
		if err == nil {
			return c.JSON(http.StatusOK, id)
			// Handle errors
		} else {
			return err
		}
	}
}
func ObtenerMensajes(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		return c.JSON(http.StatusOK, model.Obtener(db, id))
	}
}