package tests

import "fmt"

func IsStringLong(input string) bool {
	if len(input) > 5 {
		return true
	}
	return false
}

func main() {
	fmt.Println(IsStringLong("aaa"))
}
