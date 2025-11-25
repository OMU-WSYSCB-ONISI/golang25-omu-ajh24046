package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	// Week 03: ここに課題のコードを記述してください
	// 詳細な課題内容はLMSで確認してください

	fmt.Println("Week 03 課題")

	// 以下に実装してください
	http.HandleFunc("/webfortune", webfortunehandler)

	http.ListenAndServe(":8080", nil)

}

func webfortunehandler(w http.ResponseWriter, r *http.Request) {
	seed := time.Now().UnixNano()
	d := rand.New(rand.NewSource(seed))
	randnum := d.Int31n(4)
	switch randnum {
	case 0:
		fmt.Fprintln(w, "今日の運勢は大吉です")
	case 1:
		fmt.Fprintln(w, "今日の運勢は中吉です")
	case 2:
		fmt.Fprintln(w, "今日の運勢は吉です")
	case 3:
		fmt.Fprintln(w, "今日の運勢は凶です")
	}
}
