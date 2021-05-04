
all: submodule all-jitexecutor build

build: clean build-default

build-default:
	go build -o build/default/dmn_runner main.go

run:
	ENV=dev go run main.go

clean:
	$(RM) -rf ./build

submodule:
	git submodule update

# macOS
macos: clean all-jitexecutor build-macos package-macos

build-macos: 
	GO111MODULE=on GOOS=darwin GOARCH=amd64 go build -o build/darwin/dmn_runner main.go

package-macos:
	cd scripts/macos && ./build.sh

# Linux
linux: clean all-jitexecutor build-linux package-linux

build-linux:
	GOOS=linux GOARCH=amd64 go build -o build/linux/dmn_runner main.go

package-linux:
	cd build/linux && tar -pcvzf dmn_runner_linux.tar.gz dmn_runner

# Windows
win:clean all-jitexecutor build-win

build-win:
	GOOS=windows GOARCH=386 go build -ldflags "-H=windowsgui" -o build/win/dmn_runner main.go

# Jit Executor
all-jitexecutor: build-jitexecutor copy-jitexecutor

build-jitexecutor:
	mvn clean package -DskipTests -f ./kogito-apps/jitexecutor && mvn clean package -DskipTests -Pnative -am -f ./kogito-apps/jitexecutor

copy-jitexecutor:
	cp ./kogito-apps/jitexecutor/jitexecutor-runner/target/jitexecutor-runner-*-SNAPSHOT-runner jitexecutor
	chmod +x jitexecutor
