package apperrors

type MyAppError struct {
	ErrCode // 型を省略した場合、型名がそのままフィールド名になる(ErrCode型)
	Message string
	Err     error `json:"-"` // エラーチェーンのための内部エラー. jsonエンコードされないようにする

}

// ある構造体をエラーとして扱うためにはエラーインターフェースを満たしていないといけない。帰り値がstring型のError()メソッドが必要になるため、エラーメソッドを定義する
func (myErr *MyAppError) Error() string {
	return myErr.Err.Error()
}

// エラーチェーンを組み込むためにErr フィールドに格納された内部エラーを返り値として返すよう な Unwrap メソッドを作る
func (myErr *MyAppError) Unwrap() error {
	return myErr.Err
}

// 独自エラーにラップしたくなった場所で都度MyAppError構造体を作るのは良くないので、ラップメソッドを作って共通化する
func (code ErrCode) Wrap(err error, message string) error {
	return &MyAppError{ErrCode: code, Message: message, Err: err}
}
