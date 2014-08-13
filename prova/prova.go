package prova

import(
"net/http"
"gopkg.in/mgo.v2"
"gopkg.in/mgo.v2/bson"
"encoding/json"
"../conexao"
"fmt"
)

var collection *mgo.Collection = conexao.GetDb().C("Prova")

func Handler(response http.ResponseWriter , request *http.Request){
	fmt.Println(request.Method)
	switch request.Method {
		//case "GET" : 

		case "POST" :
			var prova Prova 
			json.NewDecoder(request.Body).Decode(&prova)
			err := Insere(&prova)
			fmt.Println(err)

		//case "DELETE" :

	}	
}

func Insere(prova *Prova) error{

    prova.Id = bson.NewObjectId()
    fmt.Println(prova)
	return collection.Insert(prova)
	
}