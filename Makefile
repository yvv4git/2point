.PHONY: build

BUILD_DIR='build'

build:
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/server64.bin app/cmd/main.go
	#GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/server64.exe app/cmd/main.go

tests:
	go test -v ./...