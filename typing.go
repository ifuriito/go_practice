package main

import "fmt"

func main() {

	ask(1, "hoge")

	ask(2, "hogehoge")

	ask(3, "hogehogehgoe")
}

func ask(number int, question string) {
	var input string
	fmt.Printf("[質問%d] 次の単語を入力してください: %s\n", number, question)
	fmt.Scan(&input)
	fmt.Printf("%sと入力されました\n", input)
}