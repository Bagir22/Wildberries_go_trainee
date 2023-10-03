/*

var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}

Скорее всего проблема в том что большая строка останется в памяти из-за ссылки на подстроку 
Для исправления этой проблемы, можно воспользоваться функцией copy, которая создаст новую строку содержащую только нужную подстроку.
*/

package main

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = makeCopy(v[:100])
}

func main() {
	someFunc()
}

func createHugeString(size int) string {
	hugeString := make([]byte, size)
	for i := 0; i < size; i++ {
		hugeString[i] = 'a'
	}
	
	return string(hugeString)
}

func makeCopy(s string) string {
	copyStr := make([]byte, len(s))
	copy(copyStr, s)
	return string(copyStr)
}
