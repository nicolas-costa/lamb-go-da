build-lambda-windows:
				cmd /C "SET "GOOS=linux" && SET "GOARCH=amd64" && go build main.go -o bin/main"