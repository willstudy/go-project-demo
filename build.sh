#!/bin/bash

work_dir=`dirname $0`
work_dir=`cd ${work_dir};pwd`

# unit testing
go test github.com/liuchang/stringutil

# The go tool finds the source code by looking for the `github.com/liuchang/hello` package inside the workspace specified by GOPATH 
go install github.com/liuchang/hello
