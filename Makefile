# building compidown on go

BINARY_NAME = compidown
OUTPUT_DIR = bin

#build for your own os
build:
	go build -C go -o ../$(OUTPUT_DIR)/$(BINARY_NAME) main.go

default : build

# build for any other OS

build_windows_amd:
	GOOS=windows GOARCH=amd64 go build -C go  -o ../$(OUTPUT_DIR)/$(BINARY_NAME)_win_amd.exe main.go

build_windows_arm:
	GOOS=windows GOARCH=arm64 go build -C go  -o ../$(OUTPUT_DIR)/$(BINARY_NAME)_win_arm.exe main.go

# build for mac
build_darwin_arm:
	GOOS=darwin GOARCH=arm64 go build -C go -o ../$(OUTPUT_DIR)/$(BINARY_NAME)_darwin_arm main.go

build_darwin_amd:
	GOOS=darwin GOARCH=amd64 go build -C go -o ../$(OUTPUT_DIR)/$(BINARY_NAME)_darwin_amd main.go

# build for every os:
build_all : build_darwin_amd build_darwin_arm build_windows_amd build_windows_arm build