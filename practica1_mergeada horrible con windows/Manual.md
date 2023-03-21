# Manual Tecnico 

| Nombre | Carnet |
| ------------- | ------------- |
| Oscar Daniel Oliva España          | 201902663 |

## Componentes

> Golang

Se utilizaron las siguientes importaciones para la ejecucion del servidor backend en go:

![Captura](img/Imports%20go.png)

- Se utilizaron las importaciones de sql, para poder conectarse a la BD de MySQL.
- CORS, para que las peticiones http no se bloqueen
- MUX, para poder realizar los endpoints de las peticiones recibidas http
- Encoding json y csv, para poder escribir y escribir respuestase y un archivo csv respectivamente

Siempre se deben de realizar los comandos `go get [comando]` para cada uno de los imports externos como los de `github`

Otro aspecto a tomar en cuenta, es que la aplicacion de go, no se ejecuto, para el desarrollo, en la carpeta del propio repositorio, sino que se hizo dentro de la carpeta `src` de la propia ruta principal del lenguaje `go`.
> React

El servicio de react se utilizo para hacer el frontend de la aplicacion, practicamente no se hizo nada especifico, es una aplicacion de react normal con peticiones GET y POST.

> MySQL

Se utilizaron las bases de MySQL ejecutandose con el contenedor.
Se creo el volumen para el archivo de la persistencia de los datos de MySQL, por si se llegase a apagar el contenedor, y que al volver a subirlo la base siga existiendo.

Luego se corrio el contenedor de MySQL con el comando 
```docker run -d -p 3306:3306 --name mysql-db -e MYSQL_ROOT_PASSWORD=[Contraseña] --mount src=mysql-db-data,dst=/var/lib/mysql mysql```

para que este se pueda utilizar localmente, luego se utilizo las especificaciones del docker-compose.yml que esta en el repositorio, en la parte del servicio "db" con los volumenes, y que cree la base de datos en el mismo lugar en donde se corre el comando `docker-compose up -d`.


> Script

El script es un archivo `.sh`, el cual se utilizo para sacar los logs de las operaciones de la calculadora, de igual manera, se subio a un repositorio de docker, con el cual se ejecuto por medio de los volumenes, mandando a llamar al archivo `reportes.csv`, el cual se crea dentro de la carpeta `/home/oscar/Desktop` por parte del backend, para que el script lo pueda correr sin la necesidad de buscar una carpeta dentro de la ruta base del contenedor.


> Docker

Se utilizaron Dockerfiles para el backend con mysql y uno para el frontend con docker-compose para poder subir los servicios a dockerhub

Para correr el programa en contenedores solo se necesita realizar el comando `docker-compose up -d` para poder ejecutar los contenedores ya subidos.

Los volumenes ayudaron mucho para esta aplicacion, ya que se necesita de la base de datos, y que esta persista aun despues de bajar los contenedores con el comando `docker-compose down`.
De igual manera, con los volumenes podemos capturar el archivoe `reportes.csv` el cual se podra leer para hacer obtenere los logs necesarios de la calculadora.


Se utilizaron Dokcerfile para todos los servicios, menos para la base de datos, que esa se mando a llamar al contenedor oficial de `mysql`.

Se crearon 2 archivos de `docker-compose.yml`, 1 para el frontend, backend y base de datos; y otro para el script de bash.
