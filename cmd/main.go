package main

//importing all necessary package //
import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
	handlers "github.com/lm-goutam/testing/pkg/handlers"
	"github.com/rs/cors"
)

// main method
func main() {

	router := mux.NewRouter()
	router.HandleFunc("/org", handlers.GetAllOrg).Methods("GET")
	router.HandleFunc("/status", handlers.GetAllStat).Methods("GET")
	router.HandleFunc("/cms", handlers.GetAllCms).Methods("GET")
	router.HandleFunc("/app", handlers.GetAllApp).Methods("GET")
	router.HandleFunc("/intgs", handlers.GetAllIntgs).Methods("GET")
	router.HandleFunc("/intgs", handlers.PostAllIntgs).Methods("POST")
	router.HandleFunc("/intgs/{i_id}", handlers.DeleteIntgsById).Methods("DELETE")
	router.HandleFunc("/intgs/{i_id}", handlers.UpdateIntgsById).Methods("PUT")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	fmt.Println("Server running on port : 8010")
	http.ListenAndServe(":8010", handler)
}
