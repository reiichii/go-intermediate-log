package services_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"example/api/services"
	_ "github.com/go-sql-driver/mysql"
)

var aSer *services.MyAppService

func TestMain(m *testing.M) {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	aSer = services.NewMyAppService(db)
	m.Run()
}

func BenchmarkGetArticleService(b *testing.B) { // Benchmark*という名前にする必要がある
	articleID := 1
	b.ResetTimer()             // 前処理を計測時間に含まないようにする
	for i := 0; i < b.N; i++ { // 何回実行されるかがb.N画よしなに決めてくれる
		_, err := aSer.GetArticleService(articleID)
		if err != nil {
			b.Error(err)
			break
		}
	}
}
