package main

import (
	"html/template"
	"log"
	"net/http"
)

// this struct represents the data that will be used to
// populate the form on the website
type InputData struct {
	Text    string
	Message string
}

func main() {
	// create a new http.Server and bind it to the default
	// http port (:80)
	server := &http.Server{
		Addr: ":8080",
	}

	// create a new HTTP handler for the "/" route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// check if this is a POST request
		if r.Method == "POST" {
			// parse the form data
			err := r.ParseForm()
			if err != nil {
				log.Println(err)
				return
			}

			// get the text input from the form
			text := r.FormValue("text")
			translation := translateToSK(text)
			// create a new InputData struct and populate it with
			// the user's input and a message to display
			data := InputData{
				Text:    translation,
				Message: "You entered: " + translation,
			}

			// execute the template and write the output to the
			// http.ResponseWriter
			tmpl, err := template.ParseFiles("form.html")
			if err != nil {
				log.Println(err)
				return
			}
			err = tmpl.Execute(w, data)
			if err != nil {
				log.Println(err)
				return
			}
			return
		}

		// if this is not a POST request, then render the form
		tmpl, err := template.ParseFiles("form.html")
		if err != nil {
			log.Println(err)
			return
		}

		// create a new InputData struct and populate it with
		// some default values
		data := InputData{
			Text:    "",
			Message: "",
		}

		// execute the template and write the output to the
		// http.ResponseWriter
		err = tmpl.Execute(w, data)
		if err != nil {
			log.Println(err)
			return
		}
	})

	// start the HTTP server
	log.Fatal(server.ListenAndServe())
}
