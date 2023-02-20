package main

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type History struct {
	Id        int     `json:"id"`
	Num1      float64 `json:"num1"`
	Operation string  `json:"operation"`
	Num2      float64 `json:"num2"`
	Resultado float64 `json:"resultado"`
	Fecha     string  `json:"fecha"`
}

var g bool = false

func GetHistoryEndpoint(w http.ResponseWriter, req *http.Request) {

	db, err := sql.Open("mysql", "root:adminoscar@tcp(localhost:3306)/pruebascalcu")
	if err != nil {
		fmt.Println("Error validatinfg sql.Open arguments")
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("error verifying connection with db.Ping ")
		panic(err.Error())
	}

	lll := "history"
	datos, err := db.Query("select * from " + lll)
	if err != nil {
		panic(err.Error())
	}

	defer datos.Close()

	var id int
	var num1 float64
	var operation string
	var num2 float64
	var resultado float64
	var fecha string

	BDCopy := []History{}

	for datos.Next() {
		err := datos.Scan(&id, &num1, &operation, &num2, &resultado, &fecha)
		if err != nil {
			panic(err.Error())
		}
		//		fmt.Println("ID: %d,  NUM 1: %f,   Operacion: %s, Num 2: %f,   Resultado: %f,  Fecha: %s", id, num1, operation, num2, resultado, fecha)

		fmt.Println("ID: ", id, "NUM 1: ", num1, "Operation: ", operation, " NUM2:", num2, " Resultado: ", resultado, "Fecha: ", fecha)
		BDCopy = append(BDCopy, History{id, num1, operation, num2, resultado, fecha})
	}
	fmt.Println(resultado + 5.5)
	fmt.Println("SUCCESFULL Connection to Database!")
	//Devolver los datos en formato JSON
	w.Header().Set("Content-Tyepe", "application/json")
	err = json.NewEncoder(w).Encode(BDCopy)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Printf("Error al enviar Header")
		return
	}

}

func CreateHistoryEndpoint(w http.ResponseWriter, req *http.Request) {
	var history History
	_ = json.NewDecoder(req.Body).Decode(&history)
	//json.NewEncoder(w).Encode(history)
	fmt.Println("HISTORY", history)
	//json.NewEncoder(w).Encode(history)

	db, err := sql.Open("mysql", "root:adminoscar@tcp(localhost:3306)/pruebascalcu")
	if err != nil {
		fmt.Println("Error validatinfg sql.Open arguments")
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("error verifying connection with db.Ping ")
		panic(err.Error())
	}

	fmt.Println("111111111111", history.Num1, "2222222222", history.Num2)

	fmt.Println("insert into history (num1,operation,num2,resultado,fecha) values(", history.Num1, ",\"", history.Operation, "\",", history.Num2, ",", history.Resultado, ",now());")
	fmt.Println("insert into history (num1,operation,num2,resultado,fecha) values( 9 ,\" / \", 3 , 3 ,now());")

	stmt, err := db.Prepare("INSERT INTO history(num1,operation,num2,resultado,fecha) values(?,?,?,?,now())")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var respuesta float64
	eserror := false

	if history.Operation == "+" {
		respuesta = history.Num1 + history.Num2
	} else if history.Operation == "-" {
		respuesta = history.Num1 - history.Num2
	} else if history.Operation == "*" {
		respuesta = history.Num1 * history.Num2
	} else if history.Operation == "/" {
		if history.Num2 == 0 {
			eserror = true
		} else {
			respuesta = history.Num1 / history.Num2
		}
	} else {
		eserror = true
	}

	if g == false {
		g = !g
	} else if eserror {
		valores := History{Id: history.Id, Num1: history.Num1, Operation: history.Operation, Num2: history.Num2, Resultado: respuesta, Fecha: history.Fecha}
		fmt.Println("valores", valores)
		_, err = stmt.Exec(valores.Num1, valores.Operation, valores.Num2, -123457000.00)

		datos, err := db.Query("select * from history")
		if err != nil {
			panic(err.Error())
		}
		defer datos.Close()
		newhistory := History{
			Id:        history.Id,
			Num1:      history.Num1,
			Operation: history.Operation,
			Num2:      history.Num2,
			Resultado: respuesta,
			Fecha:     history.Fecha,
		}
		fmt.Println("ULTIMOE ", newhistory)
		json.NewEncoder(w).Encode(newhistory)
	} else {
		valores := History{Id: history.Id, Num1: history.Num1, Operation: history.Operation, Num2: history.Num2, Resultado: respuesta, Fecha: history.Fecha}
		fmt.Println("valores", valores)
		_, err = stmt.Exec(valores.Num1, valores.Operation, valores.Num2, valores.Resultado)

		datos, err := db.Query("select * from history")
		if err != nil {
			panic(err.Error())
		}

		defer datos.Close()
		fmt.Println("DESPUES DEL INSERT")
		var id int
		var num1 float64
		var operation string
		var num2 float64
		var resultado float64
		var fecha string

		BDCopy := []History{}

		for datos.Next() {
			err := datos.Scan(&id, &num1, &operation, &num2, &resultado, &fecha)
			if err != nil {
				panic(err.Error())
			}
			fmt.Println("ID: ", id, "NUM 1: ", num1, "Operation: ", operation, " NUM2:", num2, " Resultado: ", resultado, "Fecha: ", fecha)
			BDCopy = append(BDCopy, History{id, num1, operation, num2, resultado, fecha})
		}
		fmt.Println(resultado + 5.5)
		fmt.Println("SUCCESFULL Connection to Database!")
		//Devolver los datos en formato JSON
		w.Header().Set("Content-Tyepe", "application/json")
		err = json.NewEncoder(w).Encode(BDCopy)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Printf("Error al enviar Header")
			return
		}
		newhistory := History{
			Id:        history.Id,
			Num1:      history.Num1,
			Operation: history.Operation,
			Num2:      history.Num2,
			Resultado: respuesta,
			Fecha:     history.Fecha,
		}
		fmt.Println("ULTIMOE ", newhistory)
		json.NewEncoder(w).Encode(history)
	}

}

func ExportCsv() error {

	db, err := sql.Open("mysql", "root:adminoscar@tcp(localhost:3306)/pruebascalcu")
	if err != nil {
		fmt.Println("Error validatinfg sql.Open arguments")
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("error verifying connection with db.Ping ")
		panic(err.Error())
	}

	// Ejecutar una consulta
	rows, err := db.Query("SELECT num1,operation,num2,resultado,fecha FROM history")
	if err != nil {
		return err
	}
	defer rows.Close()

	// Abrir un archivo para escribir el resultado
	file, err := os.Create("history.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	// Escribir los resultados en el archivo CSV
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Escribir el encabezado
	encabezado := []string{"num1", "operation", "num2", "resultado", "fecha"}
	if err := writer.Write(encabezado); err != nil {
		return err
	}

	// Procesar los resultados
	for rows.Next() {
		var num1 float64
		var num2 float64
		var operation string
		var resultado float64
		var fecha string
		if err := rows.Scan(&num1, &operation, &num2, &resultado, &fecha); err != nil {
			return err
		}
		history := []string{fmt.Sprintf("%.2f", num1), operation, fmt.Sprintf("%.2f", num2), fmt.Sprintf("%.2f", resultado), fecha}
		if err := writer.Write(history); err != nil {
			return err
		}
	}
	if err := rows.Err(); err != nil {
		return err
	}
	return nil
}

func GetCSV(w http.ResponseWriter, req *http.Request) {
	err := ExportCsv()
	if err != nil {
		fmt.Println("CUALQUIERA")
		fmt.Fprintf(w, err.Error())
	}
	fmt.Println("SSSSS")

}

func main() {

	router := mux.NewRouter()

	//endpoints

	router.HandleFunc("/history", GetHistoryEndpoint).Methods("GET")
	router.HandleFunc("/history", CreateHistoryEndpoint).Methods("POST")
	router.HandleFunc("/csv", GetCSV).Methods("GET")

	c := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":5005", c))

}

//insert, err := db.Query("insert into history (num1,operation,num2,resultado,fecha) values( 9 ,\" / \", 3 , 3 ,now());")

//insert, err := db.Query("insert into history (num1,operation,num2,resultado,fecha) values( ", history.Num1, ",'", history.Operation, "',", history.Num2, ",", history.Resultado, ",now());")
/*query := "insert into `history` (`num1`,`operation`,`num2`,`resultado`) values (?,?,?,?,now())"
insertResult, err := db.ExecContext(context.Background(), query, history.Num1, history.Operation, history.Num2, history.Resultado)
if err != nil {
	//log.Fatalf("NO JALO")
	panic(err.Error())
}

idd, err := insertResult.LastInsertId()
if err != nil {
	log.Fatalf("impossible to retrieve las inserted id: %s", err)
}
log.Printf("inserted id: %d", idd)*/
