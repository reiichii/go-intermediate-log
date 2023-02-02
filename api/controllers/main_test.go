package controllers_test

import (
	"testing"

	"example/api/controllers"
	"example/api/controllers/testdata"
)

var aCon *controllers.ArticleController

func TestMain(m *testing.M) {
	// ser := services.NewMyAppService(db) コントローラ層に渡すサービス層構造体をモック構造体 に差し替える
	ser := testdata.NewServiceMock()
	aCon = controllers.NewArticleController(ser)

	m.Run()
}
