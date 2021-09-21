package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"restgo/models"

	"database/sql"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql" // this is mysql driver import
)

// Database will establish database connection
func Database() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "ashish"
	dbName := "hellosql"
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		log.Println("database connectin error")
		return
	}
	return
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "customer to be created")

	db, err := Database()
	if err != nil {
		log.Println("db error ", err)
		return
	}
	defer db.Close()

	var customer models.Customer
	decodeErr := json.NewDecoder(r.Body).Decode(&customer)

	if decodeErr != nil {
		fmt.Print("JSON Decode problem create")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	statement, err := db.Prepare("INSERT INTO students (`fname`, `lname`) values (?, ?)")
	if err != nil {
		log.Println("prepare error ", err)
		return
	}

	insertresult, err := statement.Exec(customer.FName, customer.LName)
	if err != nil {
		log.Println("insert error ", err)
		return
	}

	insertid, err := insertresult.LastInsertId()
	log.Println("this is insert result", insertid, err)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(fmt.Sprint(insertid))
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "customer to be get")

	db, err := Database()
	if err != nil {
		return
	}
	defer db.Close()

	resultSet, err := db.Query("SELECT * FROM students")
	if err != nil {
		return
	}
	var customerList []models.Customer

	for resultSet.Next() {
		singleCustomer := models.Customer{}

		err = resultSet.Scan(&singleCustomer.ID, &singleCustomer.FName, &singleCustomer.LName)
		if err != nil {
			log.Println("error occurred while getting customers", err)
			return
		}
		customerList = append(customerList, singleCustomer)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customerList)
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "customer to be deleted ")

	db, err := Database()
	if err != nil {
		log.Println("db error ", err)
		return
	}
	defer db.Close()

	params := mux.Vars(r)

	id := params["customerNo"]
	delStatement, err := db.Prepare("DELETE FROM students WHERE id=?")
	if err != nil {
		log.Println("error occurred while deleting user", err)
		return
	}
	deleteresult, err := delStatement.Exec(id)
	if err != nil {
		log.Println("error occurred while deleting user", err)
		return
	}
	num, err := deleteresult.RowsAffected()
	log.Println("error ", err)
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode("customer deleted " + fmt.Sprint(num))
}

func welcome(w http.ResponseWriter, r *http.Request) {

	str := fmt.Sprintf("welcome to CRUD sample. " +
		"Create Read Update and Delete customer")

	json.NewEncoder(w).Encode(str)
}

// MyController ...
func MyController() {
	log.Println("hello world")

	// ip, err := utils.GetLocalIP()
	// if err != nil {
	// 	log.Println("ip find error ", err)
	// 	return
	// }
	r := mux.NewRouter()

	r.HandleFunc("/", welcome).Methods("GET", "PUT", "POST", "DELETE")
	r.HandleFunc("/api/customers", createCustomer).Methods("POST")
	r.HandleFunc("/api/customers", getCustomers).Methods("GET")
	r.HandleFunc("/api/customer/{customerNo}", deleteCustomer).Methods("DELETE")

	log.Println("application running on port ", ":1234")
	http.ListenAndServe(":1234", r)
}
