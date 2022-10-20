
package command

import (
//	"fmt"
"net/http"
	"log"
	"os/exec"
	"strings"
	"k2io.com/keshav/model"
	"encoding/json"
)



func ExecuteCommand(w http.ResponseWriter, r *http.Request) {
    
    var response model.Response
    
    s := r.FormValue("command")
    

	args := strings.Split(s, " ")
	
	cmd := exec.Command(args[0], args[1:]...)
	b, err := cmd.CombinedOutput()
	
	if err != nil{
		log.Print(err)
	}
	
    response.Status = 200
    response.Message = "Success"
    response.Output= string(b)
    if(response.Output==""){
    	response.Status = 404
    response.Message = "Command not found"
    }
 
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    json.NewEncoder(w).Encode(response)
}
 