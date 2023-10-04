package route

import (
	"log"
	"net/http"
	"pandemix/handler"

	"github.com/gorilla/mux"
)

func Run() {
	port := "3000"
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/covid", handler.GetCovidData).Methods("GET")
	router.HandleFunc("/api/patient", handler.CreatePatient).Methods("POST")
	router.HandleFunc("/api/patient", handler.GetPatient).Methods("GET")
	router.HandleFunc("/api/patient/{id}", handler.GetPatientByID).Methods("GET")
	router.HandleFunc("/api/patient/{id}", handler.UpdatePatient).Methods("PUT")
	router.HandleFunc("/api/patient/{id}", handler.DeletePatient).Methods("DELETE")

	server := http.Server{
		Addr:    "localhost:" + port,
		Handler: router,
	}

	log.Printf("Server is starting on port %s\n", port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Server couldn't start: %v", err)
	}
}
