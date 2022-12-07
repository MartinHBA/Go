package main

import (
	"net/http"
	"text/template"
)

type Page struct {
	Message string
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			// Get the text input from the form.
			text := r.FormValue("text")

			//translate text
			translatedText := (translateToSK(text))

			// Set the message to be displayed on the page.
			p := Page{Message: translatedText}

			// Execute the HTML template with the message.
			t, _ := template.ParseFiles("form.html")
			t.Execute(w, p)
		} else {
			// Display the initial form if the request is not a POST request.
			t, _ := template.ParseFiles("form.html")
			t.Execute(w, nil)
		}
	})

	// Start the server and listen for requests on port 8080.
	http.ListenAndServe(":8080", nil)
}
