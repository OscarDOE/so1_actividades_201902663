package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type datos struct {
	lnum   float32
	rnum   float32
	result float32
}

func main() {
	db, err := sql.Open("mysql", "root:adminoscar@tcp(3306:3306)/pruebasoscar")
	if err != nil {
		fmt.Printf("DIO ERRPR ESTA MADRE")
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM history")
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var campo1 string
		var campo2 int
		err = rows.Scan(&id, &campo1, &campo2)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("id: %d, campo1: %s, campo2: %d\n", id, campo1, campo2)
	}
}

/*func main() {
bd, err := getDB()
if err != nil {
	log.Printf("Error with database" + err.Error())
	return
} else {
	err = bd.Ping()
	if err != nil {
		log.Printf("error making connection to DB. Please check credentials. THe error is: " + err.Error())
		return
	}
}

router := mux.NewRouter()
setupRoutesForVideoGames(router)

port := ":8000"

server := &http.Server{
	Handler: router,
	Addr:    port,

	WriteTimeout: 15 * time.Second,
	ReadTimeout:  15 * time.Second,
}
log.Printf("Server started at %s", port)
log.Fatal(server.ListenAndServe())
*/
/*

	db, err := sql.Open("mysql", "root:adminoscar@tcp(3306:3306)/pruebasoscar")
	if err != nil {
		fmt.Println("HUBO ERROR")
	}
	fmt.Println(db)
	defer db.Close()

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {})
	fmt.Println("Hello World")
	http.ListenAndServe("localhost:5000", nil)*/
//}
