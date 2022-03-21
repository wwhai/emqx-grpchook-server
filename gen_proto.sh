#! /bin/bash
set -e
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOROOT:$GOPATH:$GOBIN

echo ">>> Generate grpchook Proto."
protoc -I ./grpchook --go_out ./grpchook --go_opt paths=source_relative \
    --go-grpc_out=./grpchook --go-grpc_opt paths=source_relative \
    ./grpchook/grpchook.proto
echo ">>> Generate grpchook Proto OK."