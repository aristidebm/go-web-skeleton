.PHONY: format run


run:
	@go run ./...


format:
	@go fmt ./...
