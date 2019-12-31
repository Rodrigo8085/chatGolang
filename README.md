# chatGolang
vacante Entropy

-Instancias requeridas:
"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"

-Diseño de la solución:
https://drive.google.com/file/d/1iBbNdy2KRDIC5eyrVr5J6K1t-WNc1XNW/view?usp=sharing

-API's implemetadas: 
https://docs.google.com/spreadsheets/d/1u9YcTf_fUFMtIAbUd5IpBBi9TTkapnkhYB0R2UsxfT8/edit?usp=sharing



Descripción del problema 
 
 
Como sabrás, las redes sociales están muy de moda, es por ello que los habitantes de                Entropyland quieren estar siempre en contacto, actualmente no existe ningún servicio           de mensajería decente y por esa razón pidieron tu ayuda, porque saben que eres uno               de los mejores programadores de la ciudad. 
 
La forma en cómo vas a ayudar a los habitantes de Entropyland es haciendo una API                Rest en donde los usuarios puedan realizar una serie de acciones que les permitirá              comunicarse con sus amigos y familiares de una forma mucho más sencilla.  
 
 
A continuación te mostramos todas las acciones que pueden hacer los usuarios de este              nuevo servicio de mensajería. 
 
 
 
● Registrarse 
 
Como parte inicial, los usuarios deben de registrarse. La información que deben de             llenar es su nombre, edad,email (el cual debe de ser único), foto de perfil y su                password para posteriormente autenticarse. 
 ● Autenticarse 
 
Una vez que ya tengan un usuario creado, los usuarios podrán autenticarse con su              email y password, este paso es demasiado importante ya que ​sino se encuentran             autenticados no podrán hacer ninguna de las demás operaciones​. 
 
 
● Añadir Contacto 
 
Un usuario podrá añadir contactos (uno a la vez). En donde debe de llenar la siguiente                información : Nombre (requerido), Apodo, email (Requerido), numero de telefono y           direccion.  
 
 
● Ver uno o mas Contactos 
 
En cualquier momento podrán consultar su lista de contactos o ver un contacto en              específico. 
 
 
● Crear una nueva conversación 
 
Un usuario puede iniciar una conversación con cualquiera de sus contactos. 
 ● Ver una o más  conversaciones 
 
Un usuario podrá ver una o más conversaciones las cuales a creado. 
 
 
● Enviar mensaje 
 
Puede enviar un mensaje a cualquiera de sus conversaciones creadas, el cuerpo del             mensaje y informacion adicional como la fecha y hora en que fueron enviados deben              ser requeridas.  
 
 
 
● Listar Mensajes 
 
Puede listar los mensajes de cualquiera de sus conversaciones. 
 ● Eliminar mensajes  
 
Los usuarios pueden eliminar mensajes, cabe destacar que estos en caso de ser             borrados , deben ser eliminados de forma bilateral, es decir que tanto quien envía el               mensaje como a quien recibió el mensaje ya no podrán verlo una vez que este fue                borrado. 
 
 
● Editar su información de perfil  
 
Todos los campos de su información de perfil puede ser modificada. 
 ● Crear un Grupo 
 
Un usuario puede crear un grupo, en donde dos o más personas pueden escribir              mensajes. Para la creación del grupo se necesita el nombre y al menos una persona,               además de la fecha de creación. El usuario que crea el grupo en automático se               convierte en administrador del grupo. 
 
 
● Enviar mensaje a un grupo 
 
Un usuario puede enviar un mensaje a cualquiera de los grupos de los que es               miembro. 
 ● Consultar grupos 
 
Un usuario podrá consultar en qué grupos se encuentra, o ver específicamente uno. 
 
 
● Editar Grupo 
 
Las acciones de edición que un usuario puede hacer sobre un grupo, son: agregar              administrador (puede ser un usuario que ya se encuentre en el grupo) o cambiar de               nombre al grupo, solo los usuarios con el rol de administrador en ese grupo pueden               hacer estos cambios. 
 
 
 
● Eliminar grupo 
 
Solo si el usuario es administrador del grupo puede eliminar un grupo. 
 
 
 
Consideraciones 
 
Los datos deben der ser persistentes , lo puedes hacer en memoria o preferentemente              en disco, es decir usar una instancia de una base de datos con la cual te sientas                 cómodo, por su puesto lo puedes complementar con algún ORM para hacer las             operaciones más sencillas , aunque esto no es estrictamente necesario.  
 
Aunque la información debe ser persistente el API debe ser stateless. 
 
 
Evaluación 
 
 
Se Evaluaran dos aspectos. 
 Diseño 
 
Se evaluará el diseño de tu API , así como el código , esto quiere decir que será de suma importancia que apliques buenas prácticas en tu código y el diseño de tu API. Cabe resaltar que aunque no está mencionado en los enunciados anteriores , tu API debe de mostrar errores cuando existan (por ejemplo se crean dos usuarios con el mismo email) y la explicación de forma explícita de cada uno de ellos. 
 
Performance 
 
Otro aspecto a evaluar es que tu API siempre traiga los resultados correctos y en un tiempo no mayor a 300 ms
