package main

import "fmt"

func main() {
	// totalScoreにaskの戻り値を代入
	totalScore := ask(1, "hoge")
	totalScore += ask(2, "hogehoge")
	totalScore += ask(3, "hogehoge")
	fmt.Println("スコア", totalScore)
}

// 戻り値の型を定義
func ask(number int, question string) int{
	var input string
	fmt.Printf("[質問%d] 次の単語を入力してください: %s\n", number, question)
	fmt.Scan(&input)
	fmt.Printf("%sと入力されました\n", input)

	// 正解、不正解の判定をスコアに反映
	if question == input {
        fmt.Println("正解!")
        return 10
	} else {
        fmt.Println("不正解!")
        return 0
	}
}