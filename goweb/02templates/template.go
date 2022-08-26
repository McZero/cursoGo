package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Estructuras
type Usuario struct {
	UserName string
	Edad     int
	Activo   bool
	Admin    bool
	Cursos   []Curso
}

type Curso struct {
	Nombre string
}

func Saludar(nombre string) string {
	return "Hola " + nombre + " desde una función"
}

// Handler
func Index(rw http.ResponseWriter, r *http.Request) {
	c1 := Curso{"Go"}
	c2 := Curso{"Python"}
	c3 := Curso{"HTML"}
	c4 := Curso{"CSS"}

	index, err := template.ParseFiles("index.html")

	cursos := []Curso{c1, c2, c3, c4}
	usuario := Usuario{"Carlos", 34, true, false, cursos}

	if err != nil {
		panic(err)
	} else {
		index.Execute(rw, usuario)
	}
}

func Saluda(rw http.ResponseWriter, r *http.Request) {
	funciones := template.FuncMap{
		"saludar": Saludar,
	}

	saluda, err := template.New("saluda.html").Funcs(funciones).ParseFiles("saluda.html")

	if err != nil {
		panic(err)
	} else {
		saluda.Execute(rw, nil)
	}
}

func main() {

	// Mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)
	mux.HandleFunc("/saluda", Saluda)

	// Crear servidor
	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}

	fmt.Println("El servidor esta corriendo en el puerto 3000")
	fmt.Println("Run server: http://localhost:3000/")
	// log.Fatal(http.ListenAndServe("localhost:3000", mux))
	log.Fatal(server.ListenAndServe())
}
