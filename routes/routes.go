package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pos-backend/controllers/logincontroller"
	"strings"

	"github.com/gorilla/mux"
)

func Routes() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", AllUsers_GET).Methods("GET")
	router.HandleFunc("/login", UserLogin_POST).Methods("POST")

	// := os.Getenv("ASPNETCORE_PORT")
	//http.ListenAndServe(":"+port, router)
	http.ListenAndServe(":8059", router)
}

func AllUsers_GET(w http.ResponseWriter, r *http.Request) {
	jsonBytes, err := json.Marshal(logincontroller.ReadAllUser())
	if err != nil {
		fmt.Printf(err.Error())
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func UserLogin_POST(w http.ResponseWriter, r *http.Request) {
	var url = r.URL.Query()
	filters, present := url["username"]
	if !present || len(filters) == 0 {
		fmt.Println("filters not present")
	}

	filters1, present1 := url["password"]
	if !present1 || len(filters1) == 0 {
		fmt.Println("filters not present")
	}

	jsonBytes, err := json.Marshal(logincontroller.Login(strings.Join(filters, ","), strings.Join(filters1, ",")))
	if err != nil {
		fmt.Printf(err.Error())
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)

	//w.WriteHeader(200)
	//w.Write([]byte(strings.Join(filters, ",")))
	//w.Write([]byte(strings.Join(filters1, ",")))
}
