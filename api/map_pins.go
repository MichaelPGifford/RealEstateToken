package api

import(
	"github.com/zcaps/maps/mongodb"
	"github.com/zcaps/maps/mappin"
	"github.com/zcaps/maps/tokens"
	"fmt"
	"net/http"
	"strings"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
)

type location struct {
	Address string `json: "address"`
	Price string `json: "price"`
	Sqft string `json: "sqft"`
	Class string `json: "class"`
	Latitude float64 `json: "latitude"`
	Longitude float64 `json: "longitude"`
}

func Get_pins(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	sess := mongodb.Connect()
	defer sess.Close()
	
	var locations []mappin.Pin
	collect := sess.DB(mongodb.DBNAME).C("pins")
	err := collect.Find(nil).All(&locations)
	if err != nil {
		fmt.Println("Failed to read data from database")
	}
	fmt.Println(locations)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(locations)
	
}

func Add_Pin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	sess := mongodb.Connect()
	defer sess.Close()
	var loca mappin.Pin
	json.NewDecoder(r.Body).Decode(&loca)
	fmt.Println(loca);
	symb := strings.Fields(loca.Address)
	address := tokens.DeployContract(symb[0], loca.Address)
	loca.ContractAddress = string(address);
	sess.DB(mongodb.DBNAME).C("pins").Insert(loca)
	
	
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(loca)
}
