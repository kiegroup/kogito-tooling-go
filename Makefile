
all: all-jitexecutor build

build: clean build-default

build-default:
	go build -o build/default/runner main.go
	chmod +x ./build/default/runner

run:
	ENV=dev go run main.go

clean:
	$(RM) -rf ./build

# macOS
macos: clean all-jitexecutor build-macos package-macos

build-macos: 
	GO111MODULE=on GOOS=darwin GOARCH=amd64 go build -o build/darwin/runner main.go
	chmod +x ./build/darwin/runner

package-macos:
	cd build/darwin && zip -qry runner-macos.zip runner

# Linux
linux: clean all-jitexecutor build-linux package-linux

build-linux:
	GOOS=linux GOARCH=amd64 go build -o build/linux/runner main.go

package-linux:
	cd build/linux && tar -pcvzf runner-linux.tar.gz runner

# Windows
win:clean all-jitexecutor build-win

build-win:
	GOOS=windows GOARCH=386 go build -ldflags "-H=windowsgui" -o build/win/runner main.go
	chmod +x ./build/win/runner

# Jit Executor
all-jitexecutor: build-jitexecutor copy-jitexecutor

build-jitexecutor:
	mvn clean package -DskipTests -f ./kogito-apps/jitexecutor && mvn clean package -DskipTests -Pnative -f ./kogito-apps/jitexecutor

copy-jitexecutor:
	cp ./kogito-apps/jitexecutor/jitexecutor-runner/target/jitexecutor-runner-*-SNAPSHOT-runner jitexecutor
	chmod +x jitexecutor