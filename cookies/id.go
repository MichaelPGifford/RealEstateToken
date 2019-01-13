package cookies

import (
	"net/http"

	"github.com/gorilla/sessions"
)

//salt for the cookie encryption
var store = sessions.NewCookieStore([]byte("098219340184347589270937498714098578148723412384321939847219088145972271106311"))

//stores an id for whoever makes this request
func StoreId(w http.ResponseWriter, r *http.Request, id string) {
	session, _ := store.Get(r, "id")
	session.Values["id"] = id
	session.Save(r, w)
}

//grabs the encrypted id and decrypts it
func GetId(w http.ResponseWriter, r *http.Request) string {
	session, _ := store.Get(r, "id")
	if session.Values["id"] == nil {
		return ""
	}
	return session.Values["id"].(string)

}
