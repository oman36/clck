package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

const mainPageHtml = `<form method="POST" action="/">
    <label> Url:
        <input type="url" name="url">
    </label>
    <div style="color:red">{{.Error}}</div>
    <input type="submit">
</form>`
const successPageHtml = `<label> Короткая ссылка:
	<input type="url" name="url" readonly value="{{.Url}}">
</label>`

type MainPageData struct {
	Error string
}
type SuccessPageData struct {
	Url string
}

func renderTemplate(tmplString string, params interface{}) (string, error) {
	tpl, err := template.New("format").Parse(tmplString)
	if err != nil {
		return "", err
	}
	var buf = bytes.Buffer{}

	err = tpl.Execute(&buf, params)
	return buf.String(), err
}

func render(w http.ResponseWriter, tmplString string, params interface{}) (err error) {
	var html string
	html, err = renderTemplate(tmplString, params)
	if err != nil {
		return
	}
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	_, err = fmt.Fprintf(w, html)
	return
}

func renderMainPage(w http.ResponseWriter, data MainPageData) {
	err := render(w, mainPageHtml, data)
	if err != nil {
		panic(err)
	}
}
func renderSuccessPage(w http.ResponseWriter, data SuccessPageData) {
	err := render(w, successPageHtml, data)
	if err != nil {
		panic(err)
	}
}
