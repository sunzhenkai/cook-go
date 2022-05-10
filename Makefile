.PHONY: proto mod

mod:
	go mod tidy && go mod download

proto:
	@protoc -I=assets/proto --go_out=pkg/gen/pb assets/proto/*.proto
