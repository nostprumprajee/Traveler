angular.module('devmtnTravel', ['ui.router'])
    .config(function ($stateProvider, $urlRouterProvider) {
        $stateProvider
            .state('home',{
                url:'/',
                templateUrl: "../views/about.html"
            })
            .state('adventurers',{
                url:'/adventurers',
                parent:'home',
                templateUrl: "../views/about-adventurers.html"
            })
            .state('contact',{
                url:'/contact',
                parent:'home',
                templateUrl: "../views/contact.html"
            })
            .state('packages',{
                url:'/packages',
                templateUrl: "../views/packages.html",
                controller: 'mainCtrl',
                controllerAs: 'vm'
            })
            .state('booked',{
                url:'/booked/:id',
                templateUrl: "../views/booked.html",
                controller: 'mainCtrl',
                controllerAs: 'vm'
            })
            .state('locations',{
                url:'/locations',
                templateUrl: "../views/locations.html",
                controller: 'mainCtrl',
                controllerAs: 'vm'
            });

        $urlRouterProvider
            .otherwise('/');
    });
