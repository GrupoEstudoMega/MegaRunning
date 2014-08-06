package main

import (
	"./atleta"
	"./conexao"
	"./prova"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func main() {

	joao := atleta.Atleta{Nome: atleta.Nome{Primeiro: "teste", NomeDoMeio: "do", Sobrenome: "teste"}, Email: "teste@teste"}

	joao.Id = bson.NewObjectId()

	corrida := prova.Prova{Nome: "Corrida Noturna"}

	corrida.Id = bson.NewObjectId()

	compareceu := prova.Participacao{ValorPatrocinio: 12.3, Atleta: joao}

	corrida.Participantes = append(corrida.Participantes, compareceu)

	// atleta.Email = "teste@teste"

	// Nome:{Primeiro:"teste"},
	// {Email:}

	fmt.Println(corrida)

	sess, err := conexao.GetSessao()
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		//	    os.Exit(1)
	}
	defer sess.Close()

	//	collection := sess.DB("MegaRunning").C("Atleta")
	provas := sess.DB("MegaRunning").C("Prova")

	err = provas.Insert(&corrida)
	//	go get gopkg.in/mgo.v2
	fmt.Println(err)

}

func insereTestes() {
	joao := atleta.Atleta{Nome: atleta.Nome{Primeiro: "teste", NomeDoMeio: "do", Sobrenome: "teste"}, Email: "teste@teste"}

	//	joao.Id = bson.NewObjectId()

	corrida := prova.Prova{Nome: "Corrida Noturna"}

	corrida.Id = bson.NewObjectId()

	compareceu := prova.Participacao{ValorPatrocinio: 12.3, Atleta: joao}

	corrida.Participantes = append(corrida.Participantes, compareceu)

	// atleta.Email = "teste@teste"

	// Nome:{Primeiro:"teste"},
	// {Email:}

	fmt.Println(corrida)

	sess, err := mgo.Dial("mongodb://mega:megamega@kahana.mongohq.com:10089/MegaRunning")
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		//	    os.Exit(1)
	}
	defer sess.Close()

	//	collection := sess.DB("MegaRunning").C("Atleta")
	provas := sess.DB("MegaRunning").C("Prova")

	err = provas.Insert(&corrida)
	//	go get gopkg.in/mgo.v2
	fmt.Println(err)

}

//func getAtleta(email string) atleta.Atleta {

//}
