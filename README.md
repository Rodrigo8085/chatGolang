# chatGolang
vacante Entropy

-Requerimientos
  -Go version 1.13.5
  -MicroFramework Echo
  -SQLite 3 para go 

Instalacion de instancias: 
go get github.com/labstack/echo
go get github.com/mattn/go-sqlite3

-Diseño de la solución:
Para el analisis de requerimientos se hizo diagrama general de la BD
https://drive.google.com/file/d/1iBbNdy2KRDIC5eyrVr5J6K1t-WNc1XNW/view?usp=sharing

A continuacion se hizo el analisis de API's a implemetadar: 
https://docs.google.com/spreadsheets/d/1u9YcTf_fUFMtIAbUd5IpBBi9TTkapnkhYB0R2UsxfT8/edit?usp=sharing

-Implementación 
  -go build general.go
  -go run general.go
  -Esta configurado en el puerto 8000 ejemplo  "localhost:8000"

-API's y JSON's Implementados: (herramienta https://insomnia.rest/)
https://drive.google.com/file/d/1v1oS4gWyIZRzqElex3ejsamBlV3DBEXW/view?usp=sharing

Mejoras:
**no se separo de la mejor forma el proyecto, en main dejo los endpoints lo cual hace mas dificil de leer el documento principal. Lo mejor hubiera sido dentro del paquete de la API hacer un archivo o bien otro paquete con las rutas del servidor aparte de separar el codigo con asteriscos (es mejor crear nuevos archivos).
Completado 
** Por otro lado no hay pruebas unitarias, ni de los modelos ni de los endpoints lo cual está dentro de las buenas practicas para hacer software

**nomenclatura de las funciones en el proyecto no es uniforme ya que hay partes en donde el codigo esta en ingles y otras en español, incluso se combinan por ejemplo en la funcion "CreateGrupo" . Eso perjudica la facilidad de entender el codigo, cosa que es fundamental cuando se trabaja en equipos.
Completado