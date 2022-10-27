package main

import (
    "log"
    "net/http"
    "html/template"
)

type User struct{
    Username string
    Email    string
    Password string
}

func handler(w http.ResponseWriter, req *http.Request) {
    temp, err := template.ParseFiles("tmpl/singup.tmpl")
    if err != nil {
        log.Fatal(err)
    }

    err = temp.Execute(w, nil)
    if err != nil{
        log.Fatal(err)
    }
    username := req.FormValue("username")
    email := req.FormValue("email")
    password := req.FormValue("password")
    log.Println(username)
    log.Println(email)
    log.Println(password)

    // create db conenction
    // insert into db
    // check if exists on db
}

func main(){
    http.HandleFunc("/signup", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
