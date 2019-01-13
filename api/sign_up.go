package api
import(
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/zcaps/maps/user"
	"encoding/json"
	"fmt"
)
func SignUp(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	var user user.User
	json.NewDecoder(r.Body).Decode(&user.Info)
	
	fmt.Println(user.CreateUser())
	fmt.Println(user.Info)
}