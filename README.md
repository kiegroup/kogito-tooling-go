# Kogito Local Server

## Requirements

Golang version: `1.16`

## Application Parameters

- `-p <PORT_NUMBER>`: Sets app port, otherwise it will use config.yaml port.

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

The runner must be compiled from the "kogito-apps" repository and copied to the `pkg/kogito` folder.

## Fedora

To use this application on Fedora, it's necessary to install some additional packages and enable the Gnome App Indicator.
Firstly install the following packages:

- `sudo dnf install gtk3-devel libappindicator-gtk3-devel-12.10.0-29.fc33.x86_64`

To enable the App Indicator extension
https://extensions.gnome.org/extension/615/appindicator-support/

## Configuration

In the `config.yaml` file you will be able to configure Proxy, Runner and Modeler properties as runner location or modeler URL. Runner ip is `127.0.0.1` and port is a random free port.

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
