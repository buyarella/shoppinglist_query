# Buyarella Shopping List Query service

This service serves queries of all sorts concerning Shopping Lists

## Build

### API
This service provides a grpc API. To generate the new stubs after changing it:
* make sure to have [protoc](https://github.com/protocolbuffers/protobuf/releases) installed
* make sure to have [protoc-gen-go](https://github.com/golang/protobuf) installed; run `install_proto.sh`
* run `make api`