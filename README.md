# Kogito Local Server

## Requirements

Golang version: `1.15`

## Run

To run simply execute `make run`.

## Build

So far the build was tested on MacOs

To build execute `make build` from root path.

For specific platforms please execute:

- `make mac`
- `make win`
- `make linux`

The binaries are going to appear in each OS folder. But if you execute just `make build` the `default` folder will contain the binaries.

## Runner

The runner must be compiled and copied to `runner` folder. Remember to change the configuration file to match the runner name.

## Configuration

In the `config.yaml` file you will be able to configure Proxy, Runner and Modeler properties as ip, port, runner location or modeler URL.

## Next Steps

- Change Java Quarkus with Native Quarkus runner
- Provide compilation pipeline for Windows
- Provide compilation pipeline for Linux
- Limit GraalVM Heap Size
- Check version alignment between Online Editor and DMN Runner.
- Installers for all platforms

## Extrass

### How do I create the image.go?

`cat icon2.png | /Users/aparedes/go/bin/2goarray Data images > icon.go`
