package main

import (
	"context"
	"fmt"

	"google.golang.org/api/idtoken"
)

func main() {
	googleClientID := "960218929572-3p258eica7kuenvl9j9e3mctekj5vvja.apps.googleusercontent.com"
	idToken, err := readFile("./idtoken.txt")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	// バリデータ構造体を生成
	tokenValidator, err := idtoken.NewValidator(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	// 検証の実施
	payload, err := tokenValidator.Validate(context.Background(), idToken, googleClientID)
	if err != nil {
		fmt.Println("validate err: ", err)
		return
	}

	fmt.Println(payload.Claims["name"])
}
