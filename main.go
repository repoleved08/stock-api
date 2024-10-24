package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/repoleved08/stock-api/router"
)

func main() {
	r := router.Router()

	fmt.Println("Server starting in port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
