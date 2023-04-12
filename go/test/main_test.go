package tests

import (
	"testing"
)

// 確認（アサート）してなくても、実行されていればカバレッジは計測される（確認不在のテスト）
func Test_IsStringLong1(t *testing.T) {
	_ = IsStringLong("abc")
}
