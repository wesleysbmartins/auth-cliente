package routes

import (
	controllers "auth-users-aws/domain/controller"

	"github.com/gorilla/mux"
)

func Routes(router *mux.Router) {
	router.HandleFunc("/", controllers.Home).Methods("GET")
	router.HandleFunc("/getClienteByCpf", controllers.GetClientByCpfHandler).Methods("POST")
	router.HandleFunc("/getAllClientesfromCache", controllers.GetAllClientesHandler).Methods("GET")
}
