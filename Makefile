all: test lint

test:
	mkdir -p build
	go test -v -cover -coverprofile=build/coverage.txt
	go tool cover -html=build/coverage.txt -o=build/coverage.html

lint:
	go get -v github.com/alecthomas/gometalinter
	gometalinter --install
	gometalinter -E gofmt -D gotype -D vetshadow -D dupl -D goconst -D interfacer -D gas -D gocyclo\
	 --tests --deadline=10m --concurrency=2 --enable-gc

clean:
	rm -rf build

.PHONY: test lint clean
