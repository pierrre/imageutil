all: test lint

test:
	go test -v

lint:
	go get -v -u github.com/alecthomas/gometalinter
	gometalinter --install --update --no-vendored-linters
	GOGC=800 gometalinter --enable-all -D dupl -D lll -D gas -D goconst -D gocyclo -D gotype -D interfacer -D safesql -D test -D testify -D vetshadow\
	 --tests --deadline=10m --concurrency=4 --enable-gc

clean:
	rm -rf build

.PHONY: test lint clean
