package repositories_test

import (
	"example/api/models"
	_ "github.com/go-sql-driver/mysql"
	"repositories"
	"testing"
)

func TestSelectCommentList(t *testing.T) {
	articleID := 1
	expectedNum := 2
	got, err := repositories.SelectCommentList(testDB, articleID)
	if err != nil {
		t.Fatal(err)
	}
	if num := len(got); num != expectedNum {
		t.Errorf("want %d but get %d articles\n", expectedNum, num)
	}
}

func TestInsertComment(t *testing.T) {
	articleID := 1
	comment := models.Comment{
		ArticleID: articleID,
		Message:   "test",
	}
	before, err := repositories.SelectCommentList(testDB, articleID)
	if err != nil {
		t.Error(err)
	}
	newComment, err := repositories.InsertComment(testDB, comment)
	if err != nil {
		t.Error(err)
	}
	after, err := repositories.SelectCommentList(testDB, articleID)
	if len(after)-len(before) != 1 {
		t.Errorf("InsertComment is not expected")
	}

	t.Cleanup(func() {
		const sqlStr = `
		delete from comments
		where comment_id = ?
		`
		testDB.Exec(sqlStr, newComment.CommentID)
	})
}
