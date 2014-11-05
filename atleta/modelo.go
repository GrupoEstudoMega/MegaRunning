package atleta

import (
//"gopkg.in/mgo.v2/bson"
)

type Atleta struct {
	//Id    bson.ObjectId `bson:"_id" json:"id"`
	Nome  string `bson:"nome"`
	Email string `bson:"email"`
}
