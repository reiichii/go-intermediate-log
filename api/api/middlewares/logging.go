package middlewares

import (
	"log"
	"net/http"

	"example/api/common"
)

// 委譲によって、Header メソッド・Write メソッド・WriteHeaderメソッドを持ち、http.ResponseWriter インターフェースを満たす
type resLoggingWriter struct {
	http.ResponseWriter // 元々使用していた http.ResponseWriter を格納するためのフィールド
	code                int
}

func NewResLoggingWriter(w http.ResponseWriter) *resLoggingWriter {
	return &resLoggingWriter{ResponseWriter: w, code: http.StatusOK}
}

// ハンドラが HTTP レスポンスコードを書き込むときに使うメソッド
func (rsw *resLoggingWriter) WriteHeader(code int) {
	// resLoggingWriter構造体のcodeフィールドに、使うレスポンスコードを保存する
	rsw.code = code
	// HTTPレスポンスに使うレスポンスコードを指定(WriteHeader本来の機能の呼び出し)
	rsw.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		traceID := newTraceID()
		// log.Println(req.RequestURI, req.Method)
		log.Printf("[%d]%s %s\n", traceID, req.RequestURI, req.Method)
		// 自作のResponseWriterを作ってそれをハンドラに渡す
		// ServeHTTPにwを渡すやり方だと、next.ServeHTTPによって実行されたハンドラが、レスポンスに何と書き込んだのか、中身を確認できないため
		ctx := req.Context()
		// ctx := SetTraceID(req.Context(), traceID)
		ctx = common.SetTraceID(ctx, traceID)
		req = req.WithContext(ctx)
		rlw := NewResLoggingWriter(w)
		next.ServeHTTP(rlw, req)

		log.Printf("[%d]res: %d", traceID, rlw.code)
	})
}
