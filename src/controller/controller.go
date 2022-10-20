package controller
 
import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
 
//    "../model"
    "k2io.com/keshav/model"
 "k2io.com/keshav/config"
 
//    "../config"
)
 
// AllEmployee = Select Employee API
func AllEmployee(w http.ResponseWriter, r *http.Request) {
    var employee model.Employee
    var response model.Response
    var arrEmployee []model.Employee
 
    db := config.Connect()
    defer db.Close()
 
    rows, err := db.Query("SELECT * FROM employee")
 
    if err != nil {
        log.Print(err)
        response.Status = 404
    response.Message = err.Error()
    }else{
    	for rows.Next() {
        err = rows.Scan(&employee.Name, &employee.Age, &employee.Salary, &employee.Location)
        if err != nil {
            log.Fatal(err.Error())
        } else {
            arrEmployee = append(arrEmployee, employee)
        }
    }
 
    response.Status = 200
    response.Message = "Success"
    response.Data = arrEmployee
    }
 
    
 
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    json.NewEncoder(w).Encode(response)
}
 
 
 //insert employee api
 
 
// func InsertEmployee(w http.ResponseWriter, r *http.Request) {
//    var response model.Response
// 
//    db := config.Connect()
//    defer db.Close()
// 
//    err := r.ParseMultipartForm(4096)
//    if err != nil {
//        panic(err)
//    }
//    name := r.FormValue("name")
//    age := r.FormValue("age")
//    salary := r.FormValue("salary")
//    location := r.FormValue("lacation")
// 
//    _, err = db.Exec("INSERT INTO employee(name, age, salary, location) VALUES(?, ?, ?, ?)", name, age, salary, location)
// 
//    if err != nil {
//        log.Print(err)
//        return
//    }
//    response.Status = 200
//    response.Message = "Insert data successfully"
//    fmt.Print("Insert data to database")
// 
//    w.Header().Set("Content-Type", "application/json")
//    w.Header().Set("Access-Control-Allow-Origin", "*")
//    json.NewEncoder(w).Encode(response)
//}
 
 /////////////////
 ///// insert employee api 
 func InsertEmployee(w http.ResponseWriter, r *http.Request) {
    db := config.Connect()
    var response model.Response
    if r.Method == "POST" {
        name := r.FormValue("name")
	    age := r.FormValue("age")
	    salary := r.FormValue("salary")
	    location := r.FormValue("location")
	    
        insForm, err := db.Prepare("INSERT INTO Employee(name, age, salary, location) VALUES(?, ?, ?, ?)")
        if err != nil {
            panic(err.Error())
        }
        insForm.Exec(name, age, salary, location)

    }
    defer db.Close()
    response.Status = 200
    response.Message = "Insert data successfully"
//    response.Data="Inserted"
    fmt.Print("Insert data to database")
 
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
//    w.Write([]byte(`{"message":"hi this is me"}`))
    json.NewEncoder(w).Encode(response)
    http.Redirect(w, r, "/", 301)
}
 
 
 
 
 //Update employee api
 
 func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
    db := config.Connect()
     var response model.Response
    if r.Method == "PUT" {
        name := r.FormValue("name")
	    age := r.FormValue("age")
	    salary := r.FormValue("salary")
	    location := r.FormValue("location")
	    
        insForm, err := db.Prepare("UPDATE Employee SET age=?, salary=?, location=? WHERE name=?")
        if err != nil {
            panic(err.Error())
            
        }
        insForm.Exec(age, salary, location, name)

    }
    defer db.Close()
    
    response.Status = 200
    response.Message = "Data Upadte successfully"
//    response.Data="Inserted"
    fmt.Print("update data in database")
 
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
//    w.Write([]byte(`{"message":"hi this is me"}`))
    json.NewEncoder(w).Encode(response)
    http.Redirect(w, r, "/", 301)
    
}


//delete employee

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
    db := config.Connect()
    var response model.Response
//    emp := r.URL.Query().Get("id")
    delForm, err := db.Prepare("DELETE FROM Employee WHERE name=?")
    if err != nil {
        panic(err.Error())
    }
    emp_name := r.FormValue("name")
    delForm.Exec(emp_name)
    log.Println("DELETE")
    defer db.Close()
    response.Status = 200
    response.Message = "Employee Deleted successfully"
//    response.Data="Inserted"
    fmt.Print("delete employee in database")
 
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
//    w.Write([]byte(`{"message":"hi this is me"}`))
    json.NewEncoder(w).Encode(response)
    http.Redirect(w, r, "/", 301)
    
}

