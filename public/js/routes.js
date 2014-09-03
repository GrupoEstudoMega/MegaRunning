console.log(megaRunningApp)
megaRunningApp.config(['$routeProvider',
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