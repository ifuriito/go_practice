package main

import "fmt"

func main() {
	totalScore := 0
	ask(1, "hoge", &totalScore)
	ask(2, "hogehoge", &totalScore)
	ask(3, "hogehogehoge", &totalScore)
	fmt.Println("スコア", totalScore)
}

// スコアの計算もポインタを使って、ask関数内で行う
func ask(number int, question string, scorePtr *int) {
	var input string
	fmt.Printf("[質問%d] 次の単語を入力してください: %s\n", number, question)
	fmt.Scan(&input)
	fmt.Printf("%sと入力されました\n", input)

	// 正解、不正解の判定をスコアに反映
	if question == input {
		fmt.Println("正解!")
		*scorePtr += 10
	} else {
        fmt.Println("不正解!")
	}
}