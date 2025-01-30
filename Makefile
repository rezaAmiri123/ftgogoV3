EVANS_PORT=8085
install-tools:
	@echo installing tools
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go install github.com/bufbuild/buf/cmd/buf@latest
	@go install github.com/vektra/mockery/v2@latest
	@go install github.com/go-swagger/go-swagger/cmd/swagger@latest
	@go install github.com/cucumber/godog/cmd/godog@latest
	@echo done

docker-clean-monolith:
	docker image rm ftgogov3-monolith -f

docker-build-monolith:
	docker build -t ftgogov3-monolith --file docker/Dockerfile .

docker-run-monolith:
	docker compose --profile monolith up
docker-rm-volume-pgdata:
	docker volume rm ftgogov3_pgdata ftgogov3_jsdata
docker-run:
	docker compose up
run-dev:
	GODEBUG=httpmuxgo121=1 go run cmd/ftgogo/*
evans:
	evans --host localhost --port $(EVANS_PORT) -r repl

psql:
	PGPASSWORD=ftgogo_pass psql -h postgres -U ftgogo_user -d ftgogo

pprof-heap:
	go tool pprof -http localhost:9090 http://localhost:6060/debug/pprof/heap?debug=1

pprof-goroutine:
	go tool pprof -http localhost:9090 http://localhost:6060/debug/pprof/goroutine?debug=1
test:
	go test ./... -count 1