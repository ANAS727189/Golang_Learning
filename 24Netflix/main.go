// main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/ANAS727189/Netflix-Api/controllers"
	"github.com/ANAS727189/Netflix-Api/routers"
)

func main() {
	r := routers.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4040", r))
	fmt.Println("Server is running on port 4040")
}
