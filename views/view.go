package views

import (
	"fmt"
	"html/template"
)

func NewView(files ...string) *View {
	files = append(files, "views/layouts/top.html", "views/layouts/footer.html")

	fmt.Println(files)

	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
	}
}

type View struct {
	Template *template.Template
}
