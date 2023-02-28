# Actividad 3 - Sistemas Operativos 1

## Enunciado

Actualmente existe un bug en la container image de frontend del ejemplo visto en clase. Lo que sucede es que al momento de cargar una URL especifica se obtiene un error 404. La actividad consiste en entender y solucionar el problema. Subir la solución al folder indicado y crear un readme.md file explicando el problema y cual es su solución.

Link del repositorio: `susguzman/so1_containers (github.com)`

> Nota: No debemos de moficar la app (frontend) para solucionar este problema.


Para poder arreglar el error, se tuvo que crear un nginx.conf, el cual es muy importante porque permite a los desarrolladores personalizar la configuración de Nginx para adaptarla a las necesidades de su aplicación web específica. Con él, se puede optimizar el rendimiento, mejorar la seguridad, y manejar solicitudes de los clientes de una manera más eficiente.
En este caso se configuraron las locaciones del index.html y ls maneras de llegar a ella, y si ocurriese algun error, indicarle que debe demostrar, se indico el puerto y el nombre del servidor.


Lo que se escribió en el archivo nginx.conf, el cual se creo dentro de la carpeta frontend,fue:


 server {

    listen       80;
    server_name  localhost;

    location / {
        root   /usr/share/nginx/html;
        index  index.html index.htm;
        try_files $uri /index.html;                 
    }

    error_page   500 502 503 504  /50x.html;
    location = /50x.html {
        root   /usr/share/nginx/html;
    }
}

De igual manera, se modificó el nginx.Dockerfile, agregandole luego de la sección "Fase # 2", los comandos:

`RUN rm /etc/nginx/conf.d/default.conf` como segunda fila

`COPY nginx.conf /etc/nginx/conf.d/` como tercera fila

`EXPOSE 80` como quinta fila

Solamente con eso bastaría para arreglar el bug, entre de las tantas maneras que existen.