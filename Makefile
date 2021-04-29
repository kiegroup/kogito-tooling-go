
all: build

clean:
	$(RM) -rf ./build

build-mac: 
	GO111MODULE=on GOOS=darwin GOARCH=amd64 go build -o build/darwin/kogito main.go

build-linux:
	GOOS=linux GOARCH=amd64 go build -o build/linux/kogito main.go

build-win:
	GOOS=windows GOARCH=386 go build -ldflags "-H=windowsgui" -o build/win/kogito main.go

build-default:
	go build -o build/default/kogito main.go

mac: clean build-mac 

linux: clean build-linux 

win:clean build-win

build: clean build-default

run:
	ENV=dev go run main.go
