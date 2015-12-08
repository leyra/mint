all: mint

mint: deps
	go build

deps:
	go get github.com/tdewolff/minify
	go get github.com/spf13/cobra
