go env -w GOOS=linux
go build -o target

go env -w GOOS=windows
go build -o target

