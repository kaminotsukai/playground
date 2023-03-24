package memory

// 文字列のスライスから重複を削除する処理をメモリ効率を意識して実装する
func RemoveDuplicates(values []string) []string {
	encountered := make(map[string]struct{})
	result := make([]string, 0, len(values)) // make関数はサイズオーバーすると2倍のメモリを再確保するため、事前に決めうちで入れる
	for _, value := range values {
		if _, ok := encountered[value]; !ok {
			encountered[value] = struct{}{} // struct{}型は空の構造体で、サイズがゼロであるためメモリ使用量を最小限に抑えることができる
			result = append(result, value)
		}
	}
	return result
}
