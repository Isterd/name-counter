BINARY := ./bin/name-counter
CMD    := ./cmd/name-counter

.PHONY: build test clean run

build:
	go build -o $(BINARY) $(CMD)

test:
	go test ./...

clean:
	rm -f $(BINARY)

run: build
	./$(BINARY) examples/names.txt