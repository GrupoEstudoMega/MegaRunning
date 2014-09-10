'use strict'

var megaRunningApp = angular.module('megaRunningApp', [
  'ngRoute',
  'megaRunningControllers',
  'megaRunningServices',
  'megaRunningRoute'
]);

var megaRunningControllers = angular.module('megaRunningControllers', []);

var megaRunningServices = angular.module('megaRunningServices', ['ngResource']);