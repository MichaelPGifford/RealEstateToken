package person
import(
	"gopkg.in/mgo.v2/bson"
	"github.com/zcaps/maps/mongodb"
	"golang.org/x/crypto/bcrypt"
)
type Person struct {
	
	FName string `json:"fname"`
	LName string `json:"lname"`
	Email string `json:"email"`
	Password string `json:"password"`
	CPassword string `json:omitempty`

}
////
//	UsernameExists returns TRUE if a username DOES NOT EXIST
////
func (p *Person) UsernameExists() bool{
	sess := mongodb.Connect()
	defer sess.Close()
	
	count, err := sess.DB("maps").C("Users").Find(bson.M{ "info.email": p.Email }).Count()
	
	if err != nil {
		return false
	}
	if count >= 1 {
		return false
	}
	return true
}
////
//	PasswordCheck returns TRUE if a password passes all error checks
////
func (p *Person) PasswordCheck() bool{
	check := true
	if(p.Password != p.CPassword){
		check = false
	}
	if(len(p.Password) < 8){
		check = false
	}
	p.CPassword = ""
	return check
}
////
//	Encrypt
////
func (p *Person) Encrypt() {
	hash, _ := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
	p.Password = string(hash)
	
}

////
//	Decrypt Compares User input to the encrypted password
////
func (p *Person) Decrypt() bool{
	pass := p.FindEncryptedPassword()
	err := bcrypt.CompareHashAndPassword([]byte(pass), []byte(p.Password))
	
	if err != nil {
		return false
	}
	return true
}
//Finds the encrypted password for a specified User.Info object and returns it
func (p *Person) FindEncryptedPassword() string{
	var user map[string] interface{}
	sess := mongodb.Connect()
	defer sess.Close()
	err := sess.DB("maps").C("Users").Find(bson.M{"info.email": p.Email}).Select(bson.M{"info.password": 1, "_id": 0}).One(&user)
	if err != nil {
		return ""
	}
	person := user["info"].(map[string] interface{})
	password := person["password"]
	
	return password.(string)
}