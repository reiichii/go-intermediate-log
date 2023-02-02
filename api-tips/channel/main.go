package main

import (
	"fmt"
	"strings"
)

func doubleInt(src int, intCh chan<- int) { // この関数内では送信専用チャネルになる
	result := src * 2
	intCh <- result // 受信側の準備ができていなかったら、できるまで待つ
}

func doubleString(src string, strCh chan<- string) {
	result := strings.Repeat(src, 2)
	strCh <- result
}

func main() {
	// var ch chan int のように書くと、chの値はnilになる
	ch1, ch2 := make(chan int), make(chan string) // make関数を使うことでchに送受信できるチャネルが格納される. 送受信可能
	defer close(ch1)
	defer close(ch2)

	go doubleInt(1, ch1)
	go doubleString("hello", ch2)
	// result := <-ch // 受診まで待ってくれる

	for i := 0; i < 2; i++ {
		select { // 複数のチャネルを待ち、受け取った順に処理していく
		case numResult := <-ch1:
			fmt.Println(numResult)
		case strResult := <-ch2:
			fmt.Println(strResult)
		}
	}

}
