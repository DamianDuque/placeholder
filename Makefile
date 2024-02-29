.PHONY: install run proto

run:
	@go docker-compose up

proto:
	@python3 -m grpc_tools.protoc -I ./protobuf --python_out=./peer/src/protobuf --pyi_out=./peer/src/protobuf --grpc_python_out=./peer/src/protobuf ./protobuf/filesystem.proto

install:
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@echo Installation complete.
	@echo Remember to add \`\export \GOPATH\=$\HOME/go\` 
	@echo and \`\export \PATH\=$\PATH:$GOPATH/bin\`
	@echo to your .bashrc file.

#https://get.golang.org/${{uname}}/go_installer
