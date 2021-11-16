package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

    db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
    // "postgres://ptntmqtillzjye:4c40f520658ff4c0ef5e2884c2c5443ea516a1dd508f88c4887bb810ec076b4f@ec2-3-92-119-83.compute-1.amazonaws.com:5432/d7mct601f3o8be"
    if err != nil {
        log.Fatalf("Error opening database: %q", err)
    }
    
    rows, err := db.Query("SELECT * FROM test")
    if err != nil {
        fmt.Println(err)
    }

    type EMPLOYEE struct {
    ID     string
    NUMBER string
}

    var es []EMPLOYEE
    for rows.Next() {
        var e EMPLOYEE
        rows.Scan(&e.ID, &e.NUMBER)
        es = append(es, e)
    }
    fmt.Printf("%v/n", es)



    engine:= gin.Default()
    engine.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "hello world",
        })
    })
    engine.Run(":" + os.Getenv("PORT"))
}