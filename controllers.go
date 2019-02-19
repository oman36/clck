package main

import (
	"fmt"
	"net/http"
	"strings"
)

func postHandler(w http.ResponseWriter, r *http.Request) {
	var validationError string
	value := strings.Trim(r.PostFormValue("url"), " ")
	if value == "" {
		validationError = "Нобходимо указать url"
	} else if valid, err := urlIsValid(value); !valid {
		validationError = "Некоректный url"
	} else if err != nil {
		panic(err)
	}
	if validationError != "" {
		renderMainPage(w, MainPageData{Error: validationError})
		return
	}

	url := Urls{}
	db.First(&url, "url = ?", value)

	if url.ID == 0 {
		url.Url = value
		db.Create(&url)
	}
	id := url.ID
	converter := NewShortUriConverter()
	uri, err := converter.IntToUri(id)
	if err != nil {
		panic(err)
	}
	renderSuccessPage(w, SuccessPageData{
		Url: fmt.Sprintf("https://%s/%s", getConfig().Host, uri),
	})

}

func getHandler(w http.ResponseWriter, r *http.Request) {
	var link = r.URL.Path[1:]
	if link == "" {
		renderMainPage(w, MainPageData{Error: ""})
	} else {
		valid, err := shortUrlIsValid(link)
		if err != nil {
			panic(err)
		}
		if !valid {
			http.NotFound(w, r)
			return
		}
		converter := NewShortUriConverter()
		id, err := converter.UriToInt(link)
		if err != nil {
			panic(err)
		}

		url := Urls{}
		db.First(&url, id)
		if url.ID == 0 {
			http.NotFound(w, r)
			return
		}
		http.Redirect(w, r, url.Url, 302)
	}
}
