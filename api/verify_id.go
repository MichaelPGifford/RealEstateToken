package api
import(
	"github.com/julienschmidt/httprouter"
	"net/http"
	"encoding/json"
	"github.com/zcaps/maps/user"
)
func VerifyId(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	u := user.GetUserFromCookie(w,r)
	if u != nil {
		json.NewEncoder(w).Encode(u)
	}
}