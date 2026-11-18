<?php

return [

    /*
    |--------------------------------------------------------------------------
    | Cross-Origin Resource Sharing (CORS) Configuration
    |--------------------------------------------------------------------------
    |
    | Autorise le front Angular à appeler l'API.
    | L'adresse du front se règle avec FRONTEND_URL dans .env
    | (http://localhost:4200 en local).
    |
    */

    'paths' => ['v1/*'],

    'allowed_methods' => ['*'],

    'allowed_origins' => explode(',', env('FRONTEND_URL', 'http://localhost:4200')),

    'allowed_origins_patterns' => [],

    'allowed_headers' => ['*'],

    'exposed_headers' => [],

    'max_age' => 0,

    'supports_credentials' => false,

];
