.PHONY: pb
pb: 
	cd pb/ && protoc -I=./ --go_out=./ --go-grpc_out=./ *.proto 