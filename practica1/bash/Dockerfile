# Define la imagen base que se usará
FROM alpine:latest

# Crea un directorio de trabajo para la aplicación
WORKDIR /app

# Copia el archivo .sh en la imagen
COPY script.sh /app/

# Establece permisos de ejecución en el archivo
RUN chmod +x script.sh

# Ejecuta el archivo .sh al iniciar el contenedor
CMD ["./script.sh"]