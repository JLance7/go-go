# Learning golang

## compiling 
To compile: ```go build <file.go>``` or ```go build -o <name> <file.go>```
- To compile for windows from mac: ```GOOS=windows GOARCH=amd64 go build -o myprogram.exe```
- To compile for mac from windows: ```SET GOOS=darwin GOARCH=amd64 go build -o myprogram```
- To compile for linux from windows: ```SET GOOS=linux GOARCH=amd64 go build -o myprogram```

## running
To run: ```./<file_generated.exe>```

## all at once
To compile and run: ```go run <file.go>```

## package manager usage
```go get <package_name>```