proto-data:
	@protoc --proto_path api/resume_v1 --go_out=pkg/resume_v1/data
			--go_opt=paths=source_relative --go-grpc_out=pkg/resume_v1/data
			--go-grpc_opt=paths=source_relative api/resume_v1/resumeDataV1.proto
server-keys:
	cd pkg/tlsconfig/cert/server/; sh gen.sh;

client-keys:
	cd pkg/tlsconfig/cert/client/; sh gen.sh;

build:
	@go build -o resumegamesrv cmd/server/resume_server.go



run:
	./resumegamesrv