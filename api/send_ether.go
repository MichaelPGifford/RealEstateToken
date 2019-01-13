package api
import(
	"github.com/julienschmidt/httprouter"
	"net/http"
	"encoding/json"
	"github.com/zcaps/maps/user"
	//"fmt"
)
func SendEther(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	var data map[string] interface{}
	
	json.NewDecoder(r.Body).Decode(&data)

	if(data["to"] != nil && data["amount"] != nil){
		u := user.GetUserFromCookie(w,r)
		if u != nil {
			u.Wallet.SendEth(data["to"].(string), data["amount"].(string))
		}
	}
	

	
}