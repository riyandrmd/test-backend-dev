package main

import (
	"fmt"
	"net/http"
	"test-backend-dev/connection"
	"test-backend-dev/routers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	db := connection.ConnectDatabase()

	eng := routers.Routers{
		R:  r,
		Db: db,
	}

	eng.NewRouters()
	fmt.Println("Runing in port 5000")
	fmt.Println(http.ListenAndServe(":5000", r))

}
