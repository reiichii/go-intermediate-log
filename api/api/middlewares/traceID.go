package middlewares

import (
	"sync"
)

var (
	logNo int = 1
	mu    sync.Mutex
)

func newTraceID() int {
	var no int

	mu.Lock()
	no = logNo
	// 異なるゴールーチンから並行に呼び出されることが考えられるためロック
	logNo += 1
	mu.Unlock()
	return no
}

// common/values.goに移行
// func SetTraceID(ctx context.Context, traceID int) context.Context {
// 	return context.WithValue(ctx, "traceID", traceID)
// }

// type traceIDKey struct{}

// func GetTraceID(ctx context.Context) int {
// 	// id := ctx.Value("traceID") // valueの戻り値はint型になってしまう
// 	id := ctx.Value(traceIDKey{})
// 	// idをint型にアサーションする
// 	if idInt, ok := id.(int); ok {
// 		return idInt
// 	}
// 	return 0
// }
