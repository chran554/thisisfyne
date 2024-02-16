
.PHONY: all
all: generate
	go build -o bin/thisisfyne ./cmd/thisisfyne

.PHONY: generate
generate:
	go generate -v -x internal/app/resources/resources.go

.PHONY: macos-app
macos-app: all
	fyne package --executable bin/thisisfyne --name "Selfie evaluator" --appVersion 0.1.0 --icon assets/macos_icon.png
	mv "Selfie evaluator.app" bin/

.PHONY: clean
clean:
	rm -r bin/

.PHONY: demo
demo:
	go run fyne.io/fyne/v2/cmd/fyne_demo@latest
