package prova

import (
	"../conexao"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

var collection *mgo.Collection = conexao.GetDb().C("Prova")

func Handler(response http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Method)

	switch request.Method {
	case "GET":
		provas := Provas()
		fmt.Println(provas)
		js, _ := json.MarshalIndent(provas, " ", "   ")
		response.Write(js)
		//resp := json.NewEncoder(response)
		//resp.Encode(provas)

	case "POST":
		var prova Prova
		json.NewDecoder(request.Body).Decode(&prova)
		err := Insere(&prova)
		fmt.Println(err)

		//case "DELETE":

	}
}

func HandlerId(response http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Method)

	switch request.Method {
	case "GET":
		vars := mux.Vars(request)
		id := bson.ObjectIdHex(vars["id"])

		prova := GetProva(id)
		fmt.Println(prova)
		js, _ := json.MarshalIndent(prova, " ", "   ")
		response.Write(js)
		//resp := json.NewEncoder(response)
		//resp.Encode(provas)
		/*
			case "PUT":
				var prova Prova
				json.NewDecoder(request.Body).Decode(&prova)
				err := Insere(&prova)
				fmt.Println(err)
		*/

		//case "DELETE":

	}
}

func Provas() []Prova {
	var provas []Prova
	collection.Find(nil).All(&provas)
	return provas
}

func GetProva(id bson.ObjectId) Prova {
	var prova Prova
	collection.FindId(id).One(&prova)
	return prova
}

func Insere(prova *Prova) error {

	prova.Id = bson.NewObjectId()
	fmt.Println(prova)
	return collection.Insert(prova)

}
