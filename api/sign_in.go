package api
import(
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/zcaps/maps/user"
	"encoding/json"
	"fmt"
)
func SignIn(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	var user user.User
	json.NewDecoder(r.Body).Decode(&user.Info)
	
	user.Login(w,r)
	fmt.Println(user)
}