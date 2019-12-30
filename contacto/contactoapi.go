package contactoapi

import (
	
	"database/sql"
	"chatGolang/contactos"
	"net/http"
	
	"github.com/labstack/echo"
)
//mostrar usuarios
func UsuariosObtener(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var usuario contactomodel.ContactoJ
		// mapea el Json con el modelo
		c.Bind(&usuario)
		actualizado := contactomodel.GetContactos(db, usuario.IDUsuario)
			return c.JSON(http.StatusOK, actualizado)
	}
}
