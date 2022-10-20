package model
 
type Employee struct {
    
    Name string `form:"name" json:"name"`
    Age int `form:"age" json:"age"`
    Salary int `form:"salary" json:"salary"`
    Location string `form:"location" json:"location"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Employee
	Output string `json:"output"`
}



//FROM golang:1.19-alpine3.16
//RUN mkdir /build
//WORKDIR /build
//RUN apk update && apk add --no-cache git
//
//RUN go get "github.com/gin-gonic/gin"
//RUN go get _ "github.com/go-sql-driver/mysql"
//RUN go get "github.com/gorilla/mux"
//	
//
//COPY go.mod go.sum ./
//RUN go mod download
//COPY . .
//
//RUN go build
//EXPOSE 8080
//CMD ["./main"]




//FROM alpine:3.16
//WORKDIR /app
//COPY --from=builder /app/main .