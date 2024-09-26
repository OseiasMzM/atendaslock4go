set CGO_ENABLED=1

go build -o build/atendaslock4go.exe

call "build/atendaslock4go.exe"

pause