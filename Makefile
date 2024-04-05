VERSION ?= 0.0.1
IMG_NAME ?= bti/lotso-airdrop-server:${VERSION}
DB_NAME ?= lotso

darwin-amd64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./bin/lotso-airdrop-darwin-amd64 .

linux-amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/lotso-airdrop-linux-amd64 .

win-amd64:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./bin/lotso-airdrop-windows-amd64 .

docker:
	docker build -t ${IMG_NAME} -f Dockerfile .

start-docker:
	IMG_NAME=${IMG_NAME} DB_NAME=${DB_NAME} docker-compose -f docker-compose.yaml up -d
