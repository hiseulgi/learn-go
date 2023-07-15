package controllers

import (
	"html/template"
	"net/http"
	"web-app/structs"
)

var funcMap = template.FuncMap{
	"unescape": func(s string) template.HTML {
		return template.HTML(s)
	},
	"avg": func(n ...int) int {
		var total = 0
		for _, each := range n {
			total += each
		}
		return total / len(n)
	},
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	data := structs.M{"name": "Sugab"}
	tmpl := template.Must(template.ParseFiles(
		"views/template/index.html",
		"views/template/_header.html",
		"views/template/_message.html",
	))

	if err := tmpl.ExecuteTemplate(w, "index", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	data := structs.M{"name": "Sugab"}
	tmpl := template.Must(template.ParseFiles(
		"views/template/about.html",
		"views/template/_header.html",
		"views/template/_message.html",
	))

	if err := tmpl.ExecuteTemplate(w, "about", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func SuperHeroHandler(w http.ResponseWriter, r *http.Request) {
	hero := structs.Superhero{
		Name:    "Sugab Slayer",
		Alias:   "Demon",
		Friends: []string{"Batman", "Flash", "Spiderman"},
	}

	tmpl := template.Must(template.ParseFiles("views/template/superhero.html"))
	if err := tmpl.ExecuteTemplate(w, "superhero", hero); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CustomFunc(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("views/template/customfunc.html").Funcs(funcMap).ParseFiles("views/template/customfunc.html"))

	if err := tmpl.ExecuteTemplate(w, "customfunc", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RenderIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("index").ParseFiles("views/template/specific.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RenderTest(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("test").ParseFiles("views/template/specific.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RenderHTML(w http.ResponseWriter, r *http.Request) {
	const view string = `<html>
		<head>
			<title>Template</title>
		</head>
		<body>
			<h1>Hello</h1>
		</body>
	</html>`

	tmpl := template.Must(template.New("main-template)").Parse(view))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	person := structs.Person{
		Name:    "Muhammad Bagus",
		Gender:  "male",
		Hobbies: []string{"Anime", "Game"},
		Info: structs.Info{
			Affiliation: "Baskoro",
			Address:     "Tembalang City"},
	}

	tmpl := template.Must(template.ParseFiles("views/template/profile.html"))
	if err := tmpl.ExecuteTemplate(w, "profile", person); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
