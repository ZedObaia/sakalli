build:
	go build -o bin/sakalli main.go

PORT = "8080"
run:
	go run main.go --port=$(PORT)

compile:
	# 32-Bit Systems
	# FreeBDS
	GOOS=freebsd GOARCH=386 go build -o bin/sakalli-freebsd-386 main.go
	# MacOS
	GOOS=darwin GOARCH=386 go build -o bin/sakalli-darwin-386 main.go
	# Linux
	GOOS=linux GOARCH=386 go build -o bin/sakalli-linux-386 main.go
	# Windows
	GOOS=windows GOARCH=386 go build -o bin/sakalli-windows-386 main.go
	# 64-Bit
	# FreeBDS
	GOOS=freebsd GOARCH=amd64 go build -o bin/sakalli-freebsd-amd64 main.go
	# MacOS
	GOOS=darwin GOARCH=amd64 go build -o bin/sakalli-darwin-amd64 main.go
	# Linux
	GOOS=linux GOARCH=amd64 go build -o bin/sakalli-linux-amd64 main.go
	# Windows
	GOOS=windows GOARCH=amd64 go build -o bin/sakalli-windows-amd64 main.go

clean:
	rm -fr bin