package main

// import (
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// type ApiServer struct {
// 	listenAddr string
// 	store      Store
// }

// func NewApiServer(listenAddr string, store Store) *ApiServer {
// 	return &ApiServer{listenAddr: listenAddr, store: store}
// }

// func (s *ApiServer) Serve() {
// 	router := mux.NewRouter()
// 	subRouter := router.PathPrefix("/api/v1").Subrouter()

// 	// Registering our services
// 	tasksService := NewTasksService(s.store)
// 	tasksService.RegisterRoutes(router)

// 	log.Println("Starting the API server at", s.listenAddr)

// 	log.Fatal(http.ListenAndServe(s.listenAddr, subRouter))
// }
