package main

import (
    "log"
    "context"
    "net/http"
    "html/template"
    "golang.org/x/crypto/bcrypt"

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
    password := []byte(req.FormValue("password"))

    hashed_password, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
    if err != nil {
        panic(err)
    }

    err = bcrypt.CompareHashAndPassword(hashed_password, password)
    if err != nil {
        log.Fatal(err)
    }

    h.helper.Insert(req.Context(), username, email, string(hashed_password))
    
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
