//Criação do App
var App = Ember.Application.create()

App.Router.map(function() {
  this.resource('provas', { path: 'provas' }, 
  	function() {
    	this.route('prova', { path: '/:prova_id' })
  })
})

App.IndexRoute = Ember.Route.extend({
  redirect: function() {
    this.transitionTo('provas')
  }
});

App.ProvasIndexRoute = Ember.Route.extend({
  model: function() {
    return this.store.find('prova')
  }
});

App.Prova = DS.Model.extend({
    Nome : DS.attr('string'),
    Url : DS.attr('string'),
    Local : DS.attr('string'),
    Data : DS.attr('date'),
    Valor : DS.attr('number')
});

App.ProvaAdapter = DS.RESTAdapter.extend({
	host: 'http://localhost:8080'
})

Ember.Handlebars.helper('datetime', function(value, options) {
  //var format = 'D MMMM, YYYY [at] HH:mm';
  var format = 'DD/MM/YYYY'
  if(options.hash.format) {
    format = options.hash.format
  }
  if(value) {
    var time = moment(value).format(format)
    return new Handlebars.SafeString('<span class="timestamp">' + time + '</span>')
  }
});