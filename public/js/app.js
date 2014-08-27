'use strict';

/* App Module */

var MegaRunningApp = angular.module('MegaRunningApp', [
  'ngRoute',
  'MegaRunningControllers',
  'MegaRunningServices'
]);

MegaRunningApp.config(['$routeProvider',
  function($routeProvider) {
    $routeProvider.
      when('/provas', {
        templateUrl: 'partials/prova-list.html',
        controller: 'ProvaListCtrl'
      }).
      when('/provas/:id', {
        templateUrl: 'partials/prova-detail.html',
        controller: 'ProvaDetailCtrl'
      }).
      otherwise({
        redirectTo: '/provas'
      });
  }]);

var MegaRunningServices = angular.module('MegaRunningServices', ['ngResource']);

MegaRunningServices .factory('Prova', ['$resource',
  function($resource){
    return $resource('http://localhost:port/provas',{port: ':8080'}, {
      //query: {method:'GET'}
      query: {method:'GET', params:{id:'provas'}, isArray:true}
    });
  }]);

var MegaRunningControllers = angular.module('MegaRunningControllers', []);

MegaRunningControllers.controller('ProvaListCtrl', ['$scope', 'Prova',
 
  function($scope, Prova) {
    $scope.provas = Prova.query();
}]);

MegaRunningControllers.controller('ProvaDetailCtrl', ['$scope', '$routeParams', 'Prova',
  function($scope, $routeParams, Prova) {
    console.log($routeParams.id);
    $scope.prova = Prova.get({id: $routeParams.id});
}]);