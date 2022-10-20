package rwfile

import(
	"fmt"
	"io/ioutil"
	"net/http"
	"k2io.com/keshav/model"
	"encoding/json"
	
)

func Rfile(w http.ResponseWriter, r *http.Request) {
    
    var response model.Response
    
    f := r.FormValue("location")
    

	content, err := ioutil.ReadFile(f)
    
    if err != nil {
        fmt.Println(err)
        response.Status = 404
    response.Message = err.Error()    
    }else{
    	fmt.Print(string(content))

	
    response.Status = 200
    response.Message = "Success"
    response.Output= string(content)
    }

   
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    json.NewEncoder(w).Encode(response)
}

////write in a file

func Wfile(w http.ResponseWriter, r *http.Request) {
    
    var response model.Response
    
    f := r.FormValue("location")
    mydata  :=r.FormValue("mydata")
    
    

    
//   var filemode os.FileMode = ()

	err := ioutil.WriteFile(f, []byte(mydata), 0777)
    
    if err != nil {
        fmt.Println(err)
        response.Status = 404
    response.Message = err.Error()    
    }else{
    	
    response.Status = 200
    response.Message = "Success"
    
    }

   
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    json.NewEncoder(w).Encode(response)
}