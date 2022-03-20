#! /bin/bash
set -e
echo ">>> Generate grpchook Proto."
protoc -I ./grpchook --go_out ./grpchook --go_opt paths=source_relative \
    --go-grpc_out=./grpchook --go-grpc_opt paths=source_relative \
    ./grpchook/grpchook.proto
echo ">>> Generate grpchook Proto OK."