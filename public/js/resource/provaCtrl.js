'use strict';

/* Controllers */

var megaRunningControllers = angular.module('megaRunningControllers', []);

megaRunningControllers.controller('ProvasListaCtrl', function($scope) {
  $scope.provas = [
    {
	  _id: "53e2a47e33e6261e20000002",
	  nome: "Corrida Noturna",
	  url: "www.itu.com.br",
	  local: "Itu",
	  data: {
	    "$date": null
	  },
	  valor: 25.5,
	  participantes: [
	    {
	      valorpatrocinio: 12.3,
	      atleta: {
	        _id: "53e2a47e33e6261e20000001",
	        nome: {
	          primeiro: "João",
	          nomeDoMeio: "Fernando",
	          sobrenome: "Barros"
	        },
	        email: "teste@teste"
	      }
	    }
	  ]
	},
	{
	  _id: "53ebd2284129e713c4000002",
	  nome: "Corrida Noturna",
	  url: "",
	  local: "",
	  data: "1-01-01T00:00:00Z",
	  valor: 0,
	  participantes: [
	    {
	      valorpatrocinio: 12.300000190734863,
	      atleta: {
	        _id: "53ebd2284129e713c4000001",
	        nome: {
	          primeiro: "teste",
	          nomeDoMeio: "do",
	          sobrenome: "teste"
	        },
	        email: "teste@teste"
	      }
	    }
	  ]
	}
  ];
});