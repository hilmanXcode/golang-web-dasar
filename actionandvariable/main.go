package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Info struct {
	Affiliation string
	Address     string
}

type Person struct {
	Name    string
	Gender  string
	Hobbies []string
	Info    Info
}

func main() {

	var address = "localhost:1337"
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		var data = Person{
			Name:    "manz",
			Gender:  "Laki-Laki",
			Hobbies: []string{"Ngoding", "Mancing"},
			Info:    Info{Affiliation: "Juned", Address: "Karawang"},
		}

		tmpl := template.Must(template.ParseFiles(
			"views/index.html",
		))

		err := tmpl.ExecuteTemplate(w, "index", data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	server := new(http.Server)
	server.Addr = address
	fmt.Println("Server run at", address)
	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}

func (t Info) GetAffiliationDetailInfo() string {
	return "Have 31 divisions"
}
