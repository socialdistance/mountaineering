DATABASE_URL := "postgres://postgres:postgres@localhost:54321/mountaineering?sslmode=disable"

migration:
	goose --dir=migrations postgres ${DATABASE_URL} up