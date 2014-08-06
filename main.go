package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
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

func main() {

	// Nome:{Primeiro:"teste"},

	atleta := Atleta{Nome: Nome{Primeiro: "teste", NomeDoMeio: "do", Sobrenome: "teste"}, Email: "teste@teste"}

	atleta.Id = bson.NewObjectId()

	// atleta.Email = "teste@teste"

	// Nome:{Primeiro:"teste"},
	// {Email:}

	fmt.Println(atleta)

	sess, err := mgo.Dial("mongodb://mega:megamega@kahana.mongohq.com:10089/MegaRunning")
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		//	    os.Exit(1)
	}
	defer sess.Close()

	collection := sess.DB("MegaRunning").C("Atleta")

	err = collection.Insert(&atleta)
	//	go get gopkg.in/mgo.v2
	fmt.Println(err)

}
