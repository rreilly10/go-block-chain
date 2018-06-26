package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kr/pretty"
	"github.com/satori/go.uuid"
)

// API Context
type API struct {
	bc *blockchain
}

// Mine endpoint
func (api *API) Mine(w http.ResponseWriter, r *http.Request) {
	fmt.Println(api.bc)
	json.NewEncoder(w).Encode(api.bc)
}

// NewTransaction endpoint
func (api *API) NewTransaction(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello World")
}

// Chain endpoint
func (api *API) Chain(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello World")
}

func main() {

	bc := blockchain{
		chain:               []block{},
		currentTransactions: []transaction{},
	}

	bc.newBlock(1, "100") // 0 block
	mine(&bc)
	bc.newTransaction("0", "new", 1)
	mine(&bc)
	pretty.Print(bc)

	api := API{
		bc: &bc,
	}

	// or error handling
	u2, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return
	}

	fmt.Printf("UUIDv4: %s\n", u2)
	router := mux.NewRouter()
	router.HandleFunc("/mine", api.Mine).Methods("GET")
	router.HandleFunc("/transactions/new", api.NewTransaction).Methods("GET")
	// router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	// router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))

}
