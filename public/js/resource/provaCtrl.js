'use strict';

/* Controllers */

megaRunningControllers.controller('ProvasListaCtrl', ['$scope', 'Prova', function($scope, Prova) {
  $scope.provas = Prova.query();

  $scope.delete = function(prova){
  	//$scope.prova = Prova.query({id: id});
    Prova.delete(prova);
  	
  	$scope.provas.splice($scope.provas.indexOf(prova), 1);
    console.log(prova);
    console.log($scope.provas.indexOf(prova));
  }
}]);

megaRunningControllers.controller('DetalheCtrl', ['$scope', '$routeParams', '$window', 'Prova',
  function($scope, $routeParams, $window, Prova) {
    Prova.get({id:$routeParams.id}).$promise.then(function(prova) {
    	for (var x = 0; x < prova.Participantes.length; x++) {
    		prova.Participantes[x].NomeAtleta = prova.Participantes[x].Atleta.Nome.Primeiro + " " + prova.Participantes[x].Atleta.Nome.Sobrenome 
    	}
    	$scope.prova = prova
    })

    $scope.submit = function() {
    	$scope.prova.$save({id:$routeParams.id});
    	$window.location.href = '#/provas';
    };
  }]);

megaRunningControllers.controller('NewCtrl', ['$scope', '$routeParams', '$window', 'Prova',
  function($scope, $routeParams, $window, Prova) {
  	$scope.prova = new Prova();
  	$scope.prova.Valor = 0;
    $scope.submit = function() {
    	$scope.prova.$save();
    	$window.location.href = '#/provas';
    };
  }]);

