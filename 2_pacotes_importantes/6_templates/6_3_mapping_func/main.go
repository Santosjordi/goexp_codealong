package main

import (
	"os"
	"strings"
	"text/template"
)

type Course struct {
	Name     string
	Price    int
	Duration int
}

type Courses []Course

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	t := template.New("content.html")

	t.Funcs(template.FuncMap{"ToUpper": ToUpper})

	t = template.Must(t.ParseFiles(templates...))

	err := t.Execute(os.Stdout, Courses{
		{"Golang", 100, 10},
		{"Python", 99, 20},
		{"Java", 20, 160},
	})
	if err != nil {
		panic(err)
	}
}
