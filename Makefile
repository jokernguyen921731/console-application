install:
	go mod download

test:
	go test ./...

run:
	go build -o build/email-application ./src/cmd
	chmod +x build/email-application