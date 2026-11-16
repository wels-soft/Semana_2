package main

import (
	"html/template"
	"net/http"
	"strconv"
)

type Persona struct {
	ID     int
	Nombre string
}

var baseDatos []Persona
var contadorID = 1

func leer(w http.ResponseWriter, r *http.Request) {
	template.Must(template.ParseFiles("index.html")).Execute(w, baseDatos)
}

func crear(w http.ResponseWriter, r *http.Request) {
	baseDatos = append(baseDatos, Persona{ID: contadorID, Nombre: r.FormValue("nombre")})
	contadorID++
	http.Redirect(w, r, "/", 303)
}

func actualizar(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	for i, p := range baseDatos {
		if p.ID == id {
			baseDatos[i].Nombre = r.FormValue("nuevo_nombre")
		}
	}
	http.Redirect(w, r, "/", 303)
}

func eliminar(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	for i, p := range baseDatos {
		if p.ID == id {
			baseDatos = append(baseDatos[:i], baseDatos[i+1:]...)
		}
	}
	http.Redirect(w, r, "/", 303)
}

func main() {
	http.HandleFunc("/", leer)
	http.HandleFunc("/crear", crear)
	http.HandleFunc("/actualizar", actualizar)
	http.HandleFunc("/eliminar", eliminar)
	http.ListenAndServe(":8080", nil)
}
