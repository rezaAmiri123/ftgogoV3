EVANS_PORT=8085

docker-clean-monolith:
	docker image rm ftgogov3-monolith -f

docker-build-monolith:
	docker build -t ftgogov3-monolith --file docker/Dockerfile .

docker-run-monolith:
	docker compose --profile monolith up
docker-rm-volume:
	docker volume rm ftgogov3_pgdata
docker-run:
	docker compose up
run-dev:
	go run cmd/ftgogo/*
evans:
	evans --host localhost --port $(EVANS_PORT) -r repl

psql:
	PGPASSWORD=ftgogo_pass psql -h postgres -U ftgogo_user -d ftgogo
