go install github.com/bufbuild/buf/cmd/buf@latest
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest
[ -n "$(go env GOBIN)" ] && export PATH="$(go env GOBIN):${PATH}"
[ -n "$(go env GOPATH)" ] && export PATH="$(go env GOPATH)/bin:${PATH}"
echo "----Done----"


# // install grpcurl
# wget https://github.com/fullstorydev/grpcurl/releases/download/v1.7.0/grpcurl_1.7.0_linux_x86_64.tar.gz
# --2024-05-05 14:36:37--  https://github.com/fullstorydev/grpcurl/releases/download/v1.7.0/grpcurl_1.7.0_linux_x86_64.tar.gz
# tar -xvf grpcurl_1.7.0_linux_x86_64.tar.gz
# chmod +x grpcurl
# ./grpcurl -help