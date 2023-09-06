build:
	go mod tidy
	go build -o ./dist/execns

clean:
	rm -fr ./dist