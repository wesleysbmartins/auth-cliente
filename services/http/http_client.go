package http

import (
	routes "auth-users-aws/domain/route"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func HttpClient() {
	router := mux.NewRouter()
	routes.Routes(router)
	http.Handle("/", router)

	port := os.Getenv("PORT")

	fmt.Println(fmt.Sprintf("Server listenning on Port %s", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
