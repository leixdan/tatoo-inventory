package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/joho/godotenv"
)

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	// Cargar variables de entorno desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Println("No se pudo cargar archivo .env, asegúrate que exista")
	}

	// Inicializar base de datos
	err = InitDB()
	if err != nil {
		log.Fatalf("Error abriendo DB: %v", err)
	}

	// Servir archivos estáticos en /static/
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Rutas CRUD
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/new", newHandler)
	http.HandleFunc("/create", createHandler)
	http.HandleFunc("/edit", editHandler)
	http.HandleFunc("/update", updateHandler)
	http.HandleFunc("/delete", deleteHandler)

	log.Println("Servidor iniciado en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Enviar slice vacío para evitar problemas con datos
	materials := []Material{}

	err := tmpl.ExecuteTemplate(w, "index.html", materials)
	if err != nil {
		http.Error(w, "Error ejecutando plantilla: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func newHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "new.html", nil)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error en el formulario", http.StatusBadRequest)
		return
	}
	qty, _ := strconv.Atoi(r.FormValue("quantity"))
	m := Material{
		Name:        r.FormValue("name"),
		Quantity:    qty,
		Description: r.FormValue("description"),
	}
	err = CreateMaterial(m)
	if err != nil {
		http.Error(w, "Error creando material", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	m, err := GetMaterialByID(id)
	if err != nil {
		http.Error(w, "Material no encontrado", http.StatusNotFound)
		return
	}
	tmpl.ExecuteTemplate(w, "edit.html", m)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error en el formulario", http.StatusBadRequest)
		return
	}
	id, _ := strconv.Atoi(r.FormValue("id"))
	qty, _ := strconv.Atoi(r.FormValue("quantity"))
	m := Material{
		ID:          id,
		Name:        r.FormValue("name"),
		Quantity:    qty,
		Description: r.FormValue("description"),
	}
	err = UpdateMaterial(m)
	if err != nil {
		http.Error(w, "Error actualizando material", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	err = DeleteMaterial(id)
	if err != nil {
		http.Error(w, "Error eliminando material", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
