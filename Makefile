build:
	go mod tidy
	go build -o ./dist/execns

install:
	install --mode +x ./dist/execns /usr/local/bin

clean:
	rm -fr ./dist