# GOOS: windows, linux, darwin
# GOARCH: 386, amd64, arm, arm64

all: windows

publish:
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o bin/be_logreport_x86.exe main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/be_logreport_x64.exe main.go
#	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o bin/be_logreport_linux_x86 main.go
#	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/be_logreport_linux_x64 main.go
#	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/be_logreport_darwin_amd64 main.go
#	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o bin/be_logreport_darwin_arm64 main.go

darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/be_logreport_darwin_amd64 main.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o bin/be_logreport_darwin_arm64 main.go

linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o bin/be_logreport_linux_x86 main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/be_logreport_linux_x64 main.go

windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o bin/be_logreport_x86.exe main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/be_logreport_x64.exe main.go

clean:
	rm -rf ./bin

.PHONY:publish
.PHONY:darwin
.PHONY:linux
.PHONY:windows
.PHONY:clean
