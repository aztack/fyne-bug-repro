set AppName="Bug Reproduction"
del "%AppName%.exe"
go-winres make
set CC=x86_64-w64-mingw32-gcc
set GOOS=windows
set GOARCH=amd64
set CGO_ENABLED=1
go build -o "%AppName%.exe" -v -ldflags -H=windowsgui .
upx "%AppName%.exe"
