package controller

import (
	model "chatGolang/models/conversacion"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)
//Crear una concersacion 
func Crear(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Nueva instacia del modelo
		var conversacion model.ConversacionJ
		// mapea el Json con el modelo
		c.Bind(&conversacion)
		// Add a task using our new model
		guardado, err := model.Crear(db, conversacion.IDContacto, conversacion.NombreGrupo, conversacion.Administrador)
		// Return a JSON response if successful
		if err != nil {
			return c.JSON(http.StatusUnauthorized, "Error al guardar")
		} else {
			return c.JSON(http.StatusCreated, guardado)
		}
	}
}
func ConversacionUsuario(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		return c.JSON(http.StatusOK, model.ConversacionUsuario(db, id))
	}
}
func ConversacionGrupo(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		return c.JSON(http.StatusOK, model.ConversacionGrupos(db, id))
	}
}