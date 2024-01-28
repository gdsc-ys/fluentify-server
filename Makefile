# Define where the *.proto files are located.
PROTO_DIR = idl/proto

# Define the output directory in which to place the build results.
PROTO_GEN_DIR = gen

CURR_DIR = $(PWD)

proto:
	docker run --rm \
	--volume $(CURR_DIR):/workspace \
	--workdir /workspace \
	bufbuild/buf generate

GO111MODULE = on
CGO_ENABLED = 0
GOOS = linux
GOARCH = amd64
GO_BUILD_OUT_DIR = build

build: proto
	mkdir -p $(GO_BUILD_OUT_DIR)
	go mod download
	go build -o $(GO_BUILD_OUT_DIR)/main .

clean:
	rm -rf $(GO_BUILD_OUT_DIR)
