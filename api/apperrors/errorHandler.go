package apperrors

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"example/api/common"
)

func ErrorHandler(w http.ResponseWriter, req *http.Request, err error) {
	var appErr *MyAppError
	if !errors.As(err, &appErr) { // errors.As(): 第三引数で受け取った err を MyAppError 構造体とみなして、内部の ErrCode フィールド等 を取り出すために、エラーインターフェース型errをMyAppError型に変換する
		appErr = &MyAppError{
			ErrCode: Unknown,
			Message: "internal process failed",
			Err:     err,
		}
	}

	// traceID := middlewares.GetTraceID(req.Context()) // ハンドラがmiddlewaresに依存してしまうのがよくない
	traceID := common.GetTraceID(req.Context())
	log.Printf("[%d]error: %s\n", traceID, appErr)

	var statusCode int

	// error codeに対応するhttpステータスコードを返すように定義
	switch appErr.ErrCode {
	case NAData:
		statusCode = http.StatusNotFound
	case NoTargetData, ReqBodyDecodeFailed, BadParam:
		statusCode = http.StatusBadRequest
	case RequiredAuthorizationHeader, Unauthorized:
		statusCode = http.StatusUnauthorized
	case NotMatchUser:
		statusCode = http.StatusForbidden
	default:
		statusCode = http.StatusInternalServerError
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(appErr)
}
