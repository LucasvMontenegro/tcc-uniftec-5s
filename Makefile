unit-test:
	gotestsum -- ./internal/... -failfast -race -coverprofile ./coverage.out

mocks:
	cd internal && rm -rf mocks && mockery --all --keeptree --inpackage

run:
	go run ./cmd/main.go

create-migration:
	@read -p "Migration file name: " MIGRATION_FILE; \
	migrate create -dir ./db/migrations -ext sql -seq $$MIGRATION_FILE