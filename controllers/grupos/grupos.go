package mesnajes

import (
	model "chatGolang/models"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)
func Crear(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Nueva instacia del modelo
		var grupo model.GrupoJ
		// mapea el Json con el modelo
		c.Bind(&grupo)
		// Add a task using our new model
		guardado, err := model.CreateGrupo(db, grupo.IDUsuario, grupo.NombreGrupo)
		// Return a JSON response if successful
		if err != nil {
			return c.JSON(http.StatusUnauthorized, "Error al guardar")
		} else {
			return c.JSON(http.StatusCreated, guardado)
		}
	}
}
func Conversacion(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		return c.JSON(http.StatusOK, model.GetConversacionGrupos(db, id))
	}
}
func Elimina(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Instantiate a new task
		var grupo model.GrupoJ
		// Map imcoming JSON body to the new Task
		c.Bind(&grupo)
		// Add a task using our new model
		respuesta, err := model.DeleteGroup(db, grupo.NombreGrupo)
		// Return a JSON response if successful
		if err == nil {
			return c.JSON(http.StatusCreated, respuesta)
			// Handle any errors
		} else {
			return err
		}
	}
}
func Editar(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Instantiate a new task
		var grupo model.ActualizarGrupoJ
		// Map imcoming JSON body to the new Task
		c.Bind(&grupo)
		// Add a task using our new model
		respuesta, err := model.UpdateGroup(db, grupo.NombreGrupoDice, grupo.NombreGrupoDebeDecir)
		// Return a JSON response if successful
		if err == nil {
			return c.JSON(http.StatusCreated, respuesta)
			// Handle any errors
		} else {
			return err
		}
	}
}