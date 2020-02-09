package endpoints

import (
	"github.com/rahul-golang/mongocrud/pkg/handler"

	"github.com/gorilla/mux"
)

//NewUserRoute All Application Routes Are defiend Here
func NewUserRoute(router *mux.Router, handler *handler.UserHandlersImpl) {
	router.HandleFunc("/v1/users", handler.CreateUser).Methods("POST")
	router.HandleFunc("/v1/users", handler.GetUsers).Methods("GET")
	router.HandleFunc("/v1/users/{id}", handler.GetUserByID).Methods("GET")
	router.HandleFunc("/v1/users/{id}", handler.UpdateUser).Methods("PUT")
	router.HandleFunc("/v1/users/{id}", handler.DeleteUser).Methods("Delete")

}
