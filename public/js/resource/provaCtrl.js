'use strict';

/* Controllers */

megaRunningControllers.controller('ProvasListaCtrl', ['$scope', 'Prova', function($scope, Prova) {
  $scope.provas = Prova.query();
  $scope.delete = function(id){
  	Prova.delete({id: id});
  	//$scope.provas = Prova.query();
  	$scope.provas.splice($scope.provas.indexOf($scope.prova), 1);
  }
}]);

megaRunningControllers.controller('DetalheCtrl', ['$scope', '$routeParams', '$window', 'Prova',
  function($scope, $routeParams, $window, Prova) {
    $scope.prova = Prova.get({id:$routeParams.id});
    $scope.submit = function() {
    	$scope.prova.$save({id:$routeParams.id});
    	$window.location.href = '#/provas';
    };
  }]);

megaRunningControllers.controller('NewCtrl', ['$scope', '$routeParams', '$window', 'Prova',
  function($scope, $routeParams, $window, Prova) {
  	$scope.prova = new Prova();
    $scope.submit = function() {
    	$scope.prova.$save();
    	$window.location.href = '#/provas';
    };
  }]);