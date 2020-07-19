package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

var t int

func init() {
	//オプションで制限時間をできる
	flag.IntVar(&t, "t", 1, "制限時間(分)")
	flag.Parse()
}

func main() {
	var (
		files     = flag.Args()
		path, err = os.Executable()
		ch_rcv    = myinput(os.Stdin)
		tm        = time.After(time.Duration(t) * time.Minute)
		words     []string
		score     = 0
	)
	if err != nil {
		//実行ファイルのパスが取得できなかったときのエラー処理
		fmt.Fprintln(os.Stderr, "読み込みに失敗しました", err)
		os.Exit(0)
	} else if len(files) != 1 {
		//コマンドの引数に1つ以上のテキストファイルが指定された時のエラー処理
		fmt.Fprintln(os.Stderr, "指定できるファイル数は1つだけです")
		os.Exit(0)
	}
	//実行ファイルのディレクトリまでのPathを取得
	path = filepath.Dir(path)
	words, err = getwords_from(filepath.Join(path, files[0]))
	if err != nil {
		//テキストファイルパスが取得できない場合のエラー処理
		fmt.Fprintln(os.Stderr, "読み込みに失敗しました", err)
		os.Exit(0)
	}

	fmt.Println("タイピングゲームを始めます。制限時間は", t, "分。1語1点、", len(words), "点満点")
	//送信用チャネル
	for i := true; i && score < len(words); {
		question := words[score]
		fmt.Print(question)
		fmt.Print(">")
		select {
		case x := <-ch_rcv:
			//標準入力に何か入力された時の処理
			// 入力された文字が一致しているかどうかをチェックする
			if question == x {
				score++
			}
		case <-tm:
			//制限時間が過ぎた際の処理
			fmt.Println("制限時間を過ぎました")
			i = false
		}
	}
	fmt.Println("あなたの点数:", score, "点 / ", len(words), " 点")

}

func myinput(r io.Reader) <-chan string {
	// サブgo ルーチン
	// 標準入力から受け取った文字列を標準出力へ出力する
	ch1 := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch1 <- s.Text()
		}
	}()
	return ch1
}

func getwords_from(txt_path string) ([]string, error) {
	// テキストファイルのパスからテキストに書き込まれた単語をリスト化して返す
	var words []string
	sf, err := os.Open(txt_path)
	if err != nil {
		return nil, err
	} else {
		scanner := bufio.NewScanner(sf)
		for scanner.Scan() {
			words = append(words, scanner.Text())
		}
	}
	return words, err
}
