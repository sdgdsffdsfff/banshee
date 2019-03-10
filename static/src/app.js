angular.module('banshee', [

  'ngSanitize',
  'ngAnimate',
  'ngAria',
  'ngMessages',
  'ngFileUpload',

  'ui.router',
  'ui.bootstrap',
  'ngMaterial',

  'pascalprecht.translate',

  'toastr',

  'banshee.tpl',

  require('./services'),
  require('./filters'),
  require('./directives'),

  require('./modules/layout'),
  require('./modules/metric'),
  require('./modules/admin')

]).config(require('./config'))
  .run(require('./run'));
