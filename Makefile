
.PHONY: all
all:
	go build -o bin/thisisfyne ./cmd/thisisfyne

.PHONY: macos
macos: all
	fyne package --executable bin/thisisfyne --name "Selfie evaluator" --appVersion 0.1.0 --icon documentation/macos_icon.png
	mv "Selfie evaluator.app" bin/

.PHONY: clear
clear:
	rm -r bin/