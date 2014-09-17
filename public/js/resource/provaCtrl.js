'use strict';

/* Controllers */

megaRunningControllers.controller('ProvasListaCtrl', ['$scope', 'Prova', function($scope, Prova) {
  $scope.provas = Prova.query();
}]);

megaRunningControllers.controller('DetalheCtrl', ['$scope', '$routeParams','Prova',
  function($scope, $routeParams, Prova) {
    $scope.prova = Prova.get({id:$routeParams.id});
    $scope.submit = function() {
    	$scope.prova.$save({id:$routeParams.id});
    };
  }]);