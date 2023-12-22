package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

const port = ":5050"

var tmpl *template.Template

func accueilHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		tmpl.ExecuteTemplate(w, "accueil.html", nil)
	} else if r.URL.Path == "/inscription" {
		http.Redirect(w, r, "/inscription.html", http.StatusSeeOther)
	} else if r.URL.Path == "/inscription.html" {
		tmpl.ExecuteTemplate(w, "inscription.html", nil)
	} else if r.URL.Path == "/connexion" {
		http.Redirect(w, r, "/connexion.html", http.StatusSeeOther)
	} else if r.URL.Path == "/connexion.html" {
		tmpl.ExecuteTemplate(w, "inscription.html", nil)
	} else {
		log.Println("Erreur lors du chargement")
	}
}
func connexionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl.ExecuteTemplate(w, "connexion.html", nil)
	}
}

func inscriptionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Utilisez la variable globale tmpl
		tmpl.Execute(w, nil)
	} else if r.Method == "POST" {

		user_name := r.FormValue("name")
		user_email := r.FormValue("mail")
		user_password := r.FormValue("pwd")

		db, err := sql.Open("sqlite3", "/home/digifemmes-22lab025/Projet/forum/forum1.db")
		if err != nil {
			log.Println("ERREUR LORS DE L'OUVERTURE DE LA BASE DE DONNEE:", err)
			http.Error(w, "Erreur lors de l'ouverture de la base de données", http.StatusInternalServerError)
			return
		}
		defer db.Close()

		// Vérifier la connexion à la base de données
		err = db.Ping()
		if err != nil {
			log.Println("ERREUR LORS DE LA CONNEXION A LA BASE DE DONNEE:", err)
			http.Error(w, "Erreur lors de la connexion à la base de données", http.StatusInternalServerError)
			return
		}

		// Exécuter la requête d'insertion
		_, err = db.Exec(`INSERT INTO users (name, mail, pwd)
                            VALUES (?, ?, ?)`,
			user_name, user_email, user_password)
		if err != nil {
			log.Println("ERREUR LORS DE L'INSERTION DANS LA BASE DE DONNEE:", err)
			http.Error(w, "Erreur lors de l'insertion dans la base de données", http.StatusInternalServerError)
			return
		}

		log.Println("Nouvel utilisateur inséré avec succès")

		http.Redirect(w, r, "/connexion.html", http.StatusSeeOther)

	} else {
		fmt.Fprint(w, "METHODE NON PRISE EN CHARGE")
	}
}

func main() {
	templateDir := "./templates"
	tmpl = template.Must(template.ParseGlob(filepath.Join(templateDir, "*.html")))

	http.HandleFunc("/", accueilHandler)
	http.HandleFunc("/inscription", inscriptionHandler)
	http.HandleFunc("/connexion", connexionHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/image/", http.StripPrefix("/image/", http.FileServer(http.Dir("image"))))

	fmt.Println("(http://localhost:5050) - server started on port", port)

	http.ListenAndServe(port, nil)
}
