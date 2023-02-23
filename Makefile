RunApp:
	docker-compose up -d

StopApp:
	docker-compose down

SqlMigrate:
	sql-migrate up	

SqlcGen:
	sqlc generate

postgesql:
	docker run --name gdn -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:14-alpine	

.PHONY: RunApp, StopApp, SqlMigrate, postgesql