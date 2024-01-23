#!/bin/bash
rm -f pb/*.go
rm -f docs/swagger/*.swagger.json
protoc --proto_path=proto --proto_path=proto/models --proto_path=proto/rpc \
    --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    --grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
    --openapiv2_out=docs/swagger --openapiv2_opt=allow_merge=true,merge_file_name=pharmago \
    proto/*.proto proto/models/*.proto proto/rpc/*.proto
cp pb/models/*pb.go pb
cp pb/rpc/*pb.go pb
cp pb/rpc/*pb.go pb
rm -r pb/models
rm -r pb/rpc
# statik -src=./docs/swagger -dest=./docs