migrateup:
	migrate -path db/migration -database "mysql://duyle:duyle95@tcp(127.0.0.1:3308)/food_delivery" -verbose up

migratedown:
	migrate -path db/migration -database "mysql://duyle:duyle95@tcp(127.0.0.1:3308)/food_delivery" -verbose down
