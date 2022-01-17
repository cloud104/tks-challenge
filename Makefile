all: compile upload

compile:
	mkdir -p ./dist
	GOARCH=amd64 GOOS=linux go build -o ./dist/challenge main.go
	cd dist ; sha256sum * > SHA256SUMS

clean:
	rm -rf ./dist

upload:
	gsutil -m cp dist/* gs://tks-challenge