package main

import (
	"fmt"
	"go-my-resume/app"
	"go-my-resume/internal/pkg/workexperience"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("internal/templates/*"))
}

func main() {
	fsStatic := http.FileServer(http.Dir("static"))
	http.Handle("/css/", fsStatic)
	http.HandleFunc("/", index)
	http.HandleFunc("/workexperience", workExperience)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func workExperience(w http.ResponseWriter, req *http.Request) {
	session, err := app.GetMongoSession()
	if err != nil {
		fmt.Println("Cannot get mongo session", err)
	}
	defer session.Close()

	workExps, err := workexperience.InquiryWorkExperience(session)
	tpl.ExecuteTemplate(w, "workexperience.gohtml", workExps)
}
