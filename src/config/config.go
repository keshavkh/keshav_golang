package config
 
import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
	"fmt"
	"k2io.com/keshav/model"
    
)
 
func Connect() *sql.DB {
    dbDriver := "mysql"
    dbUser := "root"
    dbPass := "Keshav@112"
    dbName := "Employee"
 
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(host.docker.internal:3306)/"+dbName)
    if err != nil {
        panic(err.Error())
    }
    return db
}


func GetEmps() []model.Employee{

	var employee model.Employee
	var emps []model.Employee
    dbUser := "root"
    dbPass := "Keshav@112"
    dbName := "Employee"
	db, err := sql.Open("mysql", dbUser+":"+dbPass+"@tcp(host.docker.internal:3306)/"+dbName)

	// if there is an error opening the connection, handle it
	if err != nil {
		// simply print the error to the console
		fmt.Println("Err", err.Error())
		// returns nil on error
		return nil
	}

	defer db.Close()
	results, err := db.Query("SELECT * FROM employee")

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	
	for results.Next() {

		err = results.Scan(&employee.Name, &employee.Age, &employee.Salary, &employee.Location)
		if err != nil {
			panic(err.Error()) 
		}
       
		emps = append(emps, employee)
	}

	return emps

}