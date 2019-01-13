package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/zcaps/maps/user/wallet"
)

//Generates a crypto wallet for a user
func GenerateWallet(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	wallet.GenerateWallet()
}
