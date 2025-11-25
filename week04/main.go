package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Week 04: ここに課題のコードを記述してください
	// 詳細な課題内容はLMSで確認してください

	fmt.Println("Week 04 課題")

	// 以下に実装してください
	http.HandleFunc("/info", infohandler)
	http.ListenAndServe(":8080", nil)
}

func infohandler(w http.ResponseWriter, r *http.Request) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	fmt.Fprintln(w, "現在時刻は"+(time.Now().In(jst)).Format("15:04:05")+"です")
	var h map[string][]string
	h = r.Header
	fmt.Fprintln(w, "使用しているブラウザは"+h["User-Agent"][0]+"です")
}
