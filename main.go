package main

import (
	"html/template"
	"net/http"
	"strconv"
)

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	InitDB()

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/new", newHandler)
	http.HandleFunc("/create", createHandler)
	http.HandleFunc("/edit", editHandler)
	http.HandleFunc("/update", updateHandler)
	http.HandleFunc("/delete", deleteHandler)

	http.ListenAndServe(":8080", nil)

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	materials, _ := GetAllMaterials()
	tmpl.ExecuteTemplate(w, "index.html", materials)
}

func newHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "new.html", nil)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	qty, _ := strconv.Atoi(r.FormValue("quantity"))
	m := Material{
		Name:        r.FormValue("name"),
		Quantity:    qty,
		Description: r.FormValue("description"),
	}
	CreateMaterial(m)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	m, _ := GetMaterialByID(id)
	tmpl.ExecuteTemplate(w, "edit.html", m)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id, _ := strconv.Atoi(r.FormValue("id"))
	qty, _ := strconv.Atoi(r.FormValue("quantity"))
	m := Material{
		ID:          id,
		Name:        r.FormValue("name"),
		Quantity:    qty,
		Description: r.FormValue("description"),
	}
	UpdateMaterial(m)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	DeleteMaterial(id)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
