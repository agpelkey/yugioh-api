build:
	go build -o bin/yugioh 

run: build
	./bin/yugioh

clean:
	rm -rf ./bin/
