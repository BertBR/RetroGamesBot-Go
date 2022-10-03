start: ensure
	air

ensure:
	bash -x scripts/ensure.sh

dev: ensure
	docker-compose up -d --build

log: 
	docker logs -f retrogames
