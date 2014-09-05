'use strict';

var megaRunningRoute = angular.module('megaRunningRoute', []);

megaRunningRoute.config(['$routeProvider',
  function($routeProvider) {
    $routeProvider.
      when('/provas', {
        templateUrl: 'views/provas/lista.html',
        controller: 'ProvasListaCtrl'
      }).
      otherwise({
        redirectTo: '/'
      });
  }]);