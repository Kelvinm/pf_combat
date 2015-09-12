package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

type Character struct {
	Name  string
	Class string
	Level int64
}

func SayHelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func ReturnCharacter(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("returned character"))
	c := Character{"Derpy", "Fighter", 5}
	msg, _ := json.Marshal(c)

	w.Write(msg)

}

func Greet(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	w.Write([]byte(fmt.Sprintf("Hello %s !", name)))
}

func ProcessPathVariables(w http.ResponseWriter, r *http.Request) {

	// break down the variables for easier assignment
	vars := mux.Vars(r)
	name := vars["name"]
	job := vars["job"]
	age := vars["age"]
	w.Write([]byte(fmt.Sprintf("Name is %s ", name)))
	w.Write([]byte(fmt.Sprintf("Job is %s ", job)))
	w.Write([]byte(fmt.Sprintf("Age is %s ", age)))
}

func main() {
	mx := mux.NewRouter()

	mx.HandleFunc("/ret_char", ReturnCharacter)
	mx.HandleFunc("/", SayHelloWorld)
	mx.HandleFunc("/{name}", Greet)
	loggedRouter := handlers.LoggingHandler(os.Stdout, mx)

	//to handle URL like
	//http://website:8080/person/Boo/CEO/199

	//http://website:8080/person/Boo/CEO/199 <- if age > 199, will cause 404 error

	mx.HandleFunc("/person/{name}/{job}/{age:[0-199]+}", ProcessPathVariables)

	http.ListenAndServe(":8080", loggedRouter)
}
