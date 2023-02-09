build:
	go build -o bin/yugioh ./cmd/api/main.go

run: build
	./bin/yugioh

clean:
	rm -rf ./bin/
