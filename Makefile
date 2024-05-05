run: build
	@./game-gotdd

build: 
	@go build .

test: 
	@grc go test ./... -v -cover