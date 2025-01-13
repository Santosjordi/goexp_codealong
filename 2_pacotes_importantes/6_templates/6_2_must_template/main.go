package main

// CSS doesn't work yet

import (
	"html/template"
	"net/http"
)

type Course struct {
	Name     string
	Price    int
	Duration int
}

type Courses []Course

func (t *Course) Mod(x, y int) int {
	return x % y
}

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("content.html").ParseFiles(templates...))
		err := t.Execute(w, Courses{
			{"Golang", 100, 10},
			{"Python", 99, 20},
			{"Java", 20, 160},
		})
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8082", nil)

}
