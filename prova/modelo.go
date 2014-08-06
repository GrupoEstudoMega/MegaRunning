package prova

import (
	"../atleta"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Prova struct {
	Id            bson.ObjectId `bson:"_id" json:"id"`
	Nome          string        `bson:"nome"`
	Url           string        `bson:"url"`
	Local         string        `bson:"local"`
	Data          time.Time     `bson:"data"`
	Valor         float32       `bson:"valor"`
	Participantes []Participacao
}

type Participacao struct {
	ValorPatrocinio float32 `bson:"valorpatrocinio"`
	Atleta          atleta.Atleta
}
