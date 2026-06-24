build-go:
	if [ ! -d "go-server-grpc-local" ]; then mkdir go-server-grpc-local; fi
	protoc --experimental_allow_proto3_optional -I ./ --go_out=go-server-grpc-local --go-grpc_out=go-server-grpc-local ./proto/*/*.proto

build-vue:
	if [ ! -d "vue-client-grpc-local" ]; then mkdir vue-client-grpc-local; fi
	protoc -I ./ --experimental_allow_proto3_optional --plugin=protoc-gen-ts=./node_modules/.bin/protoc-gen-ts --ts_out=vue-client-grpc-local ./proto/*/*.proto

prepare-building-vue:
	if [ ! -d "vue-client-grpc-local" ]; then mkdir vue-client-grpc-local; fi
	cd vue-client-grpc-local && npm install --save-dev @protobuf-ts/plugin && npm install @protobuf-ts/runtime

build-swift:
	protoc \
		--swift_out=swift-grpc-local \
		--grpc-swift_out=swift-grpc-local \
		--swift_opt=Visibility=Public \
		--plugin=protoc-gen-grpc-swift=/opt/homebrew/Cellar/protoc-gen-grpc-swift/2.1.1/bin/protoc-gen-grpc-swift-2 \
		--grpc-swift_opt=Visibility=Public,Client=true,Server=false \
		./proto/*/*.proto

build-desc:
	protoc \
		--descriptor_set_out=output.desc \
 		--include_imports \
 		./proto/*/*.proto
