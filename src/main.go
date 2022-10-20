
package main

import (
	 _ 
	"fmt"
	"log"
	"net/http"
    "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

    "k2io.com/keshav/controller"
	"k2io.com/keshav/config"
	"k2io.com/keshav/command"
	"k2io.com/keshav/rwfile"

 
)

func main() {
	
	
	
	
	
	//using gorilla/mux
	go func(){
		
		router := mux.NewRouter()
	router.HandleFunc("/getEmployee", controller.AllEmployee).Methods("GET")
	router.HandleFunc("/insertEmployee", controller.InsertEmployee).Methods("POST")
	router.HandleFunc("/updateEmployee", controller.UpdateEmployee).Methods("PUT")
	router.HandleFunc("/deleteEmployee", controller.DeleteEmployee).Methods("DELETE")
	router.HandleFunc("/executeCommand", command.ExecuteCommand).Methods("GET")
	router.HandleFunc("/readFile", rwfile.Rfile).Methods("GET")
	router.HandleFunc("/writeFile", rwfile.Wfile).Methods("POST")
//	http.Handle("/", router)
	http.Handle("/", router)
	fmt.Println("Connected to port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
	}()
	
	
	
	////task 6 using gin
	r := gin.Default()
	r.GET("/getemps", func(c *gin.Context) {
	emps := config.GetEmps()

	if emps == nil || len(emps) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, emps)
	}
})
	r.Run(":8090")
	
	
	
}



//func Function() {
//	
//}
//
//func function() {
//	
//}
//
//go mod init
//go get 
//go mod tidy
//go run
//go list 

//export GOROOT ="/usr/local/go"
//export GOPATH="/Users/keshav/Desktop/GoProject/src"
//
//
//export PATH="/Users/keshav/Desktop/GoProject/bin"