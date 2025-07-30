BINARY_NAME=cci-attendance-apis
FOLDER_NAME=dist
.DEFAULT_GOAL := run

build:
	GOOARCH=amd64 GOOS=darwin go build -o ./${FOLDER_NAME}/${BINARY_NAME}-darwin
	GOOARCH=amd64 GOOS=linux go build -o ./${FOLDER_NAME}/${BINARY_NAME}-linux
	GOOARCH=amd64 GOOS=windows go build -o ./${FOLDER_NAME}/${BINARY_NAME}-windows

run: build
	./${FOLDER_NAME}/${BINARY_NAME}-darwin

build_and_run: build run

clean:
	go clean
	rm ./${FOLDER_NAME}/${BINARY_NAME}-darwin
	rm ./${FOLDER_NAME}/${BINARY_NAME}-linux
	rm ./${FOLDER_NAME}/${BINARY_NAME}-windows

test:
	go test ./...

test_coverage:
	go test ./... -coverprofile=coverage.out

dep:
	go mod download

vet:
	go vet

lint:
	golangci-lint run --enable-all