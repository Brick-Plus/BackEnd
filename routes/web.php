<?php

use Illuminate\Support\Facades\Route;

Route::get('/', fn () => response()->json([
    'name' => 'BrickPlus API',
    'api' => url('/v1'),
]));
