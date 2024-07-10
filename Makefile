
.PHONY: test
test:
	go test ./... -v

.PHONY: run
run:
	docker-compose up --build
