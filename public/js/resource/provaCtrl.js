'use strict';

/* Controllers */

megaRunningControllers.controller('ProvasListaCtrl', ['$scope', 'Prova', function($scope, Prova) {
  $scope.provas = Prova.query();
}]);