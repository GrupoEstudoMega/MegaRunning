package conexao

import (
	"gopkg.in/mgo.v2"
)

var sessao *mgo.Session

func GetSessao() (*mgo.Session, error) {
	if sessao != nil {
		return sessao, nil
	} else {
		sessao, err := mgo.Dial("mongodb://mega:megamega@kahana.mongohq.com:10089/MegaRunning")
		if err != nil {
			return nil, err
		} else {
			return sessao, nil
		}
	}
}
