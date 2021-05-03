
all: build build-default

build: clean jitexecutor

linux: build build-linux

mac: build build-mac

windows: build build-windows

clean:
	$(RM) -rf ./build

mac: clean build-mac 

build-mac: 
	GO111MODULE=on GOOS=darwin GOARCH=amd64 go build -o build/darwin/runner main.go

linux: clean build-linux 

build-linux:
	GOOS=linux GOARCH=amd64 go build -o build/linux/runner main.go

win:clean build-win

build-win:
	GOOS=windows GOARCH=386 go build -ldflags "-H=windowsgui" -o build/win/runner main.go

build-default:
	go build -o build/default/runner main.go

jitexecutor: build-jitexecutor copy-jitexecutor

build-jitexecutor:
	mvn clean package -DskipTests -f ./kogito-apps/jitexecutor && mvn clean package -DskipTests -Pnative -f ./kogito-apps/jitexecutor

copy-jitexecutor:
	cp ./kogito-apps/jitexecutor/jitexecutor-runner/target/jitexecutor-runner-*-SNAPSHOT-runner jitexecutor

run:
	ENV=dev go run main.go
