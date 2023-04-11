<?php

namespace App\Http\Controllers;

use App\Models\User;
use Illuminate\Http\Request;

class PaginateController extends Controller
{
    public function paginate()
    {
        $users = User::paginate(5);
        return response()->json($users->toArray());
    }
}
