
clean:
	$(RM) -rf ./build

copy-mac:
	cp -r runner/ build/darwin/runner
	cp config.yaml build/darwin/

copy-linux:
	cp -r runner/ build/linux/runner
	cp config.yaml build/linux/

copy-win:
	cp -r runner/ build/win/runner
	cp config.yaml build/win/

copy-default:
	cp -r runner/ build/default/runner
	cp config.yaml build/default/

build-mac: 
	GO111MODULE=on GOOS=darwin GOARCH=amd64 go build -o build/darwin/kogito main.go

build-linux:
	GOOS=linux GOARCH=amd64 go build -o build/linux/kogito main.go

build-win:
	GOOS=windows GOARCH=386 go build -ldflags "-H=windowsgui" -o build/win/kogito main.go

build-default:
	go build -o build/default/kogito main.go

mac: clean build-mac copy-mac

linux: clean build-linux copy-linux

win:clean build-win copy-win

build: clean build-default copy-default

run:
	go run main.go
