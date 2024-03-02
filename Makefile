APP_BUILD 	?= 1
APP_ID 		= thisisfyne.cha.se
APP_VERSION ?= 0.0.1

.PHONY: all
all: init clean generate
	go build -o bin/thisisfyne ./cmd/thisisfyne

.PHONY: generate
generate:
	go generate -v -x internal/app/resources/resources.go

.PHONY: app-darwin
app-darwin: all
	fyne package --executable bin/thisisfyne --name "This is Fyne" --appBuild=$(APP_BUILD) --appID=$(APP_ID) --appVersion=$(APP_VERSION) --icon assets/thisisfyne_icon.png
	mv "This is Fyne.app" bin/

.PHONY: cross-all
cross-all: clean cross-linux cross-windows cross-darwin

.PHONY: cross-linux
cross-linux: generate
	fyne-cross linux -arch=amd64,arm64 -icon=assets/thisisfyne_icon.png -app-build=$(APP_BUILD) -app-id=$(APP_ID) -app-version=$(APP_VERSION) ./cmd/thisisfyne

.PHONY: cross-windows
cross-windows: generate
	fyne-cross windows -arch=amd64,arm64 -icon=assets/thisisfyne_icon.png -app-build=$(APP_BUILD) -app-id=$(APP_ID) -app-version=$(APP_VERSION) ./cmd/thisisfyne

.PHONY: cross-darwin
cross-darwin: generate
	fyne-cross darwin -arch=amd64,arm64 -icon=assets/thisisfyne_icon.png -app-build=$(APP_BUILD) -app-id=$(APP_ID) -app-version=$(APP_VERSION) ./cmd/thisisfyne

.PHONY: clean
clean: init
	rm -rf bin/*
	rm -rf fyne-cross

.PHONY: init
init:
	mkdir -p bin

.PHONY: demo
demo:
	go run fyne.io/fyne/v2/cmd/fyne_demo@latest
