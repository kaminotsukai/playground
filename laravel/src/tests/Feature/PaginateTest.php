<?php

namespace Tests\Feature;

use App\Models\User;
use Illuminate\Foundation\Testing\RefreshDatabase;
use Tests\TestCase;

class PaginateTest extends TestCase
{
    use RefreshDatabase;

    public function test_paginate()
    {
        User::factory()->count(30)->create();
        $res = $this->get('/api/paginate');
        dump(json_decode($res->getContent(), true));

        dump("====================================");

        // 以下の箇所でリクエストに`page`があればその値を表示したいページとして保持しているため、paginateで指定しなくても良い
        // see: https://github.com/illuminate/pagination/blob/master/PaginationState.php#L20
        $res = $this->get('/api/paginate?page=2');
        dump(json_decode($res->getContent(), true));

        User::truncate();
    }
}
