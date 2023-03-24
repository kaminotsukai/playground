<?php

// カプセル化
class Capsule
{
    // プライベート変数が外部から直接アクセスできないことで、オブジェクトの不正な状態が防止される（データの保護）
    // パブリック変数にするとバリデーション処理を追加できないので適切に制御できなくなる
    private string $name;
    private int    $age;

    const DEFAULT_NAME = 'anonymous';

    public function __construct(string $name, int $age)
    {
        $this->name = $name;
        $this->age  = $age;
    }

    // 不正な値が入らないようにメソッドを経由して変数を変更する（データの保護）
    // 例：空文字の場合はデフォルト名を代入する、31文字以上の名前はつけれないなど
    public function setName(string $name): void
    {
        if (empty($name)) {
            $this->name = self::DEFAULT_NAME;
        } else {
            $this->name = $name;
        }
    }

    // 単純なアクセサメソッドでも値の正当性をチェックするために意味がある
    // 将来的にルールが追加されるかもしれないのでアクセサメソッドを介して変数を変更するほうが拡張性が高い
    public function setAge(int $age): void
    {
        $this->age = $age;
    }

    public function toArray(): array
    {
        return [
            'name' => $this->name,
            'age'  => $this->age,
        ];
    }
}

$user = new Capsule("Tom", 21);
$user->setName("");
print_r($user->toArray(), true);
