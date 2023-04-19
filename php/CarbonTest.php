<?php declare(strict_types=1);
require __DIR__ . '/vendor/autoload.php';

use Carbon\Carbon;
use PHPUnit\Framework\TestCase;

final class CarbonTest extends TestCase
{
    /**
     * @dataProvider carbonDataProvider
     *
     * @return void
     */
    public function testCarbon(array $input, array $expected): void
    {
        $next = Carbon::create($expected[0], $expected[1], $expected[2], 0, 0, 0);
        $actual = Carbon::create($input[0], $input[1], $input[2], 0, 0, 0)->addMonthNoOverflow(); // 月末問題を解消するためのメソッド
        $this->assertEquals($next, $actual);

        $next = Carbon::create($expected[0], $expected[1], $expected[2], 0, 0, 0);
        $actual = Carbon::create($input[0], $input[1], $input[2], 0, 0, 0)->addMonth();
        $this->assertEquals($next, $actual);
    }

    public static function carbonDataProvider()
    {
        return [
            [[2023, 1, 25], [2023, 2, 25]],
            [[2023, 2, 25], [2023, 3, 25]],
            [[2023, 3, 25], [2023, 4, 25]],
            [[2023, 4, 25], [2023, 5, 25]],
            [[2023, 5, 25], [2023, 6, 25]],
            [[2023, 6, 25], [2023, 7, 25]],
            [[2023, 7, 25], [2023, 8, 25]],
            [[2023, 8, 25], [2023, 9, 25]],
            [[2023, 9, 25], [2023, 10, 25]],
            [[2023, 10, 25], [2023, 11, 25]],
            [[2023, 11, 25], [2023, 12, 25]],
            [[2023, 12, 25], [2024, 1, 25]],
        ];
    }
}
