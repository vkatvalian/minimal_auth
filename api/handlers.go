package api

import (
    "log"
    "net/http"
    "html/template"
    "golang.org/x/crypto/bcrypt"
    "github.com/vkatvalian/auth/database"
)

type API struct {
    DB *database.Repository
}

func NewAPI(db *database.Repository) *API {
   return &API {
        DB: db,
   } 
}

func (api *API) LoginHandler(w http.ResponseWriter, req *http.Request) {
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
        username, email, password, err := api.DB.FetchUsers(req.Context(), username_form)
	if err != nil {}

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

func (api *API) SignupHandler(w http.ResponseWriter, req *http.Request) {
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

    if username != "" && email != "" && password != "" {
        hashed_password, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
        if err != nil {
            panic(err)
        }

        err = bcrypt.CompareHashAndPassword(hashed_password, []byte(password))
        if err != nil {
            log.Fatal(err)
        }

	err = api.DB.InsertUsers(req.Context(), username, email, string(hashed_password))
	log.Println(err)
    } else {
        log.Println("all fields are required")
    }
}
