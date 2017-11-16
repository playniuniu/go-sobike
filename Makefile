# Binary name
BINARY=sobike
VERSION=0.1
LDFLAGS='-w -s'

# Builds the project
build:
	go clean
	GOPATH=${GOPATH}:`pwd` go build -o ${BINARY} -ldflags ${LDFLAGS} src/${BINARY}.go

release:
	# Make release folder
	mkdir release/

	# Build for mac
	go clean
	GOPATH=${GOPATH}:`pwd` go build -ldflags ${LDFLAGS} src/${BINARY}.go
	upx ./${BINARY}
	tar czvf release/${BINARY}-mac64-${VERSION}.tar.gz ./${BINARY}

	# Build for arm
	go clean
	GOPATH=${GOPATH}:`pwd` CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags ${LDFLAGS} src/${BINARY}.go
	upx ./${BINARY}
	tar czvf release/${BINARY}-arm64-${VERSION}.tar.gz ./${BINARY}

	# Build for linux
	go clean
	GOPATH=${GOPATH}:`pwd` CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags ${LDFLAGS} src/${BINARY}.go
	upx ./${BINARY}
	tar czvf release/${BINARY}-linux64-${VERSION}.tar.gz ./${BINARY}

	# Build for win
	go clean
	GOPATH=${GOPATH}:`pwd` CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags ${LDFLAGS} src/${BINARY}.go
	upx ./${BINARY}.exe
	tar czvf release/${BINARY}-win64-${VERSION}.tar.gz ./${BINARY}.exe

	# Clean binary
	rm -rf ${BINARY}
	rm -rf ${BINARY}.exe

# Cleans our projects: deletes binaries
clean:
	go clean
	rm -rf ${BINARY}
	rm -rf ${BINARY}.exe
	rm -rf release/

.PHONY:  clean build
