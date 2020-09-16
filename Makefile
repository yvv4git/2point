.PHONY: build

BUILD_DIR='build'

build:
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/server64.bin app/cmd/main.go
	#GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/server64.exe app/cmd/main.go

tests:
	go test -v ./...

start:
	./build/server64.bin 

docker_build:
	docker-compose build

docker_start:
	docker-compose up -d --force-recreate --build

docker_stop:
	docker-compose stop

curl_count:
	curl -X GET http://localhost:8080/count

curl_userid:
	curl -X GET http://localhost:8080/?user_id=123