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
	router.GET("/Balance", handlers.BalanceHandler)

	router.GET("/css/*cssfile", handlers.CSS_Handler)
	router.GET("/js/*jsfile", handlers.JS_Handler)
	router.GET("/images/*imgfile", handlers.Images_Handler)

	//Manage routing of API calls
	router.GET("/api/get-pins", api.Get_pins)
	router.POST("/api/add-pin", api.Add_Pin)
	//router.GET("/api/generate-wallet", api.GenerateWallet)
	router.POST("/api/sign-up", api.SignUp)
	router.POST("/api/sign-in", api.SignIn)
	router.GET("/api/verify-id", api.VerifyId)
	router.GET("/api/ether-balance", api.EtherBalance)
	router.POST("/api/send-ether", api.SendEther)

	//Manage Routing of Admin Pages
	router.GET("/admin/add-property", handlers.Admin_Add_Pins)
	//Serve the Webserver to Users
	log.Fatal(http.ListenAndServe(":8080", router))
}
