docker-build-monolith:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o monolith ../../cmd/ftgogo/main.go
	docker build -t ftgogov3-monolith .
	rm monolith
