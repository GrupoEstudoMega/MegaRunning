megaRunningServices.factory('Prova', ['$resource',
  function($resource){
    return $resource('http://localhost:8080/api/provas/:id', {}, {
    	
    });
  }]);