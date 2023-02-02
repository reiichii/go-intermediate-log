package middlewares

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"google.golang.org/api/idtoken"

	"example/api/apperrors"
	"example/api/common"
)

// common/values.goに移行
// type userNameKey struct{}

// func GetUserName(ctx context.Context) string {
// 	id := ctx.Valuse(userNameKey{})
// 	if usernameStr, ok := id.(string); ok {
// 		return usernameStr
// 	}
// 	return ""
// }

// func SetUserName(req *http.Request, name string) *http.Request {
// 	ctx := req.Context()

// 	ctx = context.WithValue(ctx, userNameKey{}, name)
// 	req = req.WithContext(ctx)

// 	return req
// }

const (
	googleClientID = "960218929572-3p258eica7kuenvl9j9e3mctekj5vvja.apps.googleusercontent.com"
)

// 認証工程の中でエラーが発生したらハンドラに処理を回さず、即座に エラー処理に移行させる( next.ServeHTTP が実行されないようにしている)
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorization := req.Header.Get("Authorization")

		// Bearer id_token 形式になっているかの検証
		authHeaders := strings.Split(authorization, " ")
		if len(authHeaders) != 2 {
			err := apperrors.RequiredAuthorizationHeader.Wrap(errors.New("invalid req header"), "invalid header")
			apperrors.ErrorHandler(w, req, err)
			return
		}
		bearer, idToken := authHeaders[0], authHeaders[1]
		if bearer != "Bearer" || idToken == "" {
			err := apperrors.RequiredAuthorizationHeader.Wrap(errors.New("invalid req header"), "invalid header")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		// token検証
		tokenValidator, err := idtoken.NewValidator(context.Background())
		if err != nil {
			err = apperrors.CannotMakeValidator.Wrap(err, "internal auth error")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		payload, err := tokenValidator.Validate(context.Background(), idToken, googleClientID)
		if err != nil {
			err = apperrors.Unauthorized.Wrap(err, "invalid id token")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		name, ok := payload.Claims["name"]
		if !ok {
			err = apperrors.Unauthorized.Wrap(err, "invalid id token")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		// req = SetUserName(req, name.(string))
		req = common.SetUserName(req, name.(string))

		next.ServeHTTP(w, req)
	})
}
