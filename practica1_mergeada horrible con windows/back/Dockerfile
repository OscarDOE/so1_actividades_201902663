FROM golang:alpine AS build
RUN  apk add --no-cache git
WORKDIR /app
COPY ["go.mod", "go.sum", "./"]
RUN go mod download
COPY . .
RUN go build -o app -v .
FROM alpine:latest
LABEL Name="backend" 
RUN apk update
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=build /app/app .
ENTRYPOINT [ "./app" ]
EXPOSE 5005 
# docker build --pull --rm -f "Dockerfile" -t 7771512001/back_so1_practica1_1s2023:latest "./"

# Correr mysql con el volumen
# docker run -d -p 3306:3306 --name mysql-db -e MYSQL_ROOT_PASSWORD=adminoscar --mount src=mysql-db-data,dst=/var/lib/mysql mysql
