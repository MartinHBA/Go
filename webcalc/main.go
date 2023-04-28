package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Calculation struct {
	Num1   int
	Num2   int
	Result int
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/calculate", calculateHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}
	t.Execute(w, nil)
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		num1, _ := strconv.Atoi(r.FormValue("num1"))
		num2, _ := strconv.Atoi(r.FormValue("num2"))
		result := num1 + num2

		data := Calculation{
			Num1:   num1,
			Num2:   num2,
			Result: result,
		}

		t, err := template.ParseFiles("result.html")
		if err != nil {
			fmt.Println("Error parsing template:", err)
			return
		}
		t.Execute(w, data)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
