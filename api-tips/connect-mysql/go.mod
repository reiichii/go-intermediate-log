module dbsample

go 1.19

require github.com/go-sql-driver/mysql v1.7.0 // indirect

require example.com/models v0.0.0

replace example.com/models => ./models
