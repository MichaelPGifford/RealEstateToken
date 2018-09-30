package handlers

import(
	"log"
	"net/http"
	"html/template"
	"path/filepath"
	"github.com/julienschmidt/httprouter"
	
)

func Admin_Add_Pins(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	absPath, _ := filepath.Abs("html/admin-add-pins.gohtml")
	
	tpl, err := template.ParseFiles(absPath)
	if err != nil {
		log.Fatalln(err)
	}

	//get the html data from mongodb
	

	err = tpl.ExecuteTemplate(w, "admin-add-pins", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
