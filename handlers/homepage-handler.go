package handlers

import(
	"log"
	"net/http"
	"html/template"
	"path/filepath"
	"github.com/julienschmidt/httprouter"
	"github.com/zcaps/mich.ly/cookies"

	"io/ioutil"
	"strings"
	
)


func Homepage_Handler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	cookies.Visitor(w,r)
	absPath, _ := filepath.Abs("html/map.gohtml")
	
	tpl, err := template.ParseFiles(absPath)
	if err != nil {
		log.Fatalln(err)
	}

	//get the html data from mongodb
	

	err = tpl.ExecuteTemplate(w, "homepage", nil)
	if err != nil {
		log.Fatalln(err)
	}
}


func CSS_Handler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	absPath, _ := filepath.Abs("css/")
	file := absPath + ps.ByName("cssfile")
	data, err := ioutil.ReadFile(file)

	if err == nil {
		var contentType string
		if strings.HasSuffix(file, ".css") {
		    contentType = "text/css"
		} else if strings.HasSuffix(file, ".html") {
		    contentType = "text/html"
		} else if strings.HasSuffix(file, ".js") {
		    contentType = "application/javascript"
		} else if strings.HasSuffix(file, ".png") {
		    contentType = "image/png"
		} else if strings.HasSuffix(file, ".svg") {
		    contentType = "image/svg+xml"
		} else {
		    contentType = "text/plain"
		}

		w.Header().Set("Content-Type", contentType)
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 - Something went wrong - " + http.StatusText(404)))
	}
}

func JS_Handler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	absPath, _ := filepath.Abs("js/")
	file := absPath + ps.ByName("jsfile")
	data, err := ioutil.ReadFile(file)

	if err == nil {
		var contentType string
		if strings.HasSuffix(file, ".css") {
		    contentType = "text/css"
		} else if strings.HasSuffix(file, ".html") {
		    contentType = "text/html"
		} else if strings.HasSuffix(file, ".js") {
		    contentType = "application/javascript"
		} else if strings.HasSuffix(file, ".png") {
		    contentType = "image/png"
		} else if strings.HasSuffix(file, ".svg") {
		    contentType = "image/svg+xml"
		} else {
		    contentType = "text/plain"
		}

		w.Header().Set("Content-Type", contentType)
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 - Something went wrong - " + http.StatusText(404)))
	}
}

func Images_Handler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	absPath, _ := filepath.Abs("images/")
	file := absPath + ps.ByName("imgfile")
	data, err := ioutil.ReadFile(file)

	if err == nil {
		var contentType string
		if strings.HasSuffix(file, ".css") {
		    contentType = "text/css"
		} else if strings.HasSuffix(file, ".html") {
		    contentType = "text/html"
		} else if strings.HasSuffix(file, ".js") {
		    contentType = "application/javascript"
		} else if strings.HasSuffix(file, ".png") {
		    contentType = "image/png"
		} else if strings.HasSuffix(file, ".svg") {
		    contentType = "image/svg+xml"
		} else {
		    contentType = "text/plain"
		}

		w.Header().Set("Content-Type", contentType)
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 - Something went wrong - " + http.StatusText(404)))
	}
}
