package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			bouttonAbout := r.FormValue("boutton-about")
			if bouttonAbout != "" {

				http.Redirect(w, r, "/"+bouttonAbout, http.StatusSeeOther)
				return
			}
			bouttonBlog := r.FormValue("boutton-blog")
			if bouttonBlog != "" {

				http.Redirect(w, r, "/"+bouttonBlog, http.StatusSeeOther)

				return
			}
			bouttonServices := r.FormValue("boutton-services")
			if bouttonServices != "" {

				http.Redirect(w, r, "/"+bouttonServices, http.StatusSeeOther)
				return
			}

			bouttonEvents := r.FormValue("boutton-events")
			if bouttonEvents != "" {

				http.Redirect(w, r, "/"+bouttonEvents, http.StatusSeeOther)
				return
			}

			bouttonContact := r.FormValue("boutton-contact")
			if bouttonContact != "" {

				http.Redirect(w, r, "/"+bouttonContact, http.StatusSeeOther)
				return

			}

			Nom := r.FormValue("full_name")
			Phone := r.FormValue("phone_number")
			Email := r.FormValue("email")
			Message := r.FormValue("message")

			// Se connecter à la base de données MySQL
			db, err := sql.Open("mysql", "kedithre:@1996Edith@tcp(127.0.0.1:3306)/edith")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer db.Close()

			err = db.Ping()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			insertQuery, err := db.Query("INSERT INTO edith.messages (`full_name`, `phone_number`, `email`, `message`) VALUES (?, ?, ?, ?);", Nom, Phone, Email, Message)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer insertQuery.Close()
			fmt.Println("Connexion réussie à la base de données")

			http.Redirect(w, r, "/index.html", http.StatusSeeOther)
			return
		}

		tmpl, err := template.ParseFiles("index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("/about.html", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("about.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	http.HandleFunc("/services.html", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("services.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	http.HandleFunc("/events.html", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("events.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	http.HandleFunc("/blog.html", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("blog.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("/contact.html", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("contact.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	fmt.Println("http://localhost:8888")
	http.ListenAndServe(":8888", nil)
}
