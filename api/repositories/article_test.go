package repositories_test

import (
	"testing"

	"example/api/models"
	"example/api/testdata"
	"repositories"

	_ "github.com/go-sql-driver/mysql"
)

func TestInsertArticle(t *testing.T) {
	article := models.Article{
		Title:    "insertTest",
		Contents: "testest",
		UserName: "user",
	}
	expectedArticleNum := 3
	newArticle, err := repositories.InsertArticle(testDB, article)
	if err != nil {
		t.Error(err)
	}
	if newArticle.ID != expectedArticleNum {
		t.Errorf("new article id is expected %d but got %d\n", expectedArticleNum, newArticle.ID)
	}
	t.Cleanup(func() {
		const sqlStr = `
			delete from articles
			where title = ? and contents = ? and username = ?
		`
		testDB.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	})
}

func TestSelectArticleList(t *testing.T) {
	expectedNum := len(testdata.ArticleTestData)
	got, err := repositories.SelectArticleList(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}
	if num := len(got); num != expectedNum {
		t.Errorf("want %d but got %d articles\n", expectedNum, num)
	}
}

func TestSelectArticleDetail(t *testing.T) {
	// expected := models.Article{
	// 	ID:       1,
	// 	Title:    "firstPost",
	// 	Contents: "This is my first blog",
	// 	UserName: "saki",
	// 	NiceNum:  4,
	// }
	// got, err := repositories.SelectArticleDetail(db, expected.ID)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// if got.ID != expected.ID {
	// 	t.Errorf("ID: get %d but want %d\n", got.ID, expected.ID)
	// }
	// if got.Contents != expected.Contents {
	// 	t.Errorf("Contents: get %s but want %s\n", got.Contents, expected.Contents)
	// }
	// if got.UserName != expected.UserName {
	// 	t.Errorf("Contents: get %s but want %s\n", got.UserName, expected.UserName)
	// }
	// if got.NiceNum != expected.NiceNum {
	// 	t.Errorf("Contents: get %d but want %d\n", got.NiceNum, expected.NiceNum)
	// }

	// table driven testで複数の入力値をテストするコードを書く
	tests := []struct {
		testTitle string
		expected  models.Article
	}{
		{
			testTitle: "subtest1",
			expected:  testdata.ArticleTestData[0],
		}, {
			testTitle: "subtest2",
			expected:  testdata.ArticleTestData[1],
		},
	}
	for _, test := range tests {
		t.Run(test.testTitle, func(t *testing.T) {
			got, err := repositories.SelectArticleDetail(testDB, test.expected.ID)
			if err != nil {
				t.Fatal(err)
			}
			if got.ID != test.expected.ID {
				t.Errorf("ID: get %d but want %d\n", got.ID, test.expected.ID)
			}
			if got.Title != test.expected.Title {
				t.Errorf("Title: get %s but want %s\n", got.Title, test.expected.Title)
			}
			if got.Contents != test.expected.Contents {
				t.Errorf("Title: get %s but want %s\n", got.Contents, test.expected.Contents)
			}
			if got.UserName != test.expected.UserName {
				t.Errorf("Title: get %s but want %s\n", got.UserName, test.expected.UserName)
			}
			if got.NiceNum != test.expected.NiceNum {
				t.Errorf("Title: get %d but want %d\n", got.NiceNum, test.expected.NiceNum)
			}
		})
	}
}

func TestUpdateNiceNum(t *testing.T) {
	articleID := 1
	before, err := repositories.SelectArticleDetail(testDB, articleID)
	if err != nil {
		t.Fatal("fail to get before data")
	}
	err = repositories.UpdateNiceNum(testDB, articleID)
	if err != nil {
		t.Fatal(err)
	}
	after, err := repositories.SelectArticleDetail(testDB, articleID)
	if err != nil {
		t.Fatal("fail to get after data")
	}
	if after.NiceNum-before.NiceNum != 1 {
		t.Error("fail to update nice num")
	}
}
