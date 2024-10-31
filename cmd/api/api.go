package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yann-fk-21/todo-platform/service/product"
	"github.com/yann-fk-21/todo-platform/service/user"
)

type ApiServer struct {
	Address string
	db      *sql.DB
}


func NewApiServer(address string, db *sql.DB) *ApiServer {
	return &ApiServer{
		Address: address,
		db: db,
	}
}


func (s *ApiServer) Run() error {
     router := mux.NewRouter()
	 subrouter := router.PathPrefix("/api/v1").Subrouter()
     
	 userStore := user.NewStore(s.db)
	 userService := user.NewHandler(userStore)
	 userService.RegisterRoutes(subrouter) 

	 productStore := product.NewStore(s.db)
	 productHandler := product.NewHandler(productStore)
	 productHandler.RegisterRoutes(subrouter)


	log.Println("Server listen on port 8080")
	return http.ListenAndServe(s.Address, router)
}