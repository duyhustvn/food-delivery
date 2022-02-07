1. Create database if not exists

``` shell
CREATE SCHEMA `food_delivery` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci ;

```
2. Create migration file 
``` shell
migrate create -ext sql -dir db/migration -seq init_schema
```
3. Run migration up

``` shell
make migrateup
```
