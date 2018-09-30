package main

import(
	"net/http"
	"github.com/zcaps/maps/handlers"
	"github.com/zcaps/maps/api"
	"github.com/julienschmidt/httprouter"
	"log"
	
)

func main() {
	router := httprouter.New()
	//Manage the routing of user webpages
	router.GET("/", handlers.Homepage_Handler)

	router.GET("/css/*cssfile", handlers.CSS_Handler)
	router.GET("/js/*jsfile", handlers.JS_Handler)
	router.GET("/images/*imgfile", handlers.Images_Handler)

	//Manage routing of API calls
	router.GET("/api/get-pins", api.Get_pins)
	router.POST("/api/add-pin", api.Add_Pin)

	//Manage Routing of Admin Pages
	router.GET("/admin/add-property", handlers.Admin_Add_Pins)
	//Serve the Webserver to Users
	log.Fatal(http.ListenAndServe(":8080", router))
}
