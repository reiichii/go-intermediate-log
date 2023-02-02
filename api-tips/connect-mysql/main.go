package main

import (
	"database/sql"
	"fmt"

	"example.com/models"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	// if err := db.Ping(); err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("connect to DB")
	// }

	// const sqlStr = `
	// 	select *
	// 	from articles;
	// `
	articleID := 0
	const sqlStr = `
		select *
		from articles
		where article_id = ?;
	`
	row := db.QueryRow(sqlStr, articleID)
	if err := row.Err(); err != nil {
		// 0件だった場合
		fmt.Println(err)
		return
	}
	// defer rows.Close()

	// articleArray := make([]models.Article, 0)
	var article models.Article
	var createdTime sql.NullTime
	// for rows.Next() {
	// 	// rows の中に格納されている取得レコード内容を読み出す
	// 	// err := rows.Scan(&article.Title, &article.Contents, &article.UserName, &article.NiceNum)
	// 	err := rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)
	// 	if createdTime.Valid {
	// 		article.CreatedAt = createdTime.Time
	// 	}
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	// } else {
	// 	// 	articleArray = append(articleArray, article)
	// 	// }
	// }
	err = row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)
	if err != nil {
		fmt.Println(err)
		return
	}
	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}
	fmt.Printf("%+v\n", article)

}
