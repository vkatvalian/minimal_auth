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

func (h *Handlers) signup(w http.ResponseWriter, req *http.Request) {
    temp, err := template.ParseFiles("tmpl/signup.tmpl")
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

    // check if exists on db
    if username != "" && email != "" && password != "" {
        hashed_password, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
        if err != nil {
            panic(err)
        }

        err = bcrypt.CompareHashAndPassword(hashed_password, []byte(password))
        if err != nil {
            log.Fatal(err)
        }

        h.helper.Insert(req.Context(), username, email, string(hashed_password))
    } else {
        log.Println("all fields are required")
    }
}

func (h *Handlers) signin(w http.ResponseWriter, req *http.Request) {
    temp, err := template.ParseFiles("tmpl/signin.tmpl")
    if err != nil {
        log.Fatal(err)
    }

    err = temp.Execute(w, nil)
    if err != nil{
        log.Fatal(err)
    }
    username_form := req.FormValue("username")
    password_form := req.FormValue("password")

    if username_form != "" && password_form != "" {
        username, email, password, _ := h.helper.Fetch(req.Context(), username_form)
    
        ok := bcrypt.CompareHashAndPassword([]byte(password), []byte(password_form))
        if ok != nil {
            log.Println(ok)
        }
	log.Println(username)
	log.Println(email)
    } else {
        log.Println("all fields are required")
    }
}

func main(){
    ctx := context.Background()
    db := database.Connection(ctx)
    db.CreateUsersTable(ctx)
    h := helpers.Helper{db}
    handlers := &Handlers{h}

    http.HandleFunc("/signup", handlers.signup)
    http.HandleFunc("/signin", handlers.signin)
    log.Fatal(http.ListenAndServe(":8080", nil))

}
