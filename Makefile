start: ensure
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	migrate -path ./pkg/storage/postgres/migrations -database ${DATABASE_URL} -verbose up
	air

ensure:
	bash -x scripts/ensure.sh

dev: ensure
	docker-compose up -d --build

log: 
	docker logs -f retrogames
