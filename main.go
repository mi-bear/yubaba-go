package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("湯婆婆: 契約書だよ。そこに名前を書きな。")

	for {
		name, err := promptName(reader)
		if err != nil {
			fmt.Println("なにをもたもたしてるんだい: ", err)
			return
		}

		fmt.Printf("フン。%s というのかい。贅沢な名だねぇ。\n", name)

		newName := stealName(name)
		fmt.Printf("今からお前の名前は「%s」だ。いいかい、%s だよ。分かったら返事をするんだ、%s!!\n",
			newName, newName, newName)

		// 確認
		fmt.Printf("この名前でいいかい？ (y/n): ")
		ans, _ := reader.ReadString('\n')
		ans = strings.TrimSpace(ans)
		if strings.ToLower(ans) != "y" {
			fmt.Println("そうかい...ならもう一度やろうか。")
			continue
		}

		fmt.Println("契約はこれで決まりさ。元の名は忘れな！ さっさと働きな！")
		break
	}
}

// promptName はユーザーに名前を入力させ、空でない名前を返します。
func promptName(r *bufio.Reader) (string, error) {
	for {
		fmt.Print("名前を入力: ")
		input, err := r.ReadString('\n')
		if err != nil {
			return "", err
		}
		input = strings.TrimSpace(input)
		if input == "" {
			fmt.Println("名前を空にしてどうするんだい！ しっかり書きな、もう一度だよ。")
			continue
		}
		return input, nil
	}
}

// promptNameAllowEmpty はユーザーに名前を入力させ、空でもそのまま返します。
// func promptNameAllowEmpty(r *bufio.Reader) (string, error) {
// 	fmt.Print("名前を入力: ")
// 	input, err := r.ReadString('\n')
// 	if err != nil {
// 		return "", err
// 	}
// 	input = strings.TrimSpace(input)
// 	return input, nil
// }

// stealName は name からランダムに 1 文字を選んで返します。
func stealName(name string) string {
	runes := []rune(name)
	n := len(runes)
	idx := rand.Intn(n)
	return string(runes[idx])
}
