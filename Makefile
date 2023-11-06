.PHONY: build-lb
build-lb:
	go build -o bin/lb cmd/lb.go cmd/helper.go

.PHONY: build-server
build-server:
	go build -o bin/server cmd/server.go cmd/helper.go
