* postgreはLastInsertIdを返す仕組みがないが、 insert句の末尾に`returning article_id;`句を追記し、その値を取得することで対応が可能
