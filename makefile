.PHONY: build build-win dev run clean

PROJECT="game-gorl"
BUILD_PATH="./build"

init:
	mkdir build
	mkdir runtime

build-win:
	CGO_LDFLAGS="-static-libgcc -static -lpthread"\
				GOOS=windows GOARCH=amd64 CGO_ENABLED=1\
				CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++\
				go build -a -x -o $(BUILD_PATH)/$(PROJECT)-win.exe

build:
	mkdir -p $(BUILD_PATH)
	cp -r assets/* $(BUILD_PATH)
	go build -o $(BUILD_PATH)/$(PROJECT)-linux

run:
	cd $(BUILD_PATH); ./$(PROJECT)-linux

dev:
	make build; make run

clean:
	rm -r $(BUILD_PATH)/*
	mkdir -p $(BUILD_PATH)
	cp -r assets/* $(BUILD_PATH)
