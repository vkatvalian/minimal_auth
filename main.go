package main

import (
    "log"
    "context"
    "net/http"
    "html/template"

    "github.com/vkatvalian/auth/helpers"
    "github.com/vkatvalian/auth/database"
)

type Handlers struct {
    helper helpers.Helper
}

func (h *Handlers) handler(w http.ResponseWriter, req *http.Request) {
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

    // insert into db
    h.helper.Insert(req.Context(), username, email, password)
    
    // check if exists on db
}

func main(){
    ctx := context.Background()
    db := database.Connection(ctx)
    db.CreateUsersTable(ctx)
    h := helpers.Helper{db}
    handlers := &Handlers{h}

    http.HandleFunc("/signup", handlers.handler)
    log.Fatal(http.ListenAndServe(":8080", nil))

}
