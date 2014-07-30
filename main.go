package main

import( 
    "fmt"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

type Nome struct{
		primeiro string `bson:"primeiro"`
		nomeDoMeio string `bson:"nomeDoMeio"`
		sobrenome string `bson:"sobrenome"`
	}

type Atleta struct{
	nome Nome `bson:"nome"`
	email string `bson:"email"`
}



func main(){

	// Nome:{Primeiro:"teste"},

	atleta := Atleta{nome:Nome{primeiro:"teste",nomeDoMeio:"do",sobrenome:"teste"},email:"teste@teste"}
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

    err = collection.Insert(bson.M(atleta))
//	go get gopkg.in/mgo.v2
    fmt.Println(err)

}