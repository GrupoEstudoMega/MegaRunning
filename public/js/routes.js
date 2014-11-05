'use strict';

var megaRunningRoute = angular.module('megaRunningRoute', []);

megaRunningRoute.config(['$routeProvider',
  function($routeProvider) {
    $routeProvider.
      when('/provas', {
        templateUrl: 'views/provas/lista.html',
        controller: 'ProvasListaCtrl'
      }).
      when('/provas/new', {
        templateUrl: 'views/provas/detalhe.html',
        controller: 'NewCtrl'
      }).      
      when('/provas/:id/edit', {
        templateUrl: 'views/provas/edit.html',
        controller: 'DetalheCtrl'
      }).
      when('/provas/:id', {
        templateUrl: 'views/provas/view.html',
        controller: 'DetalheCtrl'
      }).      
      otherwise({
        redirectTo: '/'
      });
  }]);