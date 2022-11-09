package main

import (
    "log"
    "context"
    "net/http"
    "github.com/vkatvalian/auth/api"
    "github.com/vkatvalian/auth/database"
)

func main(){
    ctx := context.Background()
    database := database.Connect(ctx)
    routes := api.NewAPI(database)
    http.HandleFunc("/signup", routes.SignupHandler)
    http.HandleFunc("/login", routes.LoginHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
