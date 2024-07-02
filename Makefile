gen-user:
	protoc protos/user/user.proto --go_out=./protos/go/gen/ --go_opt=paths=source_relative \
	--go-grpc_out=./protos/go/gen --go-grpc_opt=paths=source_relative

gen-product:
	protoc protos/product/product.proto --go_out=./protos/go/gen/ --go_opt=paths=source_relative \
	--go-grpc_out=./protos/go/gen --go-grpc_opt=paths=source_relative