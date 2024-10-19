EVANS_PORT=8085

clean-monolith:
	docker image rm ftgogov3-monolith -f

build-monolith:
	docker build -t ftgogov3-monolith --file docker/Dockerfile .

run-monolith:
	docker compose --profile monolith up

run:
	docker compose up

evans:
	evans --host localhost --port $(EVANS_PORT) -r repl
