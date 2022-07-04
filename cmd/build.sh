#!/bin/bash
if !(type go >/dev/null 2>&1)
then
    echo 'go: command not found'
    exit
fi

# build
if [[ $1 == "static" ]]
then
    go build -ldflags '-linkmode "external" -extldflags "-static"' -o ../bin/syscat syscat/main.go
    go build -ldflags '-linkmode "external" -extldflags "-static"' -o ../bin/sysctl sysctl/main.go

else
    export GOPROXY=https://goproxy.cn
    gofmt -w .
    go build -o ../bin/syscat syscat/main.go
    go build -o ../bin/sysctl sysctl/main.go
fi
