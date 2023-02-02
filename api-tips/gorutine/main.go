package main

import (
	"fmt"
	"time"
)

func cutIngredient() {
	fmt.Println("start cutting ingredients")
	time.Sleep(1 * time.Second)
	fmt.Println("finish cutting ingredients")
}

func boilWater() {
	fmt.Println("start boiling water")
	time.Sleep(2 * time.Second)
	fmt.Println("the water has boiled!")
}

func main() {
	go cutIngredient() // メインのとは別のゴールーチンで実行される
	boilWater()
}
