build-go:
	if [ ! -d "go-server-grpc-local" ]; then mkdir go-server-grpc-local; fi
	protoc --experimental_allow_proto3_optional -I ./ --go_out=go-server-grpc-local --go-grpc_out=go-server-grpc-local ./proto/*/*.proto

build-vue:
	if [ ! -d "vue-client-grpc-local" ]; then mkdir vue-client-grpc-local; fi
	protoc -I ./ --experimental_allow_proto3_optional --plugin=protoc-gen-ts=./node_modules/.bin/protoc-gen-ts --ts_out=vue-client-grpc-local ./proto/*/*.proto

prepare-building-vue:
	if [ ! -d "vue-client-grpc-local" ]; then mkdir vue-client-grpc-local; fi
	cd vue-client-grpc-local && npm install --save-dev @protobuf-ts/plugin && npm install @protobuf-ts/runtime
