package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Nykseli/solita-backend/db"
)

// StartServer starts the api server and starts to listen to requests
func StartServer() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/name/total", nameTotal).Methods("GET")
	// by argument can be "name" or "amount"
	router.HandleFunc("/name/sort/{by}", sortedNames).Methods("GET")
	router.HandleFunc("/name/amount/{name}", nameAmount).Methods("GET")
	http.ListenAndServe(":8888", router)
}

func notImplemented(w http.ResponseWriter, r *http.Request) {
	panic("Endpoint not implmeneted")
}

func nameTotal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	total := db.GetTotalNameAmount()
	jsonStr := fmt.Sprint("{\"total\": ", total, "}")
	w.Write([]byte(jsonStr))
}

func nameAmount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)
	name := params["name"]
	amount := db.GetNameAmount(name)
	jsonStr := fmt.Sprint("{\"amount\": ", amount, "}")
	w.Write([]byte(jsonStr))
}

func sortedNames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var nameData db.NameData

	params := mux.Vars(r)
	by := params["by"]
	switch by {
	case "name":
		nameData = db.GetAllSortedName()
	case "amount":
		nameData = db.GetAllSortedAmount()
		// TODO: invalid argument handler
	}

	json.NewEncoder(w).Encode(nameData)
}
