build:
	go mod tidy
	go build -o ./dist/nsrun

install:
	install --mode +x ./dist/nsrun /usr/local/bin

clean:
	rm -fr ./dist