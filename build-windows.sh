AppName="Bug Reproduction"
rm "$AppName.exe"
go-winres make
env CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 CGO_ENABLED=1 go build -o "$AppName.exe" -v -ldflags -H=windowsgui .
upx "$AppName.exe"