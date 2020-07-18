package main

import "fmt"

func main() {
	var input string
	fmt.Println("次の単語を入力してください: test");
	fmt.Scan(&input);
	fmt.Printf("%sと入力されました",input);
}