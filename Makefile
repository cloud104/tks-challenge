compile:
	mkdir -p ./dist
	GOARCH=amd64 GOOS=linux go build -o ./dist/challenge main.go

clean:
	rm -rf ./dist

all: compile