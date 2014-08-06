package atleta

import (
	"gopkg.in/mgo.v2/bson"
)

type Nome struct {
	Primeiro   string `bson:"primeiro"`
	NomeDoMeio string `bson:"nomeDoMeio"`
	Sobrenome  string `bson:"sobrenome"`
}

type Atleta struct {
	Id    bson.ObjectId `bson:"_id" json:"id"`
	Nome  Nome          `bson:"nome"`
	Email string        `bson:"email"`
}
