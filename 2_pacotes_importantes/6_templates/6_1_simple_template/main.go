package main

import (
	"html/template"
	"os"
)

type Course struct {
	Name     string
	Price    int
	Duration int
}

func main() {
	course := Course{"Golang", 100, 10}
	tmp := template.New("CourseTemplate")
	tmp, _ = tmp.Parse("Course Name: {{.Name}}\nPrice: ${{.Price}}\nDuration: {{.Duration}} hours\n")
	err := tmp.Execute(os.Stdout, course)
	if err != nil {
		panic(err)
	}
}
