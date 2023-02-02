package main

import (
	"errors"
	"fmt"
	"strconv"
)

type MyError struct {
	InternalError error
}

func (e *MyError) Error() string {
	return e.InternalError.Error()
}

func (e *MyError) Unwrap() error {
	return e.InternalError
}

func getMyError(err error) error {
	return &MyError{err}
}

func main() {
	var err1, err2 error
	_, err1 = strconv.Atoi("a") // strconv.NumError型
	err2 = getMyError(err1)     // MyError型（内部にNumError型

	var ok bool

	_, ok = err1.(*strconv.NumError)
	fmt.Println(ok) // 元が NumError 型のエラーインターフェースを NumError に変換 -> true
	_, ok = err2.(*strconv.NumError)
	fmt.Println(ok) // 元が MyError 型 (内部に NumError 型) のエラーインターフェースをNumErrorに変換 -> false
	// 内部に NumError 型を含んでいる err2 を NumError 型にアサーション することができていない

	_, ok = err1.(*MyError)
	fmt.Println(ok) // 元が NumError 型のエラーインターフェースを MyError に変換 -> false
	_, ok = err2.(*MyError)
	fmt.Println(ok) // 元が MyError 型 (内部に NumError 型) のエラーインターフェースを MyErrorに変換 -> true

	// errors.As
	var err3, err4 error
	_, err3 = strconv.Atoi("a")
	err4 = getMyError(err1)

	var targetNumErr *strconv.NumError
	fmt.Println(errors.As(err3, &targetNumErr)) // → true
	fmt.Println(errors.As(err4, &targetNumErr)) // → true  内部に NumError 型を含んでいる err4 を NumError 型にアサーションができるようになった

	var targetMyErr *MyError
	fmt.Println(errors.As(err3, &targetMyErr)) // → false
	fmt.Println(errors.As(err4, &targetMyErr)) // → true
}
