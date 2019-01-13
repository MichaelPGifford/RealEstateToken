package api
import(
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/zcaps/maps/user"
	"encoding/json"
)
func EtherBalance(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	u := user.GetUserFromCookie(w,r)
	if u != nil {
		balance := u.Wallet.GetBalance()
		json.NewEncoder(w).Encode(balance)
	}
}