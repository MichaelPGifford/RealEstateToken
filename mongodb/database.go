package mongodb

import(
	"gopkg.in/mgo.v2"
	
	"fmt"
	
)
const(
	SERVER = "localhost:27017"
	DBNAME = "maps"
)

func Connect() *mgo.Session {
	sess, err := mgo.Dial(SERVER)
	if err != nil {
		fmt.Println("Database Connection Error")
		//panic(err)
	}
	return sess
}
