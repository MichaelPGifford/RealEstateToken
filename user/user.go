package user
import(
	"net/http"
	"github.com/zcaps/maps/user/person"
	"github.com/zcaps/maps/user/wallet"
	"gopkg.in/mgo.v2/bson"
	"github.com/zcaps/maps/mongodb"
	"github.com/zcaps/maps/cookies"
	"fmt"
)
type User struct {
	Id            bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Info          *person.Person `json:"info"`
	Wallet 		  *wallet.Wallet
}
func (u *User)CreateUser()string{
	if(u.Info.UsernameExists() == false){
		//username Exists in DB
		return "An account is already registered with this email address."
	}
	if(u.Info.PasswordCheck() == false){
		//password is either too short or doesn't match
		return "Your password must be atleast 8 characters and the two passwords must match."
	}
	//Encrypt password
	u.Info.Encrypt()

	//Create Id
	u.Id = bson.NewObjectId()

	//Create Wallet
	u.Wallet = wallet.GenerateWallet()

	sess := mongodb.Connect()
	defer sess.Close()
	//store encrypted user into the DB
	sess.DB("maps").C("Users").Insert(u)
	return ""
}

////
//	Compares the hash and given password to see if they match
//	Stores an encrypted user ID as a cookie
////
func (u *User) Login(w http.ResponseWriter, r *http.Request) bool {

	if u.Info.Decrypt() {
		//make a cookie with encrypted user id
		//password is good
		u.FindId()
		cookies.StoreId(w, r, u.Id.Hex())
		return true

	} else {
		//password is bad
		return false
	}
}

//if the ID is the only thing stored in a User struct, this will populate the Object with data
func (u *User) FindId() {
	sess := mongodb.Connect()
	defer sess.Close()
	sess.DB("maps").C("Users").Find(bson.M{"info.email": u.Info.Email}).One(u)
}
func GetUserFromCookie(w http.ResponseWriter, r *http.Request) *User {
	id := cookies.GetId(w, r)
	if id == "" {
		return nil
	}
	sess := mongodb.Connect()
	defer sess.Close()
	objId := bson.ObjectIdHex(id)
	fmt.Println(objId)
	u := &User{}
	sess.DB("maps").C("Users").Find(bson.M{"_id": objId}).One(u)
	fmt.Println(u)
	return u
}