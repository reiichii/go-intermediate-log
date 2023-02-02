package api

import (
	"database/sql"
	"net/http"

	"example/api/api/middlewares"
	"example/api/controllers"
	"example/api/services"

	"github.com/gorilla/mux"
)

// func NewRouter(aCon *controllers.ArticleController, cCon *controllers.CommentController) *mux.Router {
func NewRouter(db *sql.DB) *mux.Router {
	ser := services.NewMyAppService(db)
	aCon := controllers.NewArticleController(ser)
	cCon := controllers.NewCommentController(ser)

	r := mux.NewRouter()
	r.HandleFunc("/hello", aCon.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", aCon.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", aCon.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", aCon.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", aCon.PostNiceHandler).Methods(http.MethodPost)

	r.HandleFunc("/comment", cCon.PostCommentHandler).Methods(http.MethodPost)

	// middlewares.LoggingMiddleware(cCon.PostCommentHandler)) を都度描く代わりにこれでラップしてくれる
	// r.Useで指定する順番注意
	// 今回の場合は、loggingの前処理,authの前処理,authの後処理,loggingの後処理 のようになる
	r.Use(middlewares.LoggingMiddleware)
	r.Use(middlewares.AuthMiddleware)

	return r
}
