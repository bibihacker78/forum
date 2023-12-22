package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func renderTemplates(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("./templates/" + tmpl + ".html")
	if err != nil {
		fmt.Fprint(w, "MODELE INTROUVABLE")
	}
	t.Execute(w, nil)
}

func inscription(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		renderTemplates(w, "inscription")
	case "POST":
		nom := r.FormValue("nom")
		prenom := r.FormValue("prenom")
		email := r.FormValue("email")
		password_hash := r.FormValue("mdp")

		db, err := sql.Open("sqlite3", "/home/digifemmes-22lab007/Projet/Tuto-Forum/forum.db")
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
		result, err := db.Exec(`INSERT INTO users (nom, prenom, email, username, password_hash)
								VALUES (?, ?, ?, ?, ?)`,
			nom, prenom, email, password_hash)
		if err != nil {
			fmt.Println("ERREUR LORS DE L'INSERTION DANS LA BASE DE DONNE")

		}

		// Récupérer l'ID de la ligne insérée
		id, err := result.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Nouvel utilisateur inséré avec l'ID: %d\n", id)
		http.Redirect(w, r, "http://localhost:9090/connexion", http.StatusSeeOther)
	default:
		fmt.Fprint(w, "METHODE NON PRIS EN CHARGE")

	}
}
